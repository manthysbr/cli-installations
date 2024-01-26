package install

import (
    "fmt"
    "os/exec"
)

func InstallGit() {
    fmt.Println("Git está sendo instalado. Aguarde...")
    
    fmt.Println("Preparando ambiente...")

    cmd := exec.Command("sudo", "apt-get", "install", "-y", "git")

    // Executa o comando de instalação
    executeInstallCommand(cmd)

    fmt.Println("Git instalado com sucesso.")
}