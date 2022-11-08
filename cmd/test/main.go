package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmplt *template.Template

type News struct {
	Headline string
	Body     string
}

func handlePage(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		tmplt, _ = template.ParseFiles("/home/oem/go/src/github.com/viktar3w/web-example/cmd/test/test.tmpl")

		event := News{
			Headline: "makeuseof.com has everything Tech",
			Body:     "Visit MUO for anything technology related",
		}

		err := tmplt.Execute(writer, event)

		if err != nil {
			return
		}
	}
}
func main() {
	runServer()
}
func runServer() {
	http.HandleFunc("/", handlePage)
	err := http.ListenAndServe(":23423", nil)

	if err != nil {
		log.Fatalln("There's an error with the server:", err)
	}
}
