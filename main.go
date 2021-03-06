package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

const (
	startingMessage string = "static-server is starting on the port: %v"
	listenPort      string = ":%v"
)

func handler(w http.ResponseWriter, r *http.Request) {

	pathname := strings.Trim(r.URL.Path, "/")
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	if pathname == "" {
		ShowDir(w, wd)
		return
	}

	wd = fmt.Sprintf("%s/%s", wd, pathname)
	file, err := os.Open(wd)

	if err != nil {
		NotFound(w)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		NotFound(w)
		return
	}

	if fileInfo.IsDir() {
		ShowDir(w, wd)
		return
	}

	ShowFile(w, file)
	return
}

func ShowFile(w http.ResponseWriter, file *os.File) {

	data, err := ioutil.ReadAll(file)

	if err != nil {
		NotFound(w)
		return
	}
	w.Write(data)
}

func ShowDir(w http.ResponseWriter, dir string) {

	fileInfos, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	list := make([]string, len(fileInfos))

	for i, file := range fileInfos {
		filename := file.Name()
		basename := path.Base(filename)
		if file.IsDir() {
			basename += "/"
		}
		list[i] = basename
	}

	data := struct {
		Title string
		Items []string
	}{
		Title: fmt.Sprintf("%s/", dir),
		Items: list,
	}

	directoryTemplate.ExecuteTemplate(w, "html", data)
}

func NotFound(w http.ResponseWriter) {

	w.WriteHeader(404)
	notFoundTemplate.ExecuteTemplate(w, "html", nil)
}

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)
}

func main() {

	port := flag.Int("port", 8000, "Port Number")
	flag.Parse()

	log.Info(fmt.Sprintf(startingMessage, *port))

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(listenPort, *port), nil)
}
