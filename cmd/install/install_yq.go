package install

import (
	"fmt"
	//"log"
	"os/exec"
	//"time"
)

func InstallYq() {
	fmt.Println("YQ está sendo instalado. Aguarde...")

	// Preparando ambiente
	prepararAmbiente()

	// Adiciona o PPA para o yq
	addCmd := exec.Command("sudo", "add-apt-repository", "-y", "ppa:rmescandon/yq")
	executeInstallCommand(addCmd, "Adicionando o repositório PPA para YQ...")

	// Atualiza os repositórios
	updateCmd := exec.Command("sudo", "apt", "update")
	executeInstallCommand(updateCmd, "Atualizando os repositórios...")

	// Instala o yq
	installCmd := exec.Command("sudo", "apt-get", "install", "-y", "yq")
	executeInstallCommand(installCmd, "Instalando o YQ. Isso pode levar alguns instantes...")
}
