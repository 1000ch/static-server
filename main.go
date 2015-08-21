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
	Port string = ":8080"
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

func main() {

	port := flag.Int("port", 8000, "Port Number")
	flag.Parse()

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
}
