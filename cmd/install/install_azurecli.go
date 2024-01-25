package install

import (
    "log"
    "os/exec"
)

func InstallAzureCLI() {
    cmd := exec.Command("curl", "-sL", "https://aka.ms/InstallAzureCLIDeb", "|", "sudo", "bash")
    err := cmd.Run()

    if err != nil {
        log.Fatalf("Failed to install Azure CLI: %s", err)
    } else {
        log.Println("Azure CLI installed successfully")
    }
}