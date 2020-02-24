#!/bin/bash
#
# SPDX-License-Identifier: Apache-2.0
#

# Installando Chaincode de tests y pruebas
#echo 'Instalando test Chaincode mycc en todos los Peers..'

#sudo docker exec -it cli peer chaincode install -n mycc -p github.com/chaincode -v v0
#sudo docker exec -it cli2 peer chaincode install -n mycc -p github.com/chaincode -v v0
#sudo docker exec -it cli3 peer chaincode install -n mycc -p github.com/chaincode -v v0

#echo 'Instanciando el Chaincode de testeo mycc en canalFilandon..'

#sudo docker exec -it cli peer chaincode instantiate -o orderer.example.com:7050 -C canalfilandon -n mycc github.com/chaincode -v v0 -c '{"Args": ["a", "100"]}'

# Definimos las variables necesarias para el CC principal
CC_SRC_PATH="github.com/chaincode/karnacc"
CHANNEL_NAME="canalfilandon"
CC_RUNTIME_LANGUAGE="golang"
VERSION="2.1"
NAME="karnacc"

# Instalando e Instanciando Chaincode
echo 'Instalando Chaincode Filandon en el peer...'
docker exec cli peer chaincode install -n "$NAME" -v "$VERSION" -p "$CC_SRC_PATH" -l "$CC_RUNTIME_LANGUAGE"
echo 'Instanciando Chaincode Filandon en el canal canalfilandon...'
docker exec cli peer chaincode instantiate -o orderer.example.com:7050 -C "$CHANNEL_NAME" -n "$NAME" -l "$CC_RUNTIME_LANGUAGE" -v "$VERSION" -c '{"Args":[]}' -P "OR ('Org1MSP.member')"

# Actualizar
# docker exec cli peer chaincode upgrade -n filandoncc -v 2.0 -p github.com/chaincode/filandoncc -C canalfilandon -o orderer.example.com:7050 -c '{"Args":[""]}'

# Finalización de Script
echo -e '\n\e[92m \u2714 Todas la tareas realizadas.. Para consultar el MVP usar <<consultar.sh>>>...\e[39m'
exit 1
# EoF 

