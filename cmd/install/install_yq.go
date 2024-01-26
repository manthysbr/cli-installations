package install

import (
    "fmt"
    "os/exec"
)

func InstallYq() {
    fmt.Println("YQ está sendo instalado. Aguarde...")

    // Preparando ambiente
    prepararAmbiente()

    // Adiciona o PPA para o yq
    addCmd := exec.Command("sudo", "add-apt-repository", "-y", "ppa:rmescandon/yq")
    executeInstallCommand(addCmd)

    // Atualiza os repositórios
    updateCmd := exec.Command("sudo", "apt", "update")
    executeInstallCommand(updateCmd)

    // Instala o yq
    installCmd := exec.Command("sudo", "apt-get", "install", "-y", "yq")
    executeInstallCommand(installCmd)
}