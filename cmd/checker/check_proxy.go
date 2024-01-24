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

// CheckProxyman verifica a instalação do Proxyman.
func CheckProxyman() {
	cmd := exec.Command("proxyman", "--version") // Substitua pelo comando correto se necessário
	output, err := cmd.CombinedOutput()
	state := SoftwareState{Name: "Proxyman"}

	if err != nil {
		state.Status = "Not installed"
	} else {
		state.Status = "Installed"
		version := strings.TrimSpace(string(output))
		state.Version = version // Ajuste conforme o formato da saída do seu comando
	}

	// Grava o estado no arquivo JSON
	saveSoftwareState(state)
}

// saveSoftwareState grava o estado de um software em um arquivo JSON.
func saveSoftwareState(state SoftwareState) {
	fileName := "software_state.json"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// Tratar erro conforme sua estratégia
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(state)
	if err != nil {
		// Tratar erro conforme sua estratégia
		return
	}
}
