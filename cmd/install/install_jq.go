package install

import (
	"fmt"
	"os/exec"
)

// Funcao de instalacao do common_install.sh para o jq (https://stedolan.github.io/jq/)
func InstallJq() {
    fmt.Println("jq está sendo instalado. Aguarde...")

    // Preparando ambiente
    prepararAmbiente()

    // Define o comando para instalar o Docker
    installCmd := exec.Command("sudo", "apt-get", "install", "-y", "jq")

    // Executa o comando de instalação e exibe um spinner enquanto aguarda a conclusão
    executeInstallCommand(installCmd)

    fmt.Println("jq instalado com sucesso.")
}