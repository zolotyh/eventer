package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	// "github.com/lestrrat-go/ical"
)

type CalendarEvent struct {
    Title string
    Desc  string
}

func main() {
	r := mux.NewRouter()

	tmpl := template.Must(template.ParseFiles("template.ics"))

	r.HandleFunc("/", func(writter http.ResponseWriter, request *http.Request) {
        log.Printf("%v: Recieved request for index\n", time.Now())
        // w.Header().Set("Content-Type", "text/calendar")
        writter.Header().Set("Content-Disposition", "inline; filename=\"event.ics\"")
        c :=CalendarEvent{Title: "title", Desc: "desc"}
        tmpl.Execute(writter, c)
	})

	var port string

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}
