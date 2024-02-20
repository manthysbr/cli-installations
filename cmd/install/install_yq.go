package install

import (
	"fmt"
	"log"
	"os/exec"
)

// InstallYq instala o yq utilizando wget para baixar o binário diretamente.
func InstallYq() {
    fmt.Println("yq está sendo instalado. Aguarde...")

    // Preparando ambiente
    PrepararAmbiente() // Certifique-se de que esta função esteja exportada corretamente no common_install.go

    // Define o comando para baixar e instalar o yq
    installCmd := exec.Command("bash", "-c", "wget https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64 -O /usr/bin/yq && chmod +x /usr/bin/yq")

    // Executa o comando de instalação
    if err := ExecuteInstallCommand(installCmd); err != nil {
        log.Fatalf("Erro ao instalar yq: %v", err)
    }

    fmt.Println("yq instalado com sucesso.")
}

// ExecuteInstallCommand executa um comando de instalação e fornece feedback ao usuário.
// Esta função deve estar definida no seu common_install.go e exportada corretamente.
func ExecuteInstallCommand(cmd *exec.Cmd) error {
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("falha ao executar comando: %w", err)
    }
    return nil
}
