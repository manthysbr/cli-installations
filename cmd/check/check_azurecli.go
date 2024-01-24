package validate

import (
    "fmt"
    "os/exec"
    "strings"
)

func CheckAzureCli() {
    cmd := exec.Command("az", "--version")
    output, err := cmd.CombinedOutput()

    if err != nil {
        fmt.Println("Azure CLI não está instalado.")
        return
    }

    version := strings.TrimSpace(string(output))
    fmt.Printf("Versão do Azure CLI encontrada: %s\n", version)
}