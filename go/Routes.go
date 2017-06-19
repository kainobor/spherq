package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "ImageIndex",
        "GET",
        "/images",
        ImageIndex,
    },
    Route{
        "ImageShow",
        "GET",
        "/images/{imageId}",
        ImageShow,
    },
    Route{
        "ImageCreate",
        "POST",
        "/images",
        ImageCreate,
    },
}
