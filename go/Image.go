package main

import (
    "time"
)

type Image struct {
    Id          int             `json:"id"`
    Name        string          `json:"name"`
    URL         string          `json:"url"`
    CreatedAt   time.Time       `json:"createdAt"`
}

type Images []Image
