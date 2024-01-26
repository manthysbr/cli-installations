package install

import (
    "fmt"
    "os/exec"
)

func InstallYq() {
    fmt.Println("YQ está sendo instalado. Aguarde...")

    // Preparando ambiente
    prepararAmbiente()

    // Install yq using go get
    installCmd := exec.Command("go", "get", "github.com/mikefarah/yq/v3")
    installCmd.Env = append(os.Environ(), "GO111MODULE=on")
    executeInstallCommand(installCmd)
}