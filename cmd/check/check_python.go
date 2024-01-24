package validate

import (
    "fmt"
    "os/exec"
    "strings"
)

func CheckPython() {
    cmd := exec.Command("python", "--version")
    output, err := cmd.CombinedOutput()

    if err != nil {
        fmt.Println("Python não está instalado.")
        return
    }

    version := strings.TrimSpace(string(output))
    fmt.Printf("Versão do Python encontrada: %s\n", version)
}