#!/usr/bin/bash

# Definindo cores para saída
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # Sem cor

# Nome do seu aplicativo e versão
APP_NAME="manthys"
APP_VERSION="v1.0.0"
BIN_DIR="/usr/local/bin" # Diretório padrão para binários globais
TAR_FILE="${APP_NAME}-${APP_VERSION}.tar.gz" # Nome do arquivo tar.gz
TEMP_DIR="/tmp/install" # Diretório temporário para extração

# Verifica se o usuário é root, necessário para mover o binário para /usr/local/bin
if [ "$(id -u)" -ne 0 ]; then
    echo -e "${RED}Este script deve ser executado como root. Use sudo ./install.sh${NC}"
    exit 1
fi

# Criando diretório temporário para extração
mkdir -p $TEMP_DIR
if [ $? -ne 0 ]; then
    echo -e "${RED}Falha ao criar diretório temporário.${NC}"
    exit 1
fi

# Descompactando o arquivo tar.gz no diretório temporário
echo -e "${GREEN}Descompactando $TAR_FILE...${NC}"
tar -xzvf $TAR_FILE -C $TEMP_DIR
if [ $? -ne 0 ]; then
    echo -e "${RED}Falha ao descompactar $TAR_FILE.${NC}"
    exit 1
fi

# Ajuste este caminho conforme necessário, dependendo de onde o binário é extraído
CURRENT_BIN_PATH="$TEMP_DIR/$APP_NAME"

# Movendo o binário para o diretório no PATH
echo -e "${GREEN}Movendo $APP_NAME para $BIN_DIR...${NC}"
mv $CURRENT_BIN_PATH $BIN_DIR/$APP_NAME
if [ $? -ne 0 ]; then
    echo -e "${RED}Falha ao mover $APP_NAME para $BIN_DIR.${NC}"
    exit 1
fi

# Tornando o binário executável
echo -e "${GREEN}Tornando $APP_NAME executável...${NC}"
chmod +x $BIN_DIR/$APP_NAME

# Limpeza do diretório temporário
rm -rf $TEMP_DIR

echo -e "${GREEN}Instalação concluída! Execute '$APP_NAME' para começar.${NC}"