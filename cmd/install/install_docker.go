package install

import (
	"fmt"
	"os/exec"
)

// InstallDocker instala o Docker utilizando as funções do common_install.
func InstallDocker() {
    fmt.Println("Docker está sendo instalado. Aguarde...")

    // Preparando ambiente
    prepararAmbiente()

    // Define o comando para instalar o Docker
    cmd := exec.Command("sudo", "apt-get", "install", "-y", "docker.io")

    // Executa o comando de instalação e exibe um spinner enquanto aguarda a conclusão
    executeInstallCommand(cmd)

    fmt.Println("Docker instalado com sucesso.")
}