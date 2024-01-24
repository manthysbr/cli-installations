package config

import (
    "encoding/json"
    "os"
)

type SoftwareState struct {
    Software map[string]map[string]string
}

var state SoftwareState

func SaveSoftwareState(software string, status string, version string) {
    if state.Software == nil {
        state.Software = make(map[string]map[string]string)
    }
    state.Software[software] = map[string]string{
        "state":   status,
        "version": version,
    }
    data, _ := json.Marshal(state)
    os.WriteFile("software_state.json", data, 0644)
}

func GetSoftwareState() SoftwareState {
    data, _ := os.ReadFile("software_state.json")
    json.Unmarshal(data, &state)
    return state
}