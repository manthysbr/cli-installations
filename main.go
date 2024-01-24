package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Uso: manthys <comando> [argumentos]")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "install":
        installPackages()
    default:
        fmt.Printf("Comando desconhecido: %s\n", os.Args[1])
        os.Exit(1)
    }
}

func installPackages() {
    packages := []string{"wget", "kubectl", "nginx"}

    fmt.Println("Instalando pacotes:")
    for _, pkg := range packages {
        fmt.Printf("Instalando %s...\n", pkg)
        cmd := exec.Command("wsl", "sudo", "apt", "install", "-y", pkg)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        err := cmd.Run()
        if err != nil {
            fmt.Printf("Erro ao instalar %s: %s\n", pkg, err)
        }
    }
}
