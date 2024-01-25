#!/bin/bash

# Definindo cores para o output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # Sem cor

echo -e "${GREEN}Iniciando a instalação da CLI Manthys...${NC}"

# Verifica se o Go está instalado
if ! command -v go &> /dev/null
then
    echo -e "${RED}Go não encontrado no sistema.${NC}"

    # Solicita confirmação do usuário
    read -p "Deseja instalar o Go? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]
    then
        echo -e "${GREEN}Instalando Go...${NC}"
        sudo apt update && sudo apt install -y golang-go
    else
        echo -e "${RED}Instalação cancelada pelo usuário.${NC}"
        exit 1
    fi
else
    echo -e "${GREEN}Go encontrado, prosseguindo com a instalação.${NC}"
fi

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
