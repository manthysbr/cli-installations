package checker

import (
    "log"
    "os/exec"
    "strings"
    "manthys/pkg/config"
)

func CheckAzureCLI() string {
    cmd := exec.Command("az", "--version")
    output, err := cmd.CombinedOutput()
    
    if err != nil {
        log.Printf("Azure CLI is not installed.")
        config.SaveSoftwareState("AzureCLI", "Not installed", "")
        return "Not installed"
    }

    version := strings.TrimSpace(string(output))
    log.Printf("Azure CLI version found: %s", version)
    config.SaveSoftwareState("AzureCLI", "Installed", version)
    return "Installed"
}