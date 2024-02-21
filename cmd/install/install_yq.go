package install

import (
	"fmt"
	"os/exec"
)

// InstallYq instala o yq utilizando wget para baixar o binário diretamente.
func InstallYq() {
    fmt.Println("yq está sendo instalado. Aguarde...")

    // Preparando ambiente
    prepararAmbiente()

    // Define o comando para baixar e instalar o yq
    cmd := exec.Command("bash", "-c", "wget https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64 -O /usr/bin/yq && chmod +x /usr/bin/yq")

    // Executa o comando de instalação
    //
    executeInstallCommand(cmd)

    fmt.Println("yq instalado com sucesso.")
}