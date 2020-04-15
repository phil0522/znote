package subcommands

import (
	"bufio"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/phil0522/znote/pkg/notesmarket"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"

	_ "github.com/phil0522/znote/statik"
)

var (
	ServeCommandFlagSet = flag.NewFlagSet("Edit", flag.ExitOnError)
)

func ServeHttp() {
	logrus.Info("Serve HTTP")

	if err := os.Chdir(notesmarket.RootDir); err != nil {
		logrus.WithField("md-root", notesmarket.RootDir).Panic("can not change directory")
	}

	box, err := fs.New()
	if err != nil {
		logrus.Panic("failed to create statik folder")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logrus.WithField("path", r.URL.Path).Debug("Get request")
		path := r.URL.Path
		if path == "/" {
			path = "/index.html"
		}
		file, err := box.Open(r.URL.Path)
		if err != nil {
			logrus.WithField("path", r.URL.Path).Debug("not exist, use html.")
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
	if err := http.ListenAndServe(":3000", nil); err != nil {
		logrus.Panic("can not serve")
	}
}
