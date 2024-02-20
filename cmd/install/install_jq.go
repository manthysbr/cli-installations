
package install

import (
	"fmt"
	"os/exec"
)

// InstallDocker instala o Docker utilizando as funções do common_install.
func InstallJQ() {
    fmt.Println("jq está sendo instalado. Aguarde...")

    // Preparando ambiente
    prepararAmbiente()

    // Define o comando para instalar o Docker
    installCmd := exec.Command("sudo", "apt-get", "install", "-y", "jq")

    // Executa o comando de instalação e exibe um spinner enquanto aguarda a conclusão
    executeInstallCommand(cmd)

    fmt.Println("jq instalado com sucesso.")
}