package install

import (
    "bufio"
    "fmt"
    "log"
    "os/exec"
    "strings"
)

func InstallYqJq() {
    fmt.Println("Iniciando a instalação do yq e jq...")

    // Comandos para instalar yq e jq
    commands := []*exec.Cmd{
        exec.Command("sudo", "apt-get", "install", "-y", "jq"),
        exec.Command("sudo", "apt-get", "install", "-y", "yq"),
    }

    for _, cmd := range commands {
        // Executa cada comando
        executeInstallCommand(cmd)
    }
}

func executeInstallCommand(cmd *exec.Cmd) {
    // Cria um pipe para a saída padrão do comando
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatalf("Erro ao criar o pipe para a saída padrão: %s", err)
    }

    // Inicia o comando
    if err := cmd.Start(); err != nil {
        log.Fatalf("Falha ao iniciar a instalação: %s", err)
    }

    // Lê a saída linha por linha
    scanner := bufio.NewScanner(stdout)
    for scanner.Scan() {
        line := scanner.Text()
        // Filtra as linhas desejadas
        if strings.Contains(line, "Unpacking") || strings.Contains(line, "Setting up") {
            fmt.Println(line)
        }
    }

    // Espera o comando terminar e verifica se houve erro
    if err := cmd.Wait(); err != nil {
        log.Fatalf("Falha na instalação: %s", err)
    } else {
        fmt.Printf("%s instalado com sucesso\n", cmd.Args[len(cmd.Args)-1])
    }
}
