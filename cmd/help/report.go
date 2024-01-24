package report

import (
    "fmt"
    "manthys/pkg/config"
)

func GenerateReport() {
    // Load the software state from the JSON files
    softwareState := config.LoadSoftwareState()

    // Print a list of the software to be installed
    fmt.Println("The following software will be installed:")
    for software, status := range softwareState.Software {
        if status == "não instalado" {
            fmt.Println("- " + software)
        }
    }
}