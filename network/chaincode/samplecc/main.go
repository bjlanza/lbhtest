//SPDX-License-Identifier: Apache-2.0
// Package main prov.go - Borja Lanza - 01/02/2020
package main

//Imports

import (
	//Básicas
	"fmt"
	//Fabric
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

)

// SampleChaincode es el nombre del chaincode que creamos para manejar un activo
type SampleChaincode struct {

}

/********************************************************
	DEFAULT/ PRE REQUIS FUNCTION FOR ACCESS TO PEER LEDGER
*********************************************************/

//Método INIT
func (c *SampleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	//log.Info("Chaincode Inicializado")
	return pb.Response{
		Status:  200,
		Message: "successfully initiated",
		Payload: nil,
	}
}

func (c *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// Extract function and args from transaction proposal
	funcName, args := stub.GetFunctionAndParameters()

	//fmt.Println("invoke is running " + funcName)

	//txID := stub.GetTxID()

	// check who's calling
	/*creator, err := stub.GetCreator()
	if err != nil {
		log.Printf("Obtained error: %s", err.Error())
		return shim.Error(err.Error())
	}
	log.Printf("Transaction from %s", creator)*/

	switch funcName {
	case "Query":
		return c.Query(stub, args[0])
	default:
		return shim.Error(fmt.Sprintf("unsupported function: %s", funcName))

	}
}

// Configura el estado inicial del ledger
// Valores que queremos que tenga, objetos de prueba
func (c *SampleChaincode) initLedger(APIstub shim.ChaincodeStubInterface) pb.Response {
	//Valores de Inicio
	return shim.Success(nil)
}

/********************************************************
	Metodos Auxiliares
*********************************************************/

func (c *SampleChaincode) Query(stub shim.ChaincodeStubInterface, uid string) pb.Response {
	stockBytes, err := stub.GetState(uid)
	if err != nil {
		return shim.Error("query err:"+err.Error())
	}
	if stockBytes == nil{
		return shim.Error("stock not exists")
	}
	return shim.Success(stockBytes)
}



/********************************************************
	Función Main (Principal) que arranca el chaincode en el contenedor durante la instanciación
*********************************************************/

func main() {
	//Configuración del Log
	//log.SetFormatter(&log.JSONFormatter{})

	err := shim.Start(new(SampleChaincode))
	if err != nil {
		fmt.Println("Could not start SampleChaincode")
	} else {
		fmt.Println("SampleChaincode successfully started")
	}
}
