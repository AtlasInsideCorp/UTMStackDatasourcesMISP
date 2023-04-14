package serv

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"regexp"
	"time"

	"github.com/KbaYero/UTMStackDatasourcesMISP/config"
)

type FileServerData struct {
	Data template.HTML
	Time string
}

func ShowEvents(cnf config.Configuration) {
	// Handling page styles
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir(filepath.Join("serv", "static")))))

	fs := rootPath(http.FileServer(http.Dir(cnf.EventsPath)))

	// Serve static files
	http.Handle("/", fs)
	fmt.Println("Serving on port ", cnf.ServerPort)
	fmt.Println(http.ListenAndServe(":"+cnf.ServerPort, nil))
}

func rootPath(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// add header(s)
		w.Header().Set("Cache-Control", "no-cache")
		checkRegex := regexp.MustCompile(`(.+)?misp-instance-(.+)\.json`)
		match := checkRegex.Match([]byte(r.URL.Path))

		if !match {
			fsdata := FileServerData{"", time.Now().Format(time.Stamp)}
			templates := template.Must(template.ParseFiles("serv/templates/fs-index.html"))

			if err := templates.ExecuteTemplate(w, "fs-index.html", fsdata); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		}
		h.ServeHTTP(w, r)
	})
}
