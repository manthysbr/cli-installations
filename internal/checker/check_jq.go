package checker

import (
    "log"
    "os/exec"
    "strings"
    "manthys/pkg/config"
)

func CheckJq() string {
    cmd := exec.Command("jq", "--version")
    output, err := cmd.CombinedOutput()
    
    if err != nil {
        log.Printf("JQ is not installed.")
        config.SaveSoftwareState("YQ", "Not installed", "")
        return "Not installed"
    }

    version := strings.TrimSpace(string(output))
    log.Printf("JQ version found: %s", version)
    config.SaveSoftwareState("JQ", "Installed", version)
    return "Installed"
}