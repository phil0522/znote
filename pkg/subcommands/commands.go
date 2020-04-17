package subcommands

import (
	"flag"

	pb "github.com/phil0522/znote/proto"
)

type SubCommand struct {
	Name           string
	Flagset        *flag.FlagSet
	NewRequest     func() pb.ZNoteRequest
	ResolveRequest func(pb.ZNoteRequest) pb.ZNoteResponse
}

func makeSubCommand(name string, flagset *flag.FlagSet, newRequest func() pb.ZNoteRequest, resolveRequest func(pb.ZNoteRequest) pb.ZNoteResponse) SubCommand {
	return SubCommand{
		Name:           name,
		Flagset:        flagset,
		NewRequest:     newRequest,
		ResolveRequest: resolveRequest,
	}
}

var (
	Commands = []SubCommand{
		makeSubCommand("new", NewNoteCommandFlagSet, NewNoteCreateRequest, ResolveNewNote),
		makeSubCommand("list", ListNoteCommandFlagSet, CreateListNoteRequest, ResolveListNote),
		makeSubCommand("edit", EditNoteCommandFlagSet, CreateEditNoteRequest, ResolveEditNote),
		makeSubCommand("pick", PickNoteCommandFlagSet, CreatePickNoteRequest, ResolvePickNote),
		makeSubCommand("serve", ServeCommandFlagSet, CreateServeRequest, ResolveServe),
		makeSubCommand("refresh", RefreshCommandFlagSet, CreateRefreshMarketRequest, ResolveRefreshMarket),
		makeSubCommand("show", ShowNoteCommandFlagSet, CreateShowNoteRequest, ResolveShowNote),
		makeSubCommand("preview", ShowNoteCommandFlagSet, CreateShowNoteRequest, ResolveShowNote),
	}
)
