package main

import (
    "log"
    "net/http"
)

func main() {

    router := NewRouter()

    // router listening for requests on port 8080
    log.Fatal(http.ListenAndServe(":8080", router))
}
