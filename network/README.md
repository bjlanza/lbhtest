# Testnet Hyperledger Fabric - Filandon

Red Hyperledger Fabric como proyecto para el Hackathon España Vaciada 2019

### Pasos

1. Descargar las ejecutables de Hyperledger Fabric recomendados


```

curl -sSL http://bit.ly/2ysbOFE | bash -s 1.4.2

```


2. Generar los artefactos de canal y los certificados crytográficos necesarios


```
./generar.sh
```

3.Iniciar la Red Hyperledger Fabric 


```
./start.sh
```

4.Instalar e Instanciar los Chaincodes

```
./instalarCC.sh
```

5. Consultar los datos de la red Hyperledger Fabric

```
./consultar.sh
```

6. Apagar la red y eliminar todos los Contenedores

```
./teardown.sh
```
