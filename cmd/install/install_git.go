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
