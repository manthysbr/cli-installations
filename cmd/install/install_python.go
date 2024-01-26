package install

import (
	"fmt"
	"log"
	"os/exec"
)

func InstallPython() {
	fmt.Println("Python está sendo instalado. Aguarde...")
	
	fmt.Println("Preparando ambiente...")

	cmd := exec.Command("sudo", "apt-get", "install", "-y", "python3")

	if err := executeInstallCommand(cmd); err != nil {
		log.Fatalf("Falha ao instalar o Python: %s", err)
	}

	fmt.Println("Python instalado com sucesso.")
}

func executeInstallCommand(cmd *exec.Cmd) error {
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("falha ao iniciar a instalação: %w", err)
	}

	fmt.Println("Instalando o Python. Isso pode levar alguns instantes...")

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("falha durante a instalação: %w", err)
	}

	return nil
}
