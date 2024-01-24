package help

import "fmt"

func PrintHelp() {
    fmt.Println(`Usage: manthys <command> [arguments]

The commands are:

    check   Run a check to see if the software is already installed and save the result in the software.json file.

    install Install the software based on the software.json file. Software already installed will be reconfigured to use the cli version.
            Example: manthys install

    help    Show this help message.
            Example: manthys help

Use "manthys <command> -h" for more information about a command.
`)
}

func PrintCheckHelp() {
    fmt.Println(`Runs an autocheck if the needed software is already installed and save the result in the software.json file. 
				This json file is used by the install command based into what is already installed.`)
}


func PrintInstallHelp() {
    fmt.Println(`Install the software based on the software.json file. Software already installed will be reconfigured to use the cli version.`)
}