package install

import (
	"bufio"
	"fmt"
	"github.com/briandowns/spinner"
	"log"
	"os/exec"
	"time"
)

// executeInstallCommand executa um comando de instalação e fornece feedback ao usuário
func executeInstallCommand(cmd *exec.Cmd) {
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Erro ao criar o pipe para a saída padrão: %s", err)
	}

	s := spinner.New(spinner.CharSets[39], 100*time.Millisecond) // Cria um novo spinner
	s.Start()                                                   // Inicia o spinner

	if err := cmd.Start(); err != nil {
		s.Stop() // Para o spinner em caso de erro
		log.Fatalf("Falha ao iniciar a instalação: %s", err)
	}

    // Novo: Cria um scanner para ler a saída do comando em tempo real
    scanner := bufio.NewScanner(stdoutPipe)
    go func() {
        for scanner.Scan() {
            fmt.Println(scanner.Text()) // Exibe cada linha da saída
        }
    }()

	s.Stop() // Para o spinner logo antes de esperar o comando terminar

	if err := cmd.Wait(); err != nil {
		log.Fatalf("Falha na instalação: %s", err)
	} else {
		fmt.Printf("%s instalado com sucesso\n", cmd.Args[len(cmd.Args)-1])
	}
}

func prepararAmbiente() {
    fmt.Println("Preparando ambiente...")

    s := spinner.New(spinner.CharSets[39], 100*time.Millisecond) // Cria um novo spinner
    s.Start() // Inicia o spinner

    time.Sleep(2 * time.Second) // Simula o tempo de preparação do ambiente

    s.Stop() // Para o spinner após a simulação

    fmt.Println("Ambiente preparado.")
}