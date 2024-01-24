package config

import (
    "encoding/json"
    "os"
)

type SoftwareState struct {
    Software map[string]string
}

var state SoftwareState

func SaveSoftwareState(software, status string) {
    if state.Software == nil {
        state.Software = make(map[string]string)
    }
    state.Software[software] = status
    data, _ := json.Marshal(state)
    os.WriteFile("software_state.json", data, 0644)
}

func GetSoftwareState(software string) string {
    data, _ := os.ReadFile("software_state.json")
    json.Unmarshal(data, &state)
    return state.Software[software]
}
