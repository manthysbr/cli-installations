package install

import (
    "fmt"
    "manthys/pkg/config"
)

func RunInstall() {
    // Carrega o estado do software do arquivo JSON
    softwareState := config.LoadSoftwareState()

    // Para cada software no estado do software
    for software, state := range softwareState.Software {
        // Se o software não estiver instalado
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
            // Adicione casos para outros softwares conforme necessário
            }
        } else {
            // Se o software estiver instalado, imprime a versão
            fmt.Printf("%s is installed with version: %s\n", software, state)
        }
    }
}