package check

import (
	"manthys/internal/checker" // Certifique-se de que o caminho está correto
	"github.com/spf13/cobra"
)

// CheckCmd representa o comando check e encapsula a lógica de checagem
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Verifica o estado da instalação dos softwares",
	Long:  `Executa uma checagem para determinar se os softwares necessários estão instalados e configurados corretamente.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Chama as funções do pacote checker
		checker.CheckGit()
		checker.CheckPython()
		checker.CheckAzureCLI()
		checker.CheckDocker()
		checker.CheckProxyman()
		// Adicione chamadas para outras funções de checagem conforme necessário
	},
}

// Inicialização do comando check, pode ser usada para configurar flags e outros setup inicial
func init() {
	// Aqui você pode configurar flags e argumentos do comando
	// Por exemplo:
	// CheckCmd.Flags().StringVarP(&pythonVersion, "python-version", "p", "", "Especifique a versão do Python para verificar")
}
