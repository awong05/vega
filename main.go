package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Event struct {
	FieldData []map[string]string `json:"fieldData"`
	CreatedAt time.Time           `json:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt"`
}

type Events []Event

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/events", EventIndex)
	router.HandleFunc("/events/{eventId}", EventShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func EventIndex(w http.ResponseWriter, r *http.Request) {
	events := Events{
		Event{
			FieldData: []map[string]string{
				map[string]string{
					"email":    "go@lang.com",
					"question": "What is the meaning of life?",
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Event{
			FieldData: []map[string]string{
				map[string]string{
					"full_name": "Golang Gopher",
					"answer":    "42",
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	json.NewEncoder(w).Encode(events)
}

func EventShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]
	fmt.Fprintln(w, "Event show:", eventId)
}
