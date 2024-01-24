package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// rootCmd representa o comando base quando chamado sem subcomandos
var rootCmd = &cobra.Command{
	Use:   "manthys",
	Short: "Manthys é uma CLI para configurar rapidamente o seu ambiente de desenvolvimento",
	Long: `Uma ferramenta completa para a instalação e configuração de ambientes de desenvolvimento
com suporte a múltiplas operações e verificações.`,
	// Aqui você pode tratar a lógica quando o comando manthys for chamado sem subcomandos
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bem-vindo ao Manthys CLI! Use 'manthys [comando]' para interagir com a aplicação.")
	},
}

func init() {
	// Aqui você pode inicializar os flags e subcomandos
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
