package checker

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
)

// SoftwareState representa o estado de instalação de um software
type SoftwareState struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  string `json:"status"` // Installed, Not installed, etc.
}

// CheckDocker verifica a instalação do Docker.
func CheckDocker() {
	cmd := exec.Command("docker", "--version")
	output, err := cmd.CombinedOutput()
	state := SoftwareState{Name: "Docker"}

	if err != nil {
		state.Status = "Not installed"
	} else {
		state.Status = "Installed"
		// A saída do docker --version é geralmente "Docker version x.y.z, build aabbccdd"
		versionParts := strings.Fields(strings.TrimSpace(string(output)))
		if len(versionParts) >= 3 {
			state.Version = versionParts[2] // "x.y.z" é o terceiro elemento
		}
	}

	// Grava o estado no arquivo JSON
	saveSoftwareState(state)
}

// saveSoftwareState grava o estado de um software em um arquivo JSON.
func saveSoftwareState(state SoftwareState) {
	fileName := "software_state.json"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// Tratar erro conforme sua estratégia, por exemplo, logar ou retornar o erro
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(state)
	if err != nil {
		// Tratar erro conforme sua estratégia, por exemplo, logar ou retornar o erro
		return
	}
}
