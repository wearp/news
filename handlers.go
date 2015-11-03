package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "strconv"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func ObservationIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(observations); err != nil {
        panic(err)
    }
}

func ObservationShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    observationId := vars["observationId"]
    fmt.Fprintln(w, "Observation show:", observationId)
}

func ObservationCreate(w http.ResponseWriter, r *http.Request) {
    var observation Observation
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &observation); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unproccessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    t := RepoCreateObservation(observation)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}

func ObservationDelete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    observationId := vars["observationId"]
    id, err := strconv.Atoi(observationId)
    if err != nil {
        panic(err)
    }
    if err := RepoDestroyObservation(id); err != nil {
        panic(err)
    }
    fmt.Fprintln(w, "Observation deleted:", observationId)
}

