package validate

import (
    "fmt"
    "os/exec"
    "strings"
)

func CheckProxy() {
    cmd := exec.Command("proxyman", "--version")
    output, err := cmd.CombinedOutput()

    if err != nil {
        fmt.Println("Proxyman não está instalado.")
        return
    }

    version := strings.TrimSpace(string(output))
    fmt.Printf("Versão do Proxyman encontrada: %s\n", version)
}