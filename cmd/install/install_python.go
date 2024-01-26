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