package install

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func InstallDocker() {
	fmt.Println("Docker está sendo instalado. Aguarde...")
	
	// Preparando ambiente
	prepararAmbiente()

	cmd := exec.Command("sudo", "apt-get", "install", "-y", "docker.io")

	// Inicia o comando sem esperar que ele termine
	if err := cmd.Start(); err != nil {
		log.Fatalf("Falha ao iniciar a instalação do Docker: %s", err)
	}

	fmt.Println("Instalando o Docker. Isso pode levar alguns instantes...")

	// Espera o comando terminar
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Falha ao instalar o Docker: %s", err)
	} else {
		fmt.Println("Docker instalado. Prosseguindo instalação.")
	}
}

func prepararAmbiente() {
	// Simula o tempo de preparação do ambiente
	time.Sleep(2 * time.Second)
	fmt.Println("Preparando ambiente...")
}
