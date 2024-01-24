package install

import (
    "log"
    "os/exec"
)

func InstallGit() {
    cmd := exec.Command("sudo", "apt-get", "install", "-y", "git")
    err := cmd.Run()

    if err != nil {
        log.Fatalf("Failed to install Git: %s", err)
    } else {
        log.Println("Git installed successfully")
    }
}