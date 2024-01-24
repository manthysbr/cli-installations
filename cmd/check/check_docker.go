package check

import (
    "fmt"
    "os/exec"
    "strings"
)

func CheckDocker() {
    cmd := exec.Command("docker", "--version")
    output, err := cmd.CombinedOutput()

    if err != nil {
        fmt.Println("Docker não está instalado.")
        return
    }

    version := strings.TrimSpace(string(output))
    fmt.Printf("Versão do Docker encontrada: %s\n", version)
}