package install

import (
    "bufio"
    "fmt"
    "log"
    "os/exec"
    "strings"
)

func InstallPython() {
    fmt.Println("Iniciando a instalação do Python...")

    cmd := exec.Command("sudo", "apt-get", "install", "-y", "python3")

    // Cria um pipe para a saída padrão do comando
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatalf("Erro ao criar o pipe para a saída padrão: %s", err)
    }

    // Inicia o comando
    if err := cmd.Start(); err != nil {
        log.Fatalf("Falha ao iniciar a instalação do Python: %s", err)
    }

    // Lê a saída linha por linha
    scanner := bufio.NewScanner(stdout)
    for scanner.Scan() {
        line := scanner.Text()
        // Aqui você pode filtrar as linhas que deseja mostrar
        if strings.Contains(line, "Unpacking") || strings.Contains(line, "Setting up") {
            fmt.Println(line)
        }
    }

    // Espera o comando terminar e verifica se houve erro
    if err := cmd.Wait(); err != nil {
        log.Fatalf("Falha ao instalar o Python: %s", err)
    } else {
        fmt.Println("Python instalado com sucesso")
    }
}
