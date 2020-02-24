 #
echo 'Apagando la Red Hyperledger Fabric...'

echo 'Limpiando contenedores Docker...'

sudo docker-compose -f docker-compose-cli.yml down --remove-orphans

# Borrando los contenedores docker antiguos
sudo docker volume prune

# Borrando las redes antiguas de docker
sudo docker network prune

echo 'Todas las tareas realizadas ..Adios..'
exit 1