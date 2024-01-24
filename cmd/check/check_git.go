package validate

import (
    "fmt"
    "os/exec"
    "strings"
)

func CheckGit() {
    cmd := exec.Command("git", "--version")
    output, err := cmd.CombinedOutput()

    if err != nil {
        fmt.Println("Git não está instalado.")
        return
    }

    version := strings.TrimSpace(string(output))
    fmt.Printf("Versão do Git encontrada: %s\n", version)
}