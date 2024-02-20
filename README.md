# CLI de Instalação de Software

[![Go Reference](https://pkg.go.dev/badge/github.com/seuprojeto/cli-installation.svg)](https://pkg.go.dev/github.com/manthysbr/cli-installation)
[![go.mod](https://img.shields.io/github/go-mod/go-version/manthysbr/cli-installation)](go.mod)
[![LICENSE](https://img.shields.io/github/license/manthysbr/cli-installation)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/manthysbr/cli-installation/build.yml?branch=main)](https://github.com/seuprojeto/cli-installation/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/manthysbr/cli-installation)](https://goreportcard.com/report/github.com/seuprojeto/cli-installation)

⭐ `Star` este repositório se você achar útil e digno de manutenção.

👁 `Watch` este repositório para ser notificado sobre novas releases, issues, etc.

## Descrição

Esta é uma  CLI desenvolvida em Go para automatizar a instalação e configuração de softwares essenciais para o meu trabalho e estudos. Criei por preguiça de toda vez ter que levar horas pra configurar um ambiente e para testar minhas habilidades com go, shell e tudo mais. Manthys é por conta do Radamanthys dos Cavaleiros do Zodíaco :heart_eyes: 

![O mais brabo de todes](https://i.pinimg.com/originals/1f/24/90/1f24906a5b38e83d5e73e666419bb0b0.gif)

Adicionei:

- Integração contínua via [GitHub Actions](https://github.com/features/actions),
- Gerenciamento de dependências usando [Go Modules](https://github.com/golang/go/wiki/Modules),
- Formatação de código usando [gofumpt](https://github.com/mvdan/gofumpt),
- Análise de código com [golangci-lint](https://github.com/golangci/golangci-lint),

Ainda falta:
- Testes unitários com detecção de corrida e relatórios de cobertura de código.

## Uso para desenvolvimento

1. Clone ou baixe este repositório.
2. Altere a permissão do compile.sh para executá-lo ( chmod +x )
3. Execute o compile.sh
4. Depois da execução com sucesso do script, volte um diretório e acesse a pasta build ( home user build )
4. Copie o script install.sh para a pasta de build e altere sua permissão
5. Execute install.sh

Pronto, assim dá pra executar de qualquer lugar do ambiente sem precisar apontar nenhum path.

## Uso para usuário final

1. Baixe o tar.gz nas releases.
2. Mova o binario para o seu path global ( o install.sh já faz isso )
3. Só executar manthys

## Configuração

Instruções de configuração do ambiente de desenvolvimento:

1. Instale [Go](https://golang.org/doc/install).
2. Clone este repositório e abra-o.
3. Execute `go mod tidy` para garantir que todas as dependências estejam corretas.

## Release

O workflow de release é acionado sempre que uma tag com prefixo `v` é criada e enviada.

## Manutenção

Arquivos notáveis:

- [.github/workflows](.github/workflows) - Workflows do GitHub Actions,
- [go.mod](go.mod) - Definição do módulo Go.

## Contribuindo

Sinta-se à vontade para criar uma issue ou propor um pull request.

Siga o [Código de Conduta](CODE_OF_CONDUCT.md) quando eu terminar.
