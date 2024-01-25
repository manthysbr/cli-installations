# Manthys CLI

## Visão Geral

Manthys é uma ferramenta de linha de comando (CLI) desenvolvida para automatizar e facilitar a configuração de ambientes de desenvolvimento. Ela verifica a instalação de softwares essenciais pré-definidos, informa suas versões e, se necessário, procede com a instalação automatizada.

A idéia é sempre manter os softwares atualizados para trabalho, por isso é sempre bom a homologação dos softwares antes de qualquer tipo de implantação.

A CLI usa Golang e Cobra para uma experiência de usuário aprimorada e funcionalidades robustas.

## Recursos

- Verificação automática da presença e versão de softwares essenciais como Python, Git, Docker, AzureCLI e Proxyman (mais em desenvolvimento).
- Instalação automatizada de softwares não presentes no sistema.
- Interface de usuário amigável e informativa, com saída detalhada sobre o estado de cada software.

## Instalação

Para instalar a CLI Manthys, siga os passos abaixo:

1. Clone o repositório:
`git clone [https://github.com/manthysbr/cli-installations]`


2. Navegue até o diretório:
	 `cd manthys`

3. Compile o código:
     `go build`

4. Execute a CLI
     `./manthys`


## Uso

A CLI Manthys pode ser usada com os seguintes comandos:

- `./manthys check`: Verifica o estado atual dos softwares necessários.
- `./manthys install`: Instala todos os softwares que não estão presentes no sistema, de acordo com a checagem.

### Exemplos:

- Para verificar os softwares:
`./manthys check`

- Para instalar os softwares:
`./manthys install`


## Contribuição

Contribuições são bem-vindas! Por favor, leia o arquivo CONTRIBUTING.md para mais informações sobre como contribuir para o projeto.

## Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE.md para mais detalhes.
