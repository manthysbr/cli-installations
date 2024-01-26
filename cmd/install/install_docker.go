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
	if err := executeInstallCommand(cmd); err != nil {
		log.Fatalf("Falha ao instalar o Docker: %s", err)
	}

	fmt.Println("Docker instalado com sucesso.")
}

// executeInstallCommand executa o comando e retorna o erro, se houver
func executeInstallCommand(cmd *exec.Cmd) error {
	// Inicia o comando sem esperar que ele termine
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("falha ao iniciar a instalação: %w", err)
	}

	fmt.Println("Instalando o Docker. Isso pode levar alguns instantes...")

	// Espera o comando terminar e verifica se houve erro
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("falha durante a instalação: %w", err)
	}

	return nil
}
