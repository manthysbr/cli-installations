package checker

import (
    "log"
    "os/exec"
    "strings"
    "manthys/pkg/config"
)

func CheckYq() string {
    cmd := exec.Command("yq", "--version")
    output, err := cmd.CombinedOutput()
    
    if err != nil {
        log.Printf("YQ is not installed.")
        config.SaveSoftwareState("YQ", "Not installed", "")
        return "Not installed"
    }

    version := strings.TrimSpace(string(output))
    log.Printf("Proxyman version found: %s", version)
    config.SaveSoftwareState("YQ", "Installed", version)
    return "Installed"
}