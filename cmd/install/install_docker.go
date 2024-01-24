package install

import (
    "log"
    "os/exec"
)

func InstallDocker() {
    cmd := exec.Command("sudo", "apt-get", "install", "-y", "docker.io")
    err := cmd.Run()

    if err != nil {
        log.Fatalf("Failed to install Docker: %s", err)
    } else {
        log.Println("Docker installed successfully")
    }
}