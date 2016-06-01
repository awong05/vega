package main

import (
  "encoding/json"
  "fmt"
  "io"
  "io/ioutil"
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

  w.Header().Set("Content-Type", "application/json;charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(events); err != nil {
    panic(err)
  }
}

func EventShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  eventId := vars["eventId"]
  fmt.Fprintln(w, "Event show:", eventId)
}

func EventCreate(w http.ResponseWriter, r *http.Request) {
  var event Event
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

  if err != nil {
    panic(err)
  }
  if err := r.Body.Close(); err != nil {
    panic(err)
  }
  if err := json.Unmarshal(body, &event); err != nil {
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  //TODO: Database transaction to create event.

  w.Header().Set("Content-Type", "application/json;charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  //if err := json.NewEncoder(w).Encode(t); err != nil {
  //  panic(err)
  //}
}
