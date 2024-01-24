package check

import (
    "log"
    "os/exec"
    "strings"
    "manthys/pkg/config"
)

func CheckProxy() string {
    cmd := exec.Command("proxyman", "--version")
    output, err := cmd.CombinedOutput()
    
    if err != nil {
        log.Printf("Proxyman is not installed.")
        config.SaveSoftwareState("Proxyman", "Not installed", "")
        return "Not installed"
    }

    version := strings.TrimSpace(string(output))
    log.Printf("Proxyman version found: %s", version)
    config.SaveSoftwareState("Proxyman", "Installed", version)
    return "Installed"
}