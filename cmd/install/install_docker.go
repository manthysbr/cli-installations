package install

import (
	"fmt"
	"os/exec"
)

func InstallDocker() {
    fmt.Println("Docker está sendo instalado. Aguarde...")

    // Preparando ambiente
    fmt.Println("Preparando ambiente...")

    cmd := exec.Command("sudo", "apt-get", "install", "-y", "docker.io")

    // Executa o comando de instalação
    executeInstallCommand(cmd)

    fmt.Println("Docker instalado com sucesso.")
}
