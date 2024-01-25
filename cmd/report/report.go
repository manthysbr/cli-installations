package report

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "manthys/pkg/config"
)

func GenerateInstallationReport() bool {
    // Carrega o estado do software do arquivo JSON
    softwareState := config.LoadSoftwareState()

    fmt.Println("Os seguintes softwares serão instalados:")
    for software, state := range softwareState.Software {
        if state == "Not installed" {
            fmt.Println("- " + software)
        }
    }

    // Pede a confirmação do usuário
    return getUserConfirmation()
}

func getUserConfirmation() bool {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Deseja prosseguir com a instalação? (y/n): ")
    response, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Erro ao ler a entrada. Presumindo 'não'.")
        return false
    }
    response = strings.TrimSpace(response)
    return response == "y" || response == "Y"
}
