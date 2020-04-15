package subcommands

import (
	"bufio"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/phil0522/znote/pkg/notesmarket"
	pb "github.com/phil0522/znote/proto"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	_ "github.com/phil0522/znote/statik"
)

var (
	ServeCommandFlagSet = flag.NewFlagSet("Edit", flag.ExitOnError)
)

func CreateServeRequest() pb.ZNoteRequest {
	return pb.ZNoteRequest{
		Command: "serve",
	}
}

var (
	serverStarted              = false
	httpServer    *http.Server = nil
)

func ShutDownHttpServer() {
	if httpServer != nil {
		_ = httpServer.Shutdown(context.Background())
		httpServer = nil
	}
}

func ResolveServe(req pb.ZNoteRequest) pb.ZNoteResponse {
	if serverStarted {
		return pb.ZNoteResponse{
			Result: "Server is already started",
		}
	}
	go ServeHttp()

	return pb.ZNoteResponse{
		Result: "server is starting.",
	}
}

func ServeHttp() {
	logrus.Info("Serve HTTP")
	serverStarted = true

	if err := os.Chdir(notesmarket.RootDir); err != nil {
		logrus.WithField("md-root", notesmarket.RootDir).Panic("can not change directory")
	}

	box, err := fs.New()
	if err != nil {
		logrus.Panic("failed to create statik folder")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path
		if path == "/" {
			path = "/index.html"
		}
		logrus.WithField("path", path).Debug("Get request")
		file, err := box.Open(path)
		if err != nil {
			logrus.WithField("path", path).Debug("not exist, use html.")
			file, err = box.Open("/index.html")
			path = "/index.html"

			if err != nil {
				w.WriteHeader(404)
				_, _ = w.Write([]byte(err.Error()))
				return
			}
		}

		content, err := ioutil.ReadAll(bufio.NewReader(file))
		if err != nil {
			logrus.WithField("path", r.URL.Path).Panic("failed to read")
		}

		contentType := "text/html"
		if strings.HasSuffix(path, "js") {
			contentType = "text/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		}
		w.Header().Add("Content-Type", contentType)
		_, _ = w.Write(content)
	})
	http.Handle("/md/", http.StripPrefix("/md/", http.FileServer(http.Dir(notesmarket.RootDir))))
	httpServer = &http.Server{Addr: ":3000", Handler: nil}

	if err := httpServer.ListenAndServe(); err != nil {
		logrus.Panic("can not serve")
	}
}
