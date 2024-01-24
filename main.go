package main

import (
	"fmt"
	"os"
	"manthys/cmd/validate" // Certifique-se de que o caminho está correto com base na localização do seu pacote check
	"github.com/spf13/cobra"
)

// rootCmd representa o comando base quando chamado sem subcomandos
var rootCmd = &cobra.Command{
	Use:   "manthys",
	Short: "Manthys é uma CLI para configurar rapidamente o seu ambiente de desenvolvimento",
	Long:  `Uma ferramenta completa para a instalação e configuração de ambientes de desenvolvimento com suporte a múltiplas operações e verificações.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bem-vindo ao Manthys CLI! Use 'manthys help' para ver os comandos disponíveis.")
	},
}

func main() {
	// Adiciona o comando 'check' como um subcomando de 'rootCmd'
	rootCmd.AddCommand(check.CheckCmd)

	// Execute o comando raiz e todos os subcomandos associados
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
