package install

import (
	"fmt"
	"log"
	"os/exec"
)

func InstallGit() {
	fmt.Println("Git está sendo instalado. Aguarde...")
	
	fmt.Println("Preparando ambiente...")

	cmd := exec.Command("sudo", "apt-get", "install", "-y", "git")

	if err := executeInstallCommand(cmd); err != nil {
		log.Fatalf("Falha ao instalar o Git: %s", err)
	}

	fmt.Println("Git instalado com sucesso.")
}

func executeInstallCommand(cmd *exec.Cmd) error {
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("falha ao iniciar a instalação: %w", err)
	}

	fmt.Println("Instalando o Git. Isso pode levar alguns instantes...")

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("falha durante a instalação: %w", err)
	}

	return nil
}
