package main

import "time"

type Event struct {
  Id        int                 `json:"id"`
  FieldData []map[string]string `json:"fieldData"`
  CreatedAt time.Time           `json:"createdAt"`
  UpdatedAt time.Time           `json:"updatedAt"`
}

type Events []Event
