package help

import (
    "fmt"
    "manthys/pkg/config"
)

func GenerateReport() {
    // Load the software state from the JSON files
    softwareState := config.GetSoftwareState()

    // Print a list of the software to be installed
    fmt.Println("The following software will be installed:")
    for software, details := range softwareState.Software {
        if details["state"] == "não instalado" {
            fmt.Println("- " + software)
        }
    }
}