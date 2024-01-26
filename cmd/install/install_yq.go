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

func executeCommand(cmd *exec.Cmd, message string) {
	fmt.Println(message)

	// Inicia o comando sem esperar que ele termine
	if err := cmd.Start(); err != nil {
		log.Fatalf("Falha ao executar o comando: %s", err)
	}

	// Espera o comando terminar
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Falha durante a execução: %s", err)
	} else {
		fmt.Printf("%s instalado com sucesso.\n", cmd.Args[len(cmd.Args)-1])
	}
}
