#!/bin/bash
#
# SPDX-License-Identifier: Apache-2.0
#

# Test de la red HLF

sudo docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 -n mycc -c '{"Args":["set", "a", "20"]}' -C canalfilandon

#sleep 5

#sudo docker exec -it cli peer chaincode query -n mycc -c '{"Args":["query","a"]}' -C canalfilandon
#sudo docker exec -it cli2 peer chaincode invoke -o orderer.example.com:7050 -n mycc -c '{"Args":["set", "a", "40"]}' -C canalfilandon

#sleep 5

#sudo docker exec -it cli2 peer chaincode query -n mycc -c '{"Args":["query","a"]}' -C canalfilandon
#sudo docker exec -it cli3 peer chaincode invoke -o orderer.example.com:7050 -n mycc -c '{"Args":["set", "a", "60"]}' -C canalfilandon

#sleep 5

echo 'Querying For Result on Pueblo1 Peer'

sudo docker exec -it cli peer chaincode query -n mycc -c '{"Args":["query","a"]}' -C canalfilandon

echo 'Invocando y consultando el Chaincode Filandon'

sudo docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 -n filandoncc -c '{"Args":["crearTransaccion", "hackaton", "Qwzsrerewredfwerwerews", "hyper", "test01"]}' -C canalfilandon
sudo docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 -n filandoncc -c '{"Args":["crearBIC", "bjlanza", "Qwzs4343dsgdg4rbths", "Filandon", "El filandón es una reunión que se realiza por las noches una vez terminada la cena, en la que se cuentan en voz alta cuentos al tiempo que se trabaja en alguna labor manual.", "BIC", "tradición oral","cuentacuentos","León, Asturias"]}' -C canalfilandon

#Consultas
sudo docker exec -it cli peer chaincode query -n filandoncc -c '{"Args":["consultarTransaccionPorCreator","hackaton"]}' -C canalfilandon
sudo docker exec -it cli peer chaincode query -n filandoncc -c '{"Args":["consultarTransaccionPorReceptor","hyper"]}' -C canalfilandon

#sudo docker exec -it cli peer chaincode query -n filandoncc -c '{"Args":["consultarTransaccion","d2d14d9d-0368-11ea-8dc6-0242ac12000f"]}' -C canalfilandon
sudo docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 -n filandoncc -c '{"Args":["cambiarDescripcionBIC", "2120c7e7-0386-11ea-b890-0242ac1a000b", "Filandon Chain"]}' -C canalfilandon


#Consultar estado de la red blockchain
sudo docker exec -it cli peer channel fetch newest canalfilandon.block -c canalfilandon --orderer orderer.example.com:7050
#Comprobamos el estado en varios Peers (Pueblos) para consultar que se ha actualizado la red
sudo docker exec -it cli peer channel getinfo -c canalfilandon
sudo docker exec -it cli3 peer channel getinfo -c canalfilandon

echo -e '\n\e[92m \u2714 Todas las tareas realizadas... \e[39m'

exit 1

sudo docker exec -it cli peer chaincode invoke -o orderer.example.com:7050 -n filandoncc -c '{"Args":["crearAcuerdo", "Villablino", "Qwz4683246823u47384", "Acuerdo Vecinal", "Los vecinos acuerdan ser un municipio verde", "Propuesta de Ley", "Bando Municipal","cuentacuentos"]}' -C canalfilandon
sudo docker exec -it cli peer chaincode query -n filandoncc -c '{"Args":["consultarTransaccionPorCreator","Villablino"]}' -C canalfilandon
