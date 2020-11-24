package main

import (
    "net/http"
)

func init() {
    fs := http.FileServer(http.Dir("web"))
    http.Handle("/", fs)
}