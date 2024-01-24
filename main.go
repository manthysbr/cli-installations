package main

import (
    "fmt"
    "os"
    "manthys/cmd/check"
    "manthys/cmd/install"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Uso: manthys <comando> [argumentos]")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "check":
        check.RunCheck()
    case "install":
        install.RunInstall()
    default:
        fmt.Printf("Comando desconhecido: %s\n", os.Args[1])
        os.Exit(1)
    }
}
