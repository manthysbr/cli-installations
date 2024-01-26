package install

import (
	"fmt"
	"log"
	"os/exec"
)

func InstallDocker() {
    fmt.Println("Docker está sendo instalado. Aguarde...")

    // Preparando ambiente
    fmt.Println("Preparando ambiente...")

    cmd := exec.Command("sudo", "apt-get", "install", "-y", "docker.io")

    // Executa o comando de instalação e captura qualquer erro
    err := executeInstallCommand(cmd)
    if err != nil {
        log.Fatalf("Falha ao instalar o Docker: %s", err)
    }

    fmt.Println("Docker instalado com sucesso.")
}
