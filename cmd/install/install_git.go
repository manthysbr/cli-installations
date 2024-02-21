package install

import (
	"fmt"
	"os/exec"
)

// Funcao de instalacao do common_install.sh para o git
func InstallGit() {
	fmt.Println("Iniciando a instalação do Git...")

    // Preparando ambiente
    prepararAmbiente()

    // Define o comando para instalar o Docker
	cmd := exec.Command("sudo", "apt-get", "install", "-y", "git")

    // Executa o comando de instalação e exibe um spinner enquanto aguarda a conclusão
    executeInstallCommand(cmd)

    fmt.Println("Git instalado com sucesso.")
}

