package install

import (
    "manthys/pkg/config"
)

func RunInstall() {
    // Load the software state from the JSON file
    softwareState := config.LoadSoftwareState()

    // Install the software if it's not installed
    if softwareState.Software["Python"] == "Not installed" {
        InstallPython()
    }
    // Repeat for other software
}