package help

import "fmt"

func PrintHelp() {
    fmt.Println(`Usage: manthys <command> [arguments]

The commands are:

    check   Check if the software is already installed. 
            Example: manthys check <software_name>

    install Install the software.
            Example: manthys install <software_name>

    help    Show this help message.
            Example: manthys help

Use "manthys <command> -h" for more information about a command.
`)
}

func PrintCheckHelp() {
    fmt.Println(`Check if the software is already installed and save the result in the software.json file. 
				This json file is used by the install command based into what is already installed.`)
}


func PrintInstallHelp() {
    fmt.Println(`Install the software based on the software.json file. Software already installed will be reconfigured to use the cli version.`)
}