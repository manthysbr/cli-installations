package utils

import (
    "fmt"
    "github.com/spf13/cobra"
)

// HelpCmd representa um comando personalizado de ajuda
var HelpCmd = &cobra.Command{
    Use:   "help",
    Short: "Exibe informações de ajuda para manthys",
    Long: `Ajuda mostra informações sobre os subcomandos e temas de ajuda.
Use "manthys help [comando]" para obter informações sobre um comando específico.`,
    Run: func(cmd *cobra.Command, args []string) {
        // Implemente a lógica de ajuda aqui
        if len(args) == 0 {
            cmd.Parent().HelpFunc()(cmd, args)
        } else {
            subCmd, _, err := cmd.Parent().Find(args)
            if err != nil {
                fmt.Printf("Erro: %s\n", err)
                return
            }
            subCmd.HelpFunc()(subCmd, args)
        }
    },
}

// AddCommands adiciona comandos personalizados ao rootCmd
func AddCommands(root *cobra.Command) {
    root.AddCommand(HelpCmd)
    // Adicione outros comandos personalizados aqui, se necessário
}
