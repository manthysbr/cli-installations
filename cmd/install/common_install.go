package install

import (
    "bufio"
    "fmt"
    "log"
    "os/exec"
    "strings"
)

// executeInstallCommand executa um comando de instalação e fornece feedback ao usuário
func executeInstallCommand(cmd *exec.Cmd) {
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatalf("Erro ao criar o pipe para a saída padrão: %s", err)
    }

    if err := cmd.Start(); err != nil {
        log.Fatalf("Falha ao iniciar a instalação: %s", err)
    }

    scanner := bufio.NewScanner(stdout)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "Unpacking") || strings.Contains(line, "Setting up") {
            fmt.Println(line)
        }
    }

    if err := cmd.Wait(); err != nil {
        log.Fatalf("Falha na instalação: %s", err)
    } else {
        fmt.Printf("%s instalado com sucesso\n", cmd.Args[len(cmd.Args)-1])
    }
}

func prepararAmbiente() {
	// Simula o tempo de preparação do ambiente
	time.Sleep(2 * time.Second)
	fmt.Println("Preparando ambiente...")
}