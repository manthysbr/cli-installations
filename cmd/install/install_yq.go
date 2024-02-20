package install

import (
	"fmt"
	"os/exec"
)

// InstallDocker instala o Docker utilizando as funções do common_install.
func InstallYQ() {
    fmt.Println("yq está sendo instalado. Aguarde...")

    // Preparando ambiente
    prepararAmbiente()

    // Define o comando para instalar o Docker
    addRepoCmd := exec.Command("sudo", "add-apt-repository", "-y", "ppa:rmescandon/yq")
    executeInstallCommand(addRepoCmd)

    updateCmd := exec.Command("sudo", "apt-get", "update")
    executeInstallCommand(updateCmd)

    installCmd := exec.Command("sudo", "apt-get", "install", "-y", "yq")
    executeInstallCommand(installCmd)

    fmt.Println("yq instalado com sucesso.")
}