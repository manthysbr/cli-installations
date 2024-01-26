package install

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func InstallYq() {
	fmt.Println("YQ está sendo instalado. Aguarde...")

	// Preparando ambiente
	prepararAmbiente()

	// Adiciona o PPA para o yq
	addCmd := exec.Command("sudo", "add-apt-repository", "-y", "ppa:rmescandon/yq")
	executeCommand(addCmd, "Adicionando o repositório PPA para YQ...")

	// Atualiza os repositórios
	updateCmd := exec.Command("sudo", "apt", "update")
	executeCommand(updateCmd, "Atualizando os repositórios...")

	// Instala o yq
	installCmd := exec.Command("sudo", "apt-get", "install", "-y", "yq")
	executeCommand(installCmd, "Instalando o YQ. Isso pode levar alguns instantes...")
}

func prepararAmbiente() {
	// Simula o tempo de preparação do ambiente
	time.Sleep(2 * time.Second)
	fmt.Println("Preparando ambiente...")
}
