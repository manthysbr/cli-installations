package main

import (
    "fmt"
    "log"
    "os"
    "manthys/cmd/check"
    // ... outras importações ...
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Uso: manthys <comando> [argumentos]")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "check":
        check.RunCheck()
    default:
        fmt.Printf("Comando desconhecido: %s\n", os.Args[1])
        os.Exit(1)
    }
}
