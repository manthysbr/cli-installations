package check

import (
    "log"
    "os"
    "os/exec"
    "strings"
)

func CheckPython() {
    cmd := exec.Command("python", "--version")
    output, err := cmd.CombinedOutput()

    if err != nil {
        log.Printf("Error checking Python version: %v", err)
        return
    }

    version := strings.TrimSpace(string(output))
    log.Printf("Python version found: %s", version)
}