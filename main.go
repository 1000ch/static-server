package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

const (
	message              string = "static-server is starting on the port: %v"
	listenport           string = ":%v"
	baseTemplatePath     string = "template/base.html"
	showDirTemplatePath  string = "template/dir.html"
	notFoundTemplatePath string = "template/404.html"
)

var showDirTemplate *template.Template
var notFoundTemplate *template.Template

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

func ShowDir(w http.ResponseWriter, wd string) {
	fileInfos, err := ioutil.ReadDir(wd)
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
		Title: fmt.Sprintf("Directory listing for %s/", wd),
		Items: list,
	}
	showDirTemplate.ExecuteTemplate(w, "base", data)
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(404)
	notFoundTemplate.ExecuteTemplate(w, "base", nil)
}

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)

	showDirTemplate, _ = template.ParseFiles(baseTemplatePath, showDirTemplatePath)
	notFoundTemplate, _ = template.ParseFiles(baseTemplatePath, notFoundTemplatePath)
}

func main() {

	port := flag.Int("port", 8000, "Port Number")
	flag.Parse()

	log.Info(fmt.Sprintf(message, *port))

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(listenport, *port), nil)
}
