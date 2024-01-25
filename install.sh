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

# Check if Curl is installed
if ! command -v curl &> /dev/null; then
    echo -e "${RED}Curl not found.${NC}"
    read -p "Do you want to install Curl? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]
    then
        echo -e "${GREEN}Installing Curl... Please wait.${NC}"
        sudo apt-get install curl -y > /dev/null 2>&1
    else
        echo -e "${RED}Installation cancelled by the user.${NC}"
        exit 1
    fi
else
    echo -e "${GREEN}Curl found.${NC}"
fi

# Check if Wget is installed
if ! command -v wget &> /dev/null; then
    echo -e "${RED}Wget not found.${NC}"
    read -p "Do you want to install Wget? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]
    then
        echo -e "${GREEN}Installing Wget... Please wait.${NC}"
        sudo apt-get install wget -y > /dev/null 2>&1
    else
        echo -e "${RED}Installation cancelled by the user.${NC}"
        exit 1
    fi
else
    echo -e "${GREEN}Wget found.${NC}"
fi

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo -e "${RED}Go not found on the system.${NC}"
    read -p "Do you want to install Go? (y/n): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]
    then
        echo -e "${GREEN}Installing Go... Please wait.${NC}"
        sudo apt update > /dev/null 2>&1 && sudo apt install -y golang-go > /dev/null 2>&1
    else
        echo -e "${RED}Installation cancelled by the user.${NC}"
        exit 1
    fi
fi