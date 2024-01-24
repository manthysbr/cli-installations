package main

import (
    "flag"
    "fmt"
    "os"
    "manthys/cmd/check"
    //"manthys/cmd/install"
    "manthys/cmd/help"
    // ... other imports ...
)

func main() {
    // Define a boolean flag for -h
    helpFlag := flag.Bool("h", false, "Show help for command")
    // Parse the flags
    flag.Parse()

    if len(os.Args) < 2 {
        help.PrintHelp()
        os.Exit(1)
    }

    switch os.Args[1] {
    case "check":
        if *helpFlag {
            help.PrintCheckHelp() // call the PrintCheckHelp function
        } else {
            check.RunCheck()
        }
    case "install":
        if *helpFlag {
            help.PrintInstallHelp() // call the PrintInstallHelp function
        } else {
            // call your install function here
        }
    case "help":
        help.PrintHelp()
    default:
        fmt.Printf("Comando desconhecido: %s\n", os.Args[1])
        os.Exit(1)
    }
}