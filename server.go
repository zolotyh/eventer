package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
)

type CalendarEvent struct {
	Title         string
	Desc          string
	TZID          string
	StartDateTime string
	EndDateTime   string
}

func parseValuesFromRequest(request *http.Request) CalendarEvent {
	return CalendarEvent{
		Title:         request.URL.Query()["title"][0],
		Desc:          request.URL.Query()["desc"][0],
		TZID:          request.URL.Query()["tzid"][0],
		StartDateTime: request.URL.Query()["start-date-time"][0],
		EndDateTime:   request.URL.Query()["end-date-time"][0],
	}
}

func main() {
	r := mux.NewRouter()

	box := packr.NewBox("./templates")
	templateRaw, templateRawError := box.FindString("event.ics")

	if templateRawError != nil {
		log.Println("template raw error")
		return
	}

	tmpl := template.Must(template.New("event").Parse(templateRaw))

	r.HandleFunc("/", func(writter http.ResponseWriter, request *http.Request) {
		log.Printf("%v: Recieved request for index\n", time.Now())

		event := parseValuesFromRequest(request)

		buf := &bytes.Buffer{}

		if err := tmpl.Execute(buf, event); err != nil {
			http.Error(writter, "Hey, Request was bad!", http.StatusBadRequest) // HTTP 400 status
			panic(err)
		}

		writter.Header().Set("Content-Disposition", "inline; filename=\"event.ics\"")
		buf.WriteTo(writter)
	})

	var port string

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}
