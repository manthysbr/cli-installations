package utils

import (
	"fmt"
	"github.com/spf13/cobra"
)

// helpCmd representa um comando personalizado de ajuda
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Exibe informações de ajuda para manthys",
	Long: `Ajuda mostra informações sobre os subcomandos e temas de ajuda.
Use "manthys help [comando]" para obter informações sobre um comando específico.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Aqui você pode chamar a função de ajuda padrão do Cobra ou implementar a sua própria
		if len(args) == 0 {
			rootCmd.HelpFunc()(cmd, args)
		} else {
			subCmd, _, err := rootCmd.Find(args)
			if err != nil {
				fmt.Printf("Erro: %s\n", err)
				return
			}
			subCmd.HelpFunc()(subCmd, args)
		}
	},
}

func init() {
	rootCmd.AddCommand(helpCmd) // Adiciona o comando help como subcomando de rootCmd
}
