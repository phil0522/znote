//go:generate protoc --go_out=plugins=grpc:. proto/service.proto

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/phil0522/znote/pkg/server"
	"github.com/phil0522/znote/pkg/subcommands"
	pb "github.com/phil0522/znote/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/sevlyar/go-daemon.v0"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.WarnLevel)
}

const (
	usage = `Usage: znote <command> [<args>]

command can be of
new: create a new note
list: list all notes
pick: pick a note and edit it
edit: edit a note directly
`
)

const (
	serverAddr = "127.0.0.1:6399"
)

func executeSubCommand(ctx context.Context, client pb.ZNoteServiceClient) {

}

func killServer(ctx context.Context, client pb.ZNoteServiceClient) {
	_, err := client.QuitServer(ctx, &pb.QuitServerRequest{})
	if err != nil {
		log.Fatalf("failed to quit server rpc %s", err.Error())
	}
}

func clientCall(callback func(context.Context, pb.ZNoteServiceClient)) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create connection, %s", err.Error())
	}

	client := pb.NewZNoteServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	callback(ctx, client)
}

func serve() {
	logrus.Info("starting server")
	lis, err := net.Listen("tcp4", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer func() {
		log.Printf("server exit")
	}()
	grpcServer := grpc.NewServer()
	znoteServer := &server.ZNoteServer{
		GrpcServer: grpcServer,
	}

	znoteServer.Initialize()
	pb.RegisterZNoteServiceServer(grpcServer, znoteServer)

	_ = grpcServer.Serve(lis)
}

func executeNoteRequest(ctx context.Context, client pb.ZNoteServiceClient) {
	logrus.WithField("args", os.Args[2:]).Info("params")
	for _, command := range subcommands.Commands {
		if command.Name == os.Args[1] {
			err := command.Flagset.Parse(os.Args[2:])
			if err != nil {
				logrus.WithError(err).Panic("parse flag failure.")
			}
			req := command.NewRequest()
			resp, err := client.ExecuteCommand(ctx, &req)
			if err != nil {
				logrus.WithError(err).WithField("req", req).Panic("Failed to execute request")
			}
			fmt.Print(resp.Result)
			return
		}
	}
	fmt.Printf("%s is not a valid command\n", os.Args[1])
}

func executeCommand() {
	if os.Args[1] == "kill-server" {
		clientCall(killServer)
		return
	}
	clientCall(executeNoteRequest)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	}

	userRoot := os.Getenv("HOME")
	context := daemon.Context{
		PidFileName: filepath.Join(userRoot, "znote.lock"),
		PidFilePerm: 0644,
		LogFileName: filepath.Join(userRoot, "znote.log"),
		LogFilePerm: 0666,
		WorkDir:     userRoot,
	}
	child, _ := context.Search()
	if child != nil {
		logrus.Info("Server has been already serving")
		executeCommand()
		return
	}

	if len(os.Args) >= 2 && os.Args[1] == "kill-server" {
		logrus.Info("server is not running, doing nothing")
		return
	}

	child, err := context.Reborn()
	if err != nil {
		logrus.WithError(err).Panic("failed to reborn")
	}

	if child != nil {
		time.Sleep(time.Second * 1)
		executeCommand()
	} else {
		serve()
		defer func() {
			err := context.Release()
			if err != nil {
				logrus.WithError(err).Warn("Release context failed")
			}
		}()
	}

}
