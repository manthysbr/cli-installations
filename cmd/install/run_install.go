package install

import (
    "fmt"
    "manthys/pkg/config"
)

func RunInstall() {
    // Load the software state from the JSON file
    softwareState := config.LoadSoftwareState()

    // For each software in the software state
    for software, state := range softwareState.Software {
        // If the software is not installed
        if state == "Not installed" {
            // Call the corresponding function to install the software
            switch software {
            case "Python":
                InstallPython()
            // Add cases for other software
            }
        } else {
            // If the software is installed, print the version
            fmt.Printf("%s is installed with version: %s\n", software, state)
        }
    }
}