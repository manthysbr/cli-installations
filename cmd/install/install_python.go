package install

import (
    "log"
    "os/exec"
)

func InstallPython() {
    cmd := exec.Command("sudo", "apt-get", "install", "-y", "python3")
    err := cmd.Run()

    if err != nil {
        log.Fatalf("Failed to install Python: %s", err)
    } else {
        log.Println("Python installed successfully")
    }
}