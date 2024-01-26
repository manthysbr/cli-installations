package install

import (
    "fmt"
    "manthys/pkg/config"
    "manthys/cmd/report"
    "github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
    Use:   "install",
    Short: "Instala os softwares necessários",
    Long:  `Instala todos os softwares que estão marcados como 'Not installed' no arquivo software_state.json.`,
    Run: func(cmd *cobra.Command, args []string) {
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
                    case "YQ e JQ":
                        InstallYqJq()
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
    },
}
