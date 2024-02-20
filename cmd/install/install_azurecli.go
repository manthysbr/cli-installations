package install

import (
	"fmt"
	"os/exec"
)

// InstallDocker instala o Docker utilizando as funções do common_install.
func InstallAzureCLI() {
	fmt.Println("Iniciando a instalação do Azure CLI...")

    // Preparando ambiente
    prepararAmbiente()

    // Define o comando para instalar o Docker
	cmd := exec.Command("bash", "-c", "curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash")

    // Executa o comando de instalação e exibe um spinner enquanto aguarda a conclusão
    executeInstallCommand(cmd)

    fmt.Println("Azure CLI instalado com sucesso.")
}