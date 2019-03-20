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
	Title     string
	Desc      string
	TZID      string
	StartDate string
	EndDate   string
	StartTime string
	// EndTime   string
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

	r.HandleFunc("/api", func(writter http.ResponseWriter, request *http.Request) {
		log.Printf("%v: Recieved request for index\n", time.Now())

		description, descriptionOK := request.URL.Query()["description"]

		if !descriptionOK || len(description[0]) < 1 {
			log.Println("Url Param 'description' is missing")
			return
		}

		title, titleOK := request.URL.Query()["description"]

		if !titleOK || len(title[0]) < 1 {
			log.Println("Url Param 'title' is missing")
			return
		}

		c := CalendarEvent{Title: title[0], Desc: description[0]}

		buf := &bytes.Buffer{}

		if err := tmpl.Execute(buf, c); err != nil {
			http.Error(writter, "Hey, Request was bad!", http.StatusBadRequest) // HTTP 400 status
			panic(err)
		}

		writter.Header().Set("Content-Type", "text/calendar")
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
