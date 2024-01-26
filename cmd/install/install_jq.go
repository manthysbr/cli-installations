package install

import (
    "fmt"
    "os/exec"
)

func InstallJq() {
    fmt.Println("JQ está sendo instalado. Aguarde...")

    // Preparando ambiente
    prepararAmbiente()

    // Instala o jq
    installCmd := exec.Command("sudo", "apt-get", "install", "-y", "jq")

    // Executa o comando de instalação
    executeInstallCommand(installCmd)

    fmt.Println("JQ instalado com sucesso.")
}
