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

// CheckPython verifica a instalação do Python 3 e grava o estado em um arquivo JSON.
func CheckPython() {
    cmd := exec.Command("python3", "--version")
    output, err := cmd.CombinedOutput()
    state := SoftwareState{Name: "Python"}

    if err != nil {
        state.Status = "Not installed"
        state.Version = ""
    } else {
        state.Status = "Installed"
        version := strings.TrimSpace(string(output))
        // Extrai apenas a versão numérica se a saída for "Python 3.x.y"
        if strings.HasPrefix(version, "Python ") {
            state.Version = strings.TrimPrefix(version, "Python ")
        } else {
            state.Version = version
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
