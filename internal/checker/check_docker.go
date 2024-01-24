package checker

import (
    "log"
    "os/exec"
    "strings"
    "manthys/pkg/config"
)

func CheckDocker() string {
    cmd := exec.Command("docker", "--version")
    output, err := cmd.CombinedOutput()
    
    if err != nil {
        log.Printf("Docker is not installed.")
        config.SaveSoftwareState("Docker", "Not installed", "")
        return "Not installed"
    }

    version := strings.TrimSpace(string(output))
    log.Printf("Docker version found: %s", version)
    config.SaveSoftwareState("Docker", "Installed", version)
    return "Installed"
}