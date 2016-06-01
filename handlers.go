package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "time"

  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome!")
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

  if err := json.NewEncoder(w).Encode(events); err != nil {
    panic(err)
  }
}

func EventShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  eventId := vars["eventId"]
  fmt.Fprintln(w, "Event show:", eventId)
}
