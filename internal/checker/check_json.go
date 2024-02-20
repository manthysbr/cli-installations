package checker

import (
    "io/ioutil"
    "log"
    "os"
    "manthys/pkg/config"
)

func CheckJSON() {
    filePath := "/tmp/software_state.json"

    // Verifica se o arquivo existe
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        log.Printf("Arquivo %s não encontrado. Criando um novo arquivo.", filePath)
        
        // Conteúdo inicial para o novo arquivo software_state.json
        initialContent := []byte(`{"software":[]}`)

        // Cria o arquivo com o conteúdo inicial
        if err := ioutil.WriteFile(filePath, initialContent, 0644); err != nil {
            log.Printf("Erro ao criar o arquivo %s: %s", filePath, err)
            config.SaveSoftwareState("software_state.json", "Not created", "")
            return
        }

        log.Printf("Arquivo %s criado com sucesso.", filePath)
        config.SaveSoftwareState("software_state.json", "Created", "Empty file created")
    } else if err != nil {
        // Se ocorrer um erro ao tentar verificar o arquivo (diferente de "não encontrado")
        log.Printf("Erro ao verificar a existência do arquivo %s: %s", filePath, err)
        config.SaveSoftwareState("software_state.json", "Check failed", err.Error())
    } else {
        log.Printf("Arquivo %s encontrado.", filePath)
        config.SaveSoftwareState("software_state.json", "Found", "File exists")
    }
}
