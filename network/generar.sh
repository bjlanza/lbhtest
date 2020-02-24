#!/bin/bash
#
# SPDX-License-Identifier: Apache-2.0
#
# Finalizar al primer error, mostrar todos los mensajes de error
set -e
#Enlazamos las rutas a los binarios de Hyperledger Fabric
export PATH=${PWD}/bin:${PWD}:$PATH
export FABRIC_CFG_PATH=$PWD
CURRENT_DIR=$PWD

# Se borran las carpetas que contienen los crypto certificados y artefactos
rm -rf crypto-config
rm -rf channel-artifacts
mkdir channel-artifacts
mkdir crypto-config

export CHANNEL_NAME=canalfilandon # Nombre que se desea para el canal
export CHANNEL_PROFILE=TwoOrgsChannel #Uno de los perfiles de canales definidos en el fichero configtx.yaml
export NETWORK_CONFIG=./network-config

echo 'Generando Certificados...'
# Se usa la herramienta propia de hyperledger fabric que consume el fichero crypto-config.yaml
cryptogen generate --config=./network-config/crypto-config.yaml
# Creando el bloque genesis del sistema de ordenación - Que contiene la definición de una red HLF
# Describe el consorcio de la de red y que organizaciónes pueden crear un canal
configtxgen -profile ThreeOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
# Crear Transaciones de Canales
# El canal filandon uso la definición de ThreeOrgsChannel que incluye los 3 pueblos
configtxgen  -profile ThreeOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID canalfilandon

# Configuración de los Pares de Anclaje (Anchor Peers)
echo -e "\n\e[93m[step-5]\e[39m: Configurando Anchor Peers para cada una de las Orgs pertenecientes al  $CHANNEL_NAME con el perfil $CHANNEL_PROFILE...\e[37m\n"

configtxgen -profile $CHANNEL_PROFILE -outputAnchorPeersUpdate ./channel-artifacts/pueblo1Anchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
# Finalizando
exit 1
