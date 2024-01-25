package main

import (
    "fmt"
    "os"
    "manthys/cmd/validate"
    "manthys/cmd/install"
    "manthys/cmd/utils"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "manthys",
    Short: "Manthys é uma CLI para configurar rapidamente o seu ambiente de desenvolvimento",
    Long:  `Uma ferramenta completa para a instalação e configuração de ambientes de desenvolvimento com suporte a múltiplas operações e verificações.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(utils.GetAsciiArt())
        fmt.Println("Bem-vindo ao Manthys CLI! Use 'manthys help' para ver os comandos disponíveis.")
    },
}

func main() {
    utils.AddCommands(rootCmd) // Adiciona comandos personalizados, como o comando de ajuda
    rootCmd.AddCommand(validate.CheckCmd)
    rootCmd.AddCommand(install.RunInstallCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
