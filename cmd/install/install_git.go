package install

import (
	"fmt"
	"os/exec"
)

// InstallDocker instala o Docker utilizando as funções do common_install.
func InstallGit() {
	fmt.Println("Iniciando a instalação do Azure CLI...")

    // Preparando ambiente
    prepararAmbiente()

    // Define o comando para instalar o Docker
	cmd := exec.Command("sudo", "apt-get", "install", "-y", "git")

    // Executa o comando de instalação e exibe um spinner enquanto aguarda a conclusão
    executeInstallCommand(cmd)

    fmt.Println("Azure CLI instalado com sucesso.")
}

