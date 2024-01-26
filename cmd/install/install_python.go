package install

import (
    "fmt"
    "os/exec"
)

func InstallPython() {
    fmt.Println("Python está sendo instalado. Aguarde...")
    
    fmt.Println("Preparando ambiente...")

    cmd := exec.Command("sudo", "apt-get", "install", "-y", "python3")

    // Executa o comando de instalação
    executeInstallCommand(cmd)

    fmt.Println("Python instalado com sucesso.")
}