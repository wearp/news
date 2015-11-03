package main 

import "fmt"

var currentId int

var observations Observations

// Give us some seed data
func init() {
    RepoCreateObservation(
        Observation{
            PatientId: 1, UserId: 3, AVPU: "A", HeartRate: 50,
            RespiratoryRate: 21, SystolicBP: 34, O2Saturation: 93,
            O2Supplement: false, Temperature: 67, Completed: true,
        })
    RepoCreateObservation(
        Observation{
            PatientId: 1, UserId: 3, AVPU: "A", HeartRate: 50,
            RespiratoryRate: 21, SystolicBP: 100, O2Saturation: 93,
            O2Supplement: false, Temperature: 67, Completed: true,
        })
}

func RepoFindObservation(id int) Observation {
    for _, t := range observations {
        if t.Id == id {
            return t
        }
    }
    // return empty Observation if not found
    return Observation{}
}

func RepoCreateObservation(t Observation) Observation{
    currentId += 1
    t.Id = currentId
    t.calculateRisk()
    observations = append(observations, t)
    return t
}

func RepoDestroyObservation(id int) error {
    for i, t := range observations {
        if t.Id == id {
            observations = append(observations[:i], observations[i+1:]...)
            return nil
        }
    }

    return fmt.Errorf("Could not find Observation with id of %d to delete", id)
}

