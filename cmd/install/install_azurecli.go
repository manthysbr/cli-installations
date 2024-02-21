package install

import (
    "fmt"
    "os/exec"
)

// InstallAzureCLI instala o Azure CLI seguindo os passos da documentação da Microsoft.
func InstallAzureCLI() error {
    fmt.Println("Iniciando a instalação do Azure CLI...")

    // Preparando ambiente
    err := prepararAmbiente()
    if err != nil {
        return fmt.Errorf("erro ao preparar ambiente: %v", err)
    }

    // 1. Baixa e instala o pacote de transporte HTTPS para APT, se necessário.
    installTransportHTTPS := exec.Command("sudo", "apt-get", "update")
    err = executeInstallCommand(installTransportHTTPS)
    if err != nil {
        return fmt.Errorf("erro ao instalar transporte HTTPS: %v", err)
    }

    installTransportHTTPS = exec.Command("sudo", "apt-get", "install", "-y", "apt-transport-https", "ca-certificates", "curl", "software-properties-common")
    err = executeInstallCommand(installTransportHTTPS)
    if err != nil {
        return fmt.Errorf("erro ao instalar transporte HTTPS: %v", err)
    }

    // 2. Importa a chave de repositório da Microsoft.
    importMicrosoftGPGKey := exec.Command("bash", "-c", "curl -sL https://packages.microsoft.com/keys/microsoft.asc | sudo apt-key add -")
    err = executeInstallCommand(importMicrosoftGPGKey)
    if err != nil {
        return fmt.Errorf("erro ao importar chave GPG da Microsoft: %v", err)
    }

    // 3. Adiciona o repositório de software do Azure CLI.
    addAzureCLIRepo := exec.Command("sudo", "sh", "-c", `echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ $(lsb_release -cs) main" > /etc/apt/sources.list.d/azure-cli.list`)
    err = executeInstallCommand(addAzureCLIRepo)
    if err != nil {
        return fmt.Errorf("erro ao adicionar repositório do Azure CLI: %v", err)
    }

    // 4. Atualiza o índice de pacotes do APT e instala o pacote azure-cli.
    updateAndInstallAzureCLI := exec.Command("sudo", "apt-get", "update")
    err = executeInstallCommand(updateAndInstallAzureCLI)
    if err != nil {
        return fmt.Errorf("erro ao atualizar índice de pacotes do APT: %v", err)
    }

    installAzureCLI := exec.Command("sudo", "apt-get", "install", "-y", "azure-cli")
    err = executeInstallCommand(installAzureCLI)
    if err != nil {
        return fmt.Errorf("erro ao instalar Azure CLI: %v", err)
    }

    fmt.Println("Azure CLI instalado com sucesso.")
    return nil
}