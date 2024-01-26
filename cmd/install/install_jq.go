package install

import (
    "fmt"
    "os/exec"
)

func InstallJq() {
    fmt.Println("Iniciando a instalação do jq...")
    cmd := exec.Command("sudo", "apt-get", "install", "-y", "jq")
    executeInstallCommand(cmd)
}
