package validate

import (
	"manthys/internal/checker"
	"github.com/spf13/cobra"
)

// CheckCmd representa o comando check e encapsula a lógica de checagem
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Verifica o estado da instalação dos softwares",
	Long:  `Executa uma checagem para determinar se os softwares necessários estão instalados e configurados corretamente.`,
	Run: func(cmd *cobra.Command, args []string) {
		checker.CheckJSON()
		checker.CheckGit()
		checker.CheckPython()
		checker.CheckAzureCLI()
		checker.CheckDocker()
		checker.CheckYq()
		checker.CheckJq()
	},
}

func init() {
	// Aqui você pode configurar flags e argumentos do comando
}
