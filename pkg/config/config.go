package config

import (
    "encoding/json"
    "log"
    "os"
)

type SoftwareState struct {
    Software map[string]map[string]string
}

var state SoftwareState

// SaveSoftwareState salva o estado de um software especificado no arquivo JSON.
func SaveSoftwareState(software, status, version string) {
    if state.Software == nil {
        state.Software = make(map[string]map[string]string)
    }

    state.Software[software] = map[string]string{
        "state":   status,
        "version": version,
    }

    data, err := json.Marshal(state)
    if err != nil {
        log.Fatalf("Erro ao serializar o estado do software: %s", err)
    }

    err = os.WriteFile("/tmp/software_state.json", data, 0644)
    if err != nil {
        log.Fatalf("Erro ao escrever no arquivo software_state.json: %s", err)
    }
}

// GetSoftwareState retorna o estado atual de todos os softwares a partir do arquivo JSON.
func GetSoftwareState() SoftwareState {
    data, err := os.ReadFile("/tmp/software_state.json")
    if err != nil {
        log.Fatalf("Erro ao ler o arquivo software_state.json: %s", err)
    }

    err = json.Unmarshal(data, &state)
    if err != nil {
        log.Fatalf("Erro ao deserializar os dados do software_state.json: %s", err)
    }

    return state
}