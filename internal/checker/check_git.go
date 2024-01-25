package checker

import (
    "log"
    "os/exec"
    "strings"
    "manthys/pkg/config"
)

func CheckGit() string {
    cmd := exec.Command("git", "--version")
    output, err := cmd.CombinedOutput()
    
    if err != nil {
        log.Printf("Git is not installed.")
        config.SaveSoftwareState("Git", "Not installed", "")
        return "Not installed"
    }

    version := strings.TrimSpace(string(output))
    log.Printf("Git version found: %s", version)
    config.SaveSoftwareState("Git", "Installed", version)
    return "Installed"
}