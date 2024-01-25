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
        softwareState := config.LoadSoftwareState()
        for software, state := range softwareState.Software {
            if state == "Not installed" {
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
            }
        }
    } else {
        fmt.Printf("%s is installed with version: %s\n", software, state)
    }
}
