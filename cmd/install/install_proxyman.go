package install

import (
    "log"
    "os/exec"
)

func InstallProxyman() {
    // Replace with the actual command to install Proxyman
    cmd := exec.Command("sudo", "apt-get", "install", "-y", "proxyman")
    err := cmd.Run()

    if err != nil {
        log.Fatalf("Failed to install Proxyman: %s", err)
    } else {
        log.Println("Proxyman installed successfully")
    }
}