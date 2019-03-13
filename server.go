package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/lestrrat-go/ical"
)

func getCalendarFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v: Recieved request for index\n", time.Now())

	w.Header().Set("Content-Type", "text/calendar")
	w.Header().Set("Content-Disposition", "inline; filename=\"event.ics\"")

	c := ical.New()
	c.AddProperty("X-Foo-Bar-Baz", "value")
	tz := ical.NewTimezone()
	tz.AddProperty("TZID", "Asia/Tokyo")
	c.AddEntry(tz)

	ical.NewEncoder(os.Stdout).Encode(c)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getCalendarFile)

	var port string

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}
