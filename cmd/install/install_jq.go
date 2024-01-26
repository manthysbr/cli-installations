package install

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func InstallJq() {
	fmt.Println("JQ está sendo instalado. Aguarde...")

	// Preparando ambiente
	prepararAmbiente()

	// Instala o jq
	installCmd := exec.Command("sudo", "apt-get", "install", "-y", "jq")
	executeInstallCommand(installCmd, "Instalando o JQ. Isso pode levar alguns instantes...")
}

func prepararAmbiente() {
	// Simula o tempo de preparação do ambiente
	time.Sleep(2 * time.Second)
	fmt.Println("Preparando ambiente...")
}
