package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

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
	var port string
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "9000"
	}

	handler := &RegexpHandler{}

	handler.HandleFunc(regexp.MustCompile("^/"), getCalendarFile)
	fmt.Printf("Server is running on port %v\n", port)
	http.ListenAndServe(":"+port, handler)
}
