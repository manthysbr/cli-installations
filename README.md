# CLI de Instalação de Software - Manthys CLI

[![Go Reference](https://pkg.go.dev/badge/github.com/seuprojeto/cli-installation.svg)](https://pkg.go.dev/github.com/seuprojeto/cli-installation)
[![go.mod](https://img.shields.io/github/go-mod/go-version/seuprojeto/cli-installation)](go.mod)
[![LICENSE](https://img.shields.io/github/license/seuprojeto/cli-installation)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/seuprojeto/cli-installation/build.yml?branch=main)](https://github.com/seuprojeto/cli-installation/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/seuprojeto/cli-installation)](https://goreportcard.com/report/github.com/seuprojeto/cli-installation)

⭐ `Star` este repositório se você achar útil e digno de manutenção.

👁 `Watch` este repositório para ser notificado sobre novas releases, issues, etc.

## Descrição

Esta é uma ferramenta CLI desenvolvida em Go para automatizar a instalação de softwares essenciais como o Azure CLI e Docker, proporcionando feedback visual através de spinners durante o processo de instalação. Ideal para simplificar a configuração de ambientes de desenvolvimento ou produção.

Inclui:

- Integração contínua via [GitHub Actions](https://github.com/features/actions),
- Gerenciamento de dependências usando [Go Modules](https://github.com/golang/go/wiki/Modules),
- Formatação de código usando [gofumpt](https://github.com/mvdan/gofumpt),
- Análise de código com [golangci-lint](https://github.com/golangci/golangci-lint),
- Testes unitários com detecção de corrida e relatórios de cobertura de código.

## Uso

1. Clone ou baixe este repositório.
2. Substitua todas as ocorrências de `seuprojeto/cli-installation` pelo caminho do seu próprio repositório.
3. Execute `go build` para construir a aplicação.
4. Execute a CLI construída para instalar o software desejado, por exemplo, `./cli-installation install docker`.

## Configuração

Instruções de configuração do ambiente de desenvolvimento:

1. Instale [Go](https://golang.org/doc/install).
2. Clone este repositório e abra-o.
3. Execute `go mod tidy` para garantir que todas as dependências estejam corretas.

## Build

### Terminal

- Execute `go build` para construir a aplicação.

## Release

O workflow de release é acionado sempre que uma tag com prefixo `v` é criada e enviada.

## Manutenção

Arquivos notáveis:

- [.github/workflows](.github/workflows) - Workflows do GitHub Actions,
- [go.mod](go.mod) - Definição do módulo Go.

## Contribuindo

Sinta-se à vontade para criar uma issue ou propor um pull request.

Siga o [Código de Conduta](CODE_OF_CONDUCT.md).
