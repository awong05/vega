package main

import "time"

type Event struct {
  FieldData []map[string]string `json:"fieldData"`
  CreatedAt time.Time           `json:"createdAt"`
  UpdatedAt time.Time           `json:"updatedAt"`
}

type Events []Event
