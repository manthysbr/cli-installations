package check

import (
    "flag"
    "manthys/cmd/help"
)

func RunCheck() {
    helpFlag := flag.Bool("h", false, "Show help for command")
    flag.Parse()

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