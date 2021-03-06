/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main


import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response  {
        fmt.Println("########### example_cc Init ###########")
	_, args := stub.GetFunctionAndParameters()
	var A, B, C, D, E ,P string    // Entities
	var Aval, Bval, Cval, Dval, Eval ,Pval  int // Asset holdings
	var err error

	if len(args) != 12 {
		return shim.Error("Incorrect number of arguments. Expecting 12")
	}

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	C = args[4]
	Cval, err = strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	D = args[6]
	Dval, err = strconv.Atoi(args[7])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	E = args[8]
	Eval, err = strconv.Atoi(args[9])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	P = args[10]
	Eval, err = strconv.Atoi(args[11])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}	
	fmt.Printf("Aval = %d, Bval = %d, Cval = %d, Dval = %d, Eval = %d, Pval = %d\n", Aval, Bval, Cval, Dval, Eval, Pval)

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(C, []byte(strconv.Itoa(Cval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(D, []byte(strconv.Itoa(Dval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(E, []byte(strconv.Itoa(Eval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(P, []byte(strconv.Itoa(Pval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)


}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface) pb.Response {
		return shim.Error("Unknown supported call")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
        fmt.Println("########### example_cc Invoke ###########")
	function, args := stub.GetFunctionAndParameters()
	
	if function != "invoke" {
                return shim.Error("Unknown function call")
	}

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting at least 2")
	}

	if args[0] == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}

	if args[0] == "query" {
		// queries an entity state
		return t.query(stub, args)
	}
	if args[0] == "move" {
		// Deletes an entity from its state
		return t.move(stub, args)
	}
	return shim.Error("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'")
}

func (t *SimpleChaincode) move(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// must be an invoke
	var A, B, C, D, E, P string    // Entities
	var Aval, Bval, Cval, Dval, Eval, Pval int // Asset holdings
	var X,Y,Z,W,U,V int          // Transaction value
	var err error

	if len(args) != 13 {
		return shim.Error("Incorrect number of arguments. Expecting 13, function followed by 6 names and 6 value")
	}

	A = args[1]
	B = args[2]
	C = args[3]
	D = args[4]
	E = args[5]
	P = args[6]
	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Bvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	Cvalbytes, err := stub.GetState(C)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Cvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Cval, _ = strconv.Atoi(string(Cvalbytes))

	Dvalbytes, err := stub.GetState(D)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Dvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Dval, _ = strconv.Atoi(string(Dvalbytes))
	
	Evalbytes, err := stub.GetState(E)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Evalbytes == nil {
		return shim.Error("Entity not found")
	}
	Eval, _ = strconv.Atoi(string(Evalbytes))

	Pvalbytes, err := stub.GetState(P)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Pvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Pval, _ = strconv.Atoi(string(Pvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[7])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	Y, err = strconv.Atoi(args[8])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	Z, err = strconv.Atoi(args[9])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	W, err = strconv.Atoi(args[10])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	U, err = strconv.Atoi(args[11])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	V, err = strconv.Atoi(args[12])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}

/*
	Aval = Aval + X
	Bval = Bval + Y
	Cval = Cval + Z
	Dval = Dval + W
	Eval = U
	Pval = V
*/
	Aval =  X
	Bval =  Y
	Cval =  Z
	Dval =  W
	Eval =  U
	Pval =  V

	fmt.Printf("Aval = %d, Bval = %d, Cval = %d, Dval = %d, Eval = %d, Price = %d\n", Aval, Bval, Cval, Dval, Eval, Pval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(C, []byte(strconv.Itoa(Cval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(D, []byte(strconv.Itoa(Dval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(E, []byte(strconv.Itoa(Eval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(P, []byte(strconv.Itoa(Pval)))
	if err != nil {
		return shim.Error(err.Error())
	}

        return shim.Success(nil);
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[1]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var A string // Entities
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[1]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
