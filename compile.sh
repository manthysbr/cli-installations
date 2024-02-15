#!/bin/bash

# Definindo cores para saida
GREEN='\033[0;32m'
NC='\033[0m' # Sem cor

# Nome do seu aplicativo
APP_NAME="manthys"

# Versão do aplicativo (isso pode ser dinâmico, talvez passado como um argumento ou obtido de alguma forma automatizada)
APP_VERSION="v1.0.0"

# Caminho do diretório de build (pode ser um diretório temporário)
BUILD_DIR="./build"

# Função para criar efeito de digitação
typeout() {
  local IFS=''
  for i in $1; do
    echo -n "$i"
    sleep 0.05
  done
  echo
}

# Criando o diretório de build se não existir
mkdir -p $BUILD_DIR

# Mensagem inicial
typeout "${GREEN}Iniciando o processo de build e empacotamento para $APP_NAME $APP_VERSION...${NC}"

# Compilando o aplicativo
echo -e "${GREEN}Compilando o binário...${NC}"
go build -o $BUILD_DIR/$APP_NAME

if [ $? -ne 0 ]; then
    echo -e "${RED}Erro na compilação do binário. Processo abortado.${NC}"
    exit 1
fi

# Navegando para o diretório de build
cd $BUILD_DIR

# Empacotando o binário
echo -e "${GREEN}Empacotando o binário...${NC}"
tar -czvf $APP_NAME-$APP_VERSION.tar.gz $APP_NAME

# Gerando o checksum
echo -e "${GREEN}Gerando checksum...${NC}"
sha256sum $APP_NAME-$APP_VERSION.tar.gz > $APP_NAME-$APP_VERSION.sha256

# Mensagem final
typeout "${GREEN}Processo de build e empacotamento concluído com sucesso!${NC}"

# Retornando ao diretório original
cd -

# Mostrando os arquivos resultantes
echo -e "${GREEN}Arquivos gerados:${NC}"
ls -l $BUILD_DIR
