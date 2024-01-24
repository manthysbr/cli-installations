package install

import (
    // Importações necessárias
    "manthys/pkg/config"
)

func RunInstall() {
    // Verificar o estado dos softwares do arquivo de configuração
    state := config.GetSoftwareState("python")
    if state == "não instalado" {
        // Instalar o software
    }
    // Repetir para outros softwares
}
