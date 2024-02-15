#!/bin/bash

# Definindo cores para saída
GREEN='\033[0;32m'
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

typeout "${GREEN}Iniciando a instalação da CLI Manthys...${NC}"

# Função para instalar pacotes com feedback simplificado
install_package() {
    package_name=$1
    echo -e "Verificando a instalação de ${GREEN}$package_name${NC}..."
    if ! command -v $package_name &> /dev/null; then
        echo -e "${package_name} não encontrado."
        read -p "Deseja instalar $package_name? (y/n): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            echo -e "Instalando ${GREEN}$package_name${NC}... Por favor, aguarde."
            if sudo apt-get install $package_name -y > /dev/null 2>&1; then
                echo -e "${GREEN}$package_name instalado com sucesso.${NC}"
            else
                echo -e "Falha ao instalar $package_name. Verifique suas permissões ou conexão com a internet."
                exit 1
            fi
        else
            echo -e "Instalação cancelada pelo usuário."
            exit 1
        fi
    else
        echo -e "${GREEN}$package_name já está instalado.${NC}"
    fi
}

# Checando se o sudo pode ser acessado
if ! sudo -v &> /dev/null; then
    echo -e "Este script precisa de permissões de administrador para instalar pacotes."
    exit 1
fi

# Instala Curl, Wget e Go
install_package curl
install_package wget
install_package golang-go

# Compilando o código fonte
echo -e "Compilando o código fonte..."
if go build -o manthys; then
    echo -e "${GREEN}Compilação bem sucedida.${NC}"
else
    echo -e "Falha ao compilar o código fonte. Verifique se o Go está instalado corretamente."
    exit 1
fi

# Movendo o binário para /usr/local/bin
BIN_PATH="/usr/local/bin/manthys"
if sudo mv manthys $BIN_PATH; then
    echo -e "${GREEN}O binário foi movido com sucesso para $BIN_PATH${NC}"
else
    echo -e "Falha ao mover o binário para $BIN_PATH. Verifique suas permissões."
    exit 1
fi

# Atualizando o PATH, se necessário
if [[ ":$PATH:" != *":/usr/local/bin:"* ]]; then
    echo -e "Atualizando o PATH para incluir /usr/local/bin..."
    echo "export PATH=\$PATH:/usr/local/bin" >> ~/.bashrc
    # Recarregando o .bashrc
    source ~/.bashrc
    echo -e "${GREEN}PATH atualizado com sucesso.${NC}"
fi

echo -e "${GREEN}Instalação concluída! Execute 'manthys' para começar.${NC}"
