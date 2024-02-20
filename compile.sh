#!/usr/bin/bash
# Esse script realiza o build e empacotamento da aplicação.
# O binario é compilado e empacotado em um arquivo tar.gz, e um checksum é gerado para o arquivo resultante.
# Definindo cores para saida
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # Sem cor

# Nome da app
APP_NAME="manthys"

# Versao da app
APP_VERSION="v1.0.0"

# Caminho do diretorio de build
BUILD_DIR="/home/docker/build"

# Verifica e cria o diretorio de build se nao existir
if [ ! -d "$BUILD_DIR" ]; then
  mkdir -p $BUILD_DIR
  if [ $? -ne 0 ]; then
      echo -e "${RED}Falha ao criar o diretorio '$BUILD_DIR'. Verifique as permissoes.${NC}"
      exit 1
  fi
fi

echo -e "${GREEN}Iniciando o processo de build e empacotamento para $APP_NAME $APP_VERSION...${NC}"

# Compilando o aplicativo...
if ! /usr/local/go/bin/go build -o $BUILD_DIR/$APP_NAME .; then
    echo -e "${RED}Erro na compilação do binário. Processo abortado.${NC}"
    exit 1
fi

# Empacotando o binário
echo -e "${GREEN}Empacotando o binario...${NC}"
tar -czvf $BUILD_DIR/$APP_NAME-$APP_VERSION.tar.gz -C $BUILD_DIR $APP_NAME

# Gerando o checksum
echo -e "${GREEN}Gerando checksum...${NC}"
sha256sum $BUILD_DIR/$APP_NAME-$APP_VERSION.tar.gz > $BUILD_DIR/$APP_NAME-$APP_VERSION.sha256

echo -e "${GREEN}Processo de build e empacotamento concluido com sucesso!${NC}"

# Mostrando os arquivos resultantes
ls -l $BUILD_DIR