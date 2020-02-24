/*
 * Ejemplo de chaincode para el Hackathon de la España Vacía 2019 Musac
 * Equipo 7 - LBH
 * Chaincode Filandon
 */

package main

/* Imports
 * Librerías para formateo de cadenas, manejo de lectura y escritura de JSON
 */

 /* Importante
 * Hay que hacer vendoring de las librerias que se añadan a un proyecto HLF
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
//Librería de google para generar identificadores únicos
	"github.com/google/uuid"

//Librerías especificas de Hyperledger Fabric, usadadas para la gestión de los Smart Contracts (ChainCodes)
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Estructura básica que hace referencia al SmartContract / Chaincode
type ChaincodeFilandon struct {
}

/*
 * Estructuras de Datos
 */

type fileLog struct {
	UUID             string `json:"uuid"`
	Creator       	 string `json:"creator"`
	FileHash         string `json:"fileHash"`
	Recipient        string `json:"recipient"`
	FileName         string `json:"fileName"`
	TransferComplete bool   `json:"transferComplete"`
	CreationTime     string `json:"creationTime"`
	CompletionTime   string `json:"completionTime`
}

type bienCultural struct {
	UUID             string `json:"uuid"`
	Creator       	 string `json:"creator"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	TypeItem         string `json:"typeitem"`
	Category         string `json:"category"`
	Keywords         string `json:"keyswords"`
	Place            string `json:"place`
	FileHash         string `json:"fileHash"`	
	CreationTime     string `json:"creationTime"`
}

type acuerdo struct {
	UUID             string `json:"uuid"`
	Creator       	 string `json:"creator"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	TypeItem         string `json:"typeitem"`
	Category         string `json:"category"`
	Conditions       string `json:"conditions`
	FileHash         string `json:"fileHash"`	
	CreationTime     string `json:"creationTime"`
}

/*
 * Una buena práctica es inicializar cualquier Ledger en una función separada - ver método initLedger()
 */
func (s *ChaincodeFilandon) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	fmt.Println("########### Iniciando Chaincode Filandon ###########")

	return shim.Success(nil)
}

/*
 * El metodo Invoke se llama como resultado de una aplicación haciendo una petición al Smart Contract
 * La llamada de la aplicación tambien especifica la función concreta del Smart Contract
 * Redirecciona al método que se encarga de manejar la llamada de la invocación
 */

func (s *ChaincodeFilandon) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Recupera los métodos y argumentos del Smart Contract según el comando invocado
	function, args := APIstub.GetFunctionAndParameters()
	fmt.Println("Se ha invocado la función: " + function)


	// Redirecciona al método invocado para interactuar con el ledger del canal
	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "crearTransaccion" {
		return s.crearTransaccion(APIstub, args)
	} else if function == "consultarTransaccion" {
		return s.consultarTransaccion(APIstub, args)
	} else if function == "consultarTransaccionPorReceptor" {
		return s.consultarTransaccionPorReceptor(APIstub, args)
	} else if function == "consultarTransaccionPorCreator" {
		return s.consultarTransaccionPorCreator(APIstub, args)
	} else if function == "crearBIC" {
		return s.crearBIC(APIstub, args)
	} else if function == "crearAcuerdo" {
		return s.crearAcuerdo(APIstub, args)
	} else if function == "cambiarDescripcionBIC" {
		return s.cambiarDescripcionBIC(APIstub, args)
	}		
  	// Mostramos mensaje de error si no hay ningún método asociado
	return shim.Error("Nombre invalido de funcion del smart contract")
}


// Configura el estado inicial del ledger
// Valores que queremos que tenga, objetos de prueba
func (s *ChaincodeFilandon) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// ======================== crearTransaccion =================================================
// crearTransaccion registra un archivo en la red HLF, con un hash del documento offchain
// args[0]: Creador
// args[1]: hash offchain del ficheron (probablemente en IPFS)
// args[2]: Destinatario
// args[3]: Nombre de fichero
// =========================================================================================
func (s *ChaincodeFilandon) crearTransaccion(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Número incorrecto de argumentos. Son necesarios 4 argumentos")
	}

	//Generamos un nuevo UUID para cada transacción
	id, err := uuid.NewUUID()
	if err != nil {
		// Capturar y manejar el error
		return shim.Error("Fallo al generar el UUID del elemento")
	}
	uuid := id.String()

	creator := args[0]
	fileHash := args[1]
	recipient := args[2]
	filename := args[3]
	//Capturamos el tiempo actual para introducirlo como fecha de creacion
	now := time.Now().String()
	creationTime := now[:19]
	completionTime := ""

	if err != nil {
		return shim.Error("Fallo al ejecutar la transacción: " + err.Error())
	}

	var transaccion = fileLog{
		UUID:             uuid,
		Creator:       	  creator,
		FileHash:         fileHash,
		Recipient:        recipient,
		FileName:         filename,
		TransferComplete: false,
		CreationTime:     creationTime,
		CompletionTime:   completionTime}

	transaccionAsBytes, _ := json.Marshal(transaccion)

	APIstub.PutState(uuid, transaccionAsBytes)

	return shim.Success(nil)
}

// ======================== crearBIC =================================================
// crearBIC registra un Bien de Interes Cultural (Material o Inmaterial)
// args[0]: Creador
// args[1]: hash offchain del ficheron (probablemente en IPFS)
// args[2]: Nombre del BIC
// args[3]: Descripción del BIC
// args[4]: Tipo del BIC
// args[5]: Categoria del BIC
// args[6]: Palabras Clave
// args[7]: Lugar del BIC
// =========================================================================================
func (s *ChaincodeFilandon) crearBIC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 8 {
		return shim.Error("Número incorrecto de argumentos. Son necesarios 8 argumentos")
	}

	//Generamos un nuevo UUID para cada transacción
	id, err := uuid.NewUUID()
	if err != nil {
		// Capturar y manejar el error
		return shim.Error("Fallo al generar el UUID del elemento")
	}
	uuid := id.String()

	creator := args[0]
	fileHash := args[1]
	name := args[2]
	description := args[3]
	typeitem := args[4]
	category := args[5]
	keyswords := args[6]
	place := args[7]
	//Capturamos el tiempo actual para introducirlo como fecha de creacion
	now := time.Now().String()
	creationTime := now[:19]

	if err != nil {
		return shim.Error("Fallo al ejecutar la transacción: " + err.Error())
	}

	var transaccion = bienCultural{
		UUID:             uuid,
		Creator:       	  creator,
		FileHash:         fileHash,
		Description:	  description,
		TypeItem:		  typeitem,
		Name:             name,
		Category:		  category,
		Keywords:		  keyswords,
		Place:			  place,
		CreationTime:     creationTime}

	transaccionAsBytes, _ := json.Marshal(transaccion)

	APIstub.PutState(uuid, transaccionAsBytes)

	return shim.Success(nil)
}

// ======================== crearAcuerdo =================================================
// crearAcuerdo registra un Acuerdo en la Blockchain
// args[0]: Creador
// args[1]: hash offchain de un fichero vinculado (probablemente en IPFS)
// args[2]: Nombre del Acuerdo
// args[3]: Descripción del Acuerdo
// args[4]: Tipo del Acuerdo
// args[5]: Categoria del Acuerdo
// args[6]: Condiciones del Acuerdo
// =========================================================================================
func (s *ChaincodeFilandon) crearAcuerdo(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Número incorrecto de argumentos. Son necesarios 7 argumentos")
	}

	//Generamos un nuevo UUID para cada transacción
	id, err := uuid.NewUUID()
	if err != nil {
		// Capturar y manejar el error
		return shim.Error("Fallo al generar el UUID del elemento")
	}
	uuid := id.String()

	creator := args[0]
	fileHash := args[1]
	name := args[2]
	description := args[3]
	typeitem := args[4]
	category := args[5]
	conditions := args[6]
	//Capturamos el tiempo actual para introducirlo como fecha de creacion
	now := time.Now().String()
	creationTime := now[:19]

	if err != nil {
		return shim.Error("Fallo al ejecutar la transacción: " + err.Error())
	}

	var transaccion = acuerdo{
		UUID:             uuid,
		Creator:       	  creator,
		FileHash:         fileHash,
		Description:	  description,
		TypeItem:		  typeitem,
		Name:             name,
		Category:		  category,
		Conditions:		  conditions,
		CreationTime:     creationTime}

	transaccionAsBytes, _ := json.Marshal(transaccion)

	APIstub.PutState(uuid, transaccionAsBytes)

	return shim.Success(nil)
}


// ======================== Métodos Update =================================================

// ======================== cambiarDescripcionBIC ==========================================
// Cambia la descripción de un Bien de Interes Cultural
// args[0]: Clave identificativa del BIC
// args[1]: Nueva Descripción
// =========================================================================================

func (s *ChaincodeFilandon) cambiarDescripcionBIC(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Número incorrecto de argumentos. Se requieren 2 (Clave, Descripción)")
	}

	bicAsBytes, _ := APIstub.GetState(args[0])
	bic := bienCultural{}

	json.Unmarshal(bicAsBytes, &bic)
	bic.Description = args[1]

	bicAsBytes, _ = json.Marshal(bic)
	APIstub.PutState(args[0], bicAsBytes)

	return shim.Success(bicAsBytes)

}

// ======================== Métodos de Consulta =================================================


// ======================== consultarTransaccion =================================================
// consultarTransaccion busca (consulta) por la transacción que contenga la clave especificada.
// args[0]: clave (key) del registro que se desea buscar
// =========================================================================================
func (s *ChaincodeFilandon) consultarTransaccion(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Numero incorrecto de argumentos. Se espera 1 argumento solo")
	}

	fmt.Println("------------- Consultando Transacción -------------")
	transaccionAsBytes, _ := APIstub.GetState(args[0])

	// Capturar y manejar el error de nulo en la clave 
	if transaccionAsBytes == nil {
		return shim.Error("No se puede encontrar la clave seleccionada")
	}
	return shim.Success(transaccionAsBytes)
}


// ============= consultarTransaccionPorCreator =================================================
// consultarTransaccionPorCreator consulta por transacciones que tengan como creator el argumento proporcionado.
// Esta consulta solo esta disponible en bases de datos que soporten consultas enriquecidas (CoudDB)
// =========================================================================================
func (t *ChaincodeFilandon) consultarTransaccionPorCreator(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	creatorName := strings.ToLower(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"creator\":\"%s\"}}", creatorName)

	resultsIterator, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Registro\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
 	// Imprimimos por pantalla los resultados del buffer de la consulta
	fmt.Printf("- consultarTransaccionPorCreator:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// ============= consultarTransaccionPorReceptor =================================================
// consultarTransaccionPorReceptor consulta por transacciones que tengan como receptor el argumento proporcionado.
// Esta consulta solo esta disponible en bases de datos que soporten consultas enriquecidas (CoudDB)
// =========================================================================================
func (t *ChaincodeFilandon) consultarTransaccionPorReceptor(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	//Comprobamos si se ha pasado el número correcto de argumentos (1)
	if len(args) != 1 {
		return shim.Error("Numero incorrecto de argumentos. Se espera solo un argumento 1")
	}

	//Pasamos a minusculas el argumento
	recipientName := strings.ToLower(args[0])

	//Formateamos la consulta con el selector 
	queryString := fmt.Sprintf("{\"selector\":{\"recipient\":\"%s\"}}", recipientName)

	resultsIterator, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Registro\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- consultarTransaccionPorReceptor:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// La función main es solo relevante para testeo y unitario
func main() {

	// Creando el nuevo Chain Code
	err := shim.Start(new(ChaincodeFilandon))
	if err != nil {
		fmt.Printf("Error creando el código de cadena: %s", err)
	}
}
