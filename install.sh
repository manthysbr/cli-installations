#!/bin/bash

# Defining colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
BOLD='\033[1m'
BLINK='\033[5m'
NC='\033[0m' # Sem cor

# Função para criar efeito de digitação
typeout() {
  local IFS=''
  for i in $1; do
    echo -n "$i"
    sleep 0.05
  done
  echo
}

typeout "${GREEN}${BOLD}Iniciando a instalação da CLI Manthys...${NC}"

# Função para instalar pacotes com feedback simplificado
install_package() {
    package_name=$1
    echo -e "${GREEN}Verificando a instalação de $package_name...${NC}"
    if ! command -v $package_name &> /dev/null; then
        echo -e "${RED}$package_name não encontrado.${NC}"
        read -p "Deseja instalar $package_name? (y/n): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            echo -e "${GREEN}Instalando $package_name... Por favor, aguarde.${NC}"
            sudo apt-get install $package_name -y > /dev/null 2>&1 && echo -e "${GREEN}$package_name instalado com sucesso.${NC}"
        else
            echo -e "${RED}Instalação cancelada pelo usuário.${NC}"
            exit 1
        fi
    else
        echo -e "${GREEN}$package_name já está instalado.${NC}"
    fi
}

# Instala Curl, Wget e Go
install_package curl
install_package wget
install_package golang-go

echo -e "${GREEN}Compilando o código fonte...${NC}"
go build -o manthys
echo -e "${GREEN}Movendo o binário para /usr/local/bin...${NC}"
sudo mv manthys /usr/local/bin/manthys

# Verifica se /usr/local/bin está no PATH
if [[ ":$PATH:" != *":/usr/local/bin:"* ]]; then
    echo -e "${GREEN}Atualizando o PATH para incluir /usr/local/bin...${NC}"
    echo "export PATH=\$PATH:/usr/local/bin" >> ~/.bashrc
    source ~/.bashrc
fi

echo -e "${GREEN}Instalação concluída! Execute 'manthys' para começar.${NC}"
