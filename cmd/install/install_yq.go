package install

import (
    "fmt"
    "os/exec"
)

func InstallYq() {
    fmt.Println("Iniciando a instalação do yq...")
    cmd := exec.Command("sudo", "apt-get", "install", "-y", "yq")
    executeInstallCommand(cmd)
}
