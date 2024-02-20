package install

import (
	"fmt"
	"os/exec"
)

// InstallDocker instala o Docker utilizando as funções do common_install.
func InstallJq() {
    fmt.Println("jq está sendo instalado. Aguarde...")

    // Preparando ambiente
    prepararAmbiente()

    // Define o comando para instalar o Docker
    installCmd := exec.Command("sudo", "apt-get", "install", "-y", "jq")

    // Executa o comando de instalação e exibe um spinner enquanto aguarda a conclusão
    ExecuteInstallCommand(installCmd)

    fmt.Println("jq instalado com sucesso.")
}