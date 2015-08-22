package main

import (
	"path"
	"os"
	"fmt"
	"flag"
	"strings"
	"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

const (
	message string = "static-server is starting on the port: %v"
	listenport string = ":%v"
)

func handler(w http.ResponseWriter, r *http.Request) {

	pathname := strings.Trim(r.URL.Path, "/")
	wd, error := os.Getwd()

	if error != nil {
		log.Fatal(error)
	}

	files, error := ioutil.ReadDir(wd)

	if error != nil {
		log.Fatal(error)
	}

	for _, file := range files {

		filename := file.Name()
		basename := path.Base(filename)

		if strings.Contains(basename, pathname) {

			fullpath := path.Join(wd, filename)
			data, error := ioutil.ReadFile(fullpath)

			if error != nil {
				log.Fatal(error)
				return
			}

			w.Write(data)
			return
		}
	}
}

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)
}

func main() {

	port := flag.Int("port", 8000, "Port Number")
	flag.Parse()

	log.Info(fmt.Sprintf(message, *port))

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(listenport, *port), nil)
}
