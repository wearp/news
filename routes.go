package main

import "net/http"

// actions GET, POST, DELETE, etc. can now be specified
// handlers are defined in handlers.goj
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
        "ObservationIndex",
        "GET",
        "/news",
        ObservationIndex,
    },
    Route{
        "ObservationShow",
        "GET",
        "/news/{observationId}",
        ObservationShow,
    },
    Route{
        "ObservationCreate",
        "POST",
        "/news",
        ObservationCreate,
    },  
    Route{
        "ObservationDelete",
        "DELETE",
        "/news/{observationId}",
        ObservationDelete,
    },
}
