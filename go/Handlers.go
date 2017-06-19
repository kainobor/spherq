package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "time"
    "io/ioutil"
    "io"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func ImageIndex(w http.ResponseWriter, r *http.Request) {
    images := Images{
        Image{Id: 1, Name: "Image one", URL: "", CreatedAt: time.Now()},
        Image{Id: 2, Name: "Image two", URL: "", CreatedAt: time.Now()},
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(images); err != nil {
        panic(err)
    }
}

func ImageShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    fmt.Sprintln(w, "Image show: %1", json.NewEncoder(w).Encode(vars))
}

func ImageCreate(w http.ResponseWriter, r *http.Request) {
    var image Image
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

    if err != nil {
        panic(err)
    }

    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    if err := json.Unmarshal(body, &image); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    t := RepoCreateImage(image)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}
