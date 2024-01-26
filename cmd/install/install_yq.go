package install

import (
    "fmt"
    "os/exec"
    "time"
)

func InstallYq() {
    fmt.Println("YQ está sendo instalado. Aguarde...")

    // Preparando ambiente
    prepararAmbiente()

    // Adiciona o repositório PPA do yq
    addRepoCmd := exec.Command("sudo", "add-apt-repository", "-y", "ppa:rmescandon/yq")
    executeInstallCommand(addRepoCmd)

    // Atualiza os pacotes após adicionar o repositório
    updateCmd := exec.Command("sudo", "apt-get", "update")
    executeInstallCommand(updateCmd)

    // Instala o yq
    installCmd := exec.Command("sudo", "apt-get", "install", "-y", "yq")
    executeInstallCommand(installCmd)

    fmt.Println("YQ instalado com sucesso.")
}