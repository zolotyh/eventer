package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gorilla/mux"
)

type CalendarEvent struct {
    Title string
    Desc  string
}

func main() {
	r := mux.NewRouter()

	tmpl := template.Must(template.ParseFiles("template.ics"))

	r.HandleFunc("/api", func(writter http.ResponseWriter, request *http.Request) {
        log.Printf("%v: Recieved request for index\n", time.Now())
        // writter.Header().Set("Content-Type", "text/calendar")
        writter.Header().Set("Content-Disposition", "inline; filename=\"event.ics\"")

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

        c :=CalendarEvent{Title: title[0], Desc: description[0]}


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
