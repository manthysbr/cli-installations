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
        fmt.Println("Iniciando o processo de instalação...")
        if report.GenerateInstallationReport() {
            softwareState := config.GetSoftwareState()
            for software, details := range softwareState.Software {
                if details["state"] == "Not installed" {
                    fmt.Printf("Instalando %s...\n", software)
                    installSoftware(software)
                } else {
                    fmt.Printf("%s já está instalado com a versão: %s\n", software, details["version"])
                }
            }
        } else {
            fmt.Println("Instalação cancelada pelo usuário.")
        }
    },
}

func installSoftware(software string) {
    switch software {
    case "Python":
        InstallPython()
    case "Git":
        InstallGit()
    case "Docker":
        InstallDocker()
    case "YQ":
        InstallYq()
    case "JQ":
        InstallJq()
    case "AzureCLI":
        InstallAzureCLI()
    }
}

