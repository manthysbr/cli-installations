package install

import (
    "fmt"
    "manthys/pkg/config"
    "manthys/cmd/report"
)

func RunInstall() {
    // Gera o relatório e pede confirmação
    if report.GenerateInstallationReport() {
        // O usuário confirmou, prossegue com a instalação
        softwareState := config.GetSoftwareState()
        for software, details := range softwareState.Software {
            if details["state"] == "Not installed" {
                // Chama a função correspondente para instalar o software
                switch software {
                case "Python":
                    InstallPython()
                case "Git":
                    InstallGit()
                case "Docker":
                    InstallDocker()
                case "Proxyman":
                    InstallProxyman()
                case "AzureCLI":
                    InstallAzureCLI()
                }
            } else {
                fmt.Printf("%s is installed with version: %s\n", software, details["version"])
            }
        }
    } else {
        // Mensagem caso o usuário opte por não instalar
        fmt.Println("Instalação cancelada pelo usuário.")
    }
}
