package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

// creates router with, registering URL paths and handlers. See routes.go
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)

    // see routes.go for Routes and Route types
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc
        // see logger.go
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)

    }
    return router
  
}
