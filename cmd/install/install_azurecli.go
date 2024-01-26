package install

import (
	"fmt"
	"log"
	"os/exec"
)

func InstallAzureCLI() {
	fmt.Println("Iniciando a instalação do Azure CLI...")

	// Download e execução do script de instalação da Azure CLI
	cmd := exec.Command("bash", "-c", "curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash")

	// Inicia o comando sem esperar que ele termine
	if err := cmd.Start(); err != nil {
		log.Fatalf("Falha ao iniciar a instalação do Azure CLI: %s", err)
	}

	fmt.Println("Instalando o Azure CLI. Isso pode levar alguns instantes...")

	// Espera o comando terminar
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Falha ao instalar o Azure CLI: %s", err)
	} else {
		fmt.Println("Azure CLI instalado com sucesso.")
	}
}
