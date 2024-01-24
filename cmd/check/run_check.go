package check

import (
    "flag"
    "os"
    "manthys/cmd/help"
)

func RunCheck() {
    // Create a new flag set
    flags := flag.NewFlagSet("check", flag.ExitOnError)

    // Define the -h flag in this flag set
    helpFlag := flags.Bool("h", false, "Show help for command")

    // Parse the flags
    flags.Parse(os.Args[2:])

    if *helpFlag {
        help.PrintCheckHelp()
        return
    }

    // Chama a função de checagem para Python
    CheckPython()
    CheckGit()
    CheckDocker()
    CheckProxy()
    CheckAzureCli()
    // Chama funções de checagem para outros softwares    // Exemplo: CheckGit(), CheckDocker(), etc.
}