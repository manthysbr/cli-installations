package check

import (
    "log"
    "os/exec"
    "strings"
    "manthys/pkg/config"
)

func CheckPython() string {
    cmd := exec.Command("python3", "--version")
    output, err := cmd.CombinedOutput()
    
    if err != nil {
        log.Printf("Python is not installed.")
        config.SaveSoftwareState("Python", "Not installed")
        return "Not installed"
    }

    version := strings.TrimSpace(string(output))
    log.Printf("Python version found: %s", version)
    config.SaveSoftwareState("Python", "Installed")
    return "Installed"
}