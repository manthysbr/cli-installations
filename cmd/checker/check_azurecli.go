package checker

import (
	"os/exec"
	"strings"
	"os"
	"encoding/json"
)

// SoftwareState representa o estado de instalação de um software
type SoftwareState struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  string `json:"status"` // Installed, Not installed, etc.
}

// CheckAzureCLI verifica a instalação do Azure CLI.
func CheckAzureCLI() {
	cmd := exec.Command("az", "--version")
	output, err := cmd.CombinedOutput()
	state := SoftwareState{Name: "AzureCLI"}

	if err != nil {
		state.Status = "Not installed"
	} else {
		state.Status = "Installed"
		// Adapte esta parte para extrair a versão corretamente com base na saída do seu comando
		versionInfo := strings.Split(strings.TrimSpace(string(output)), "\n")
		for _, line := range versionInfo {
			if strings.HasPrefix(line, "azure-cli") {
				parts := strings.Fields(line)
				if len(parts) >= 2 {
					state.Version = parts[1] // assumindo que a versão é o segundo elemento
					break
				}
			}
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
