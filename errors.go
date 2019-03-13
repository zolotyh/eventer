package main

import (
    "log"
    "net/http"
)

type Error struct {

}

func (e Error) notFoundError(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text-html")
    w.WriteHeader(http.StatusNotFound)

    log.Printf("404! Not found", w)
}
