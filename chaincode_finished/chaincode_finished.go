/*
Copyright IBM Corp 2016 All Rights Reserved.

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
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type Patient struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	DateOfBirth string `json:"DateOfBirth"`
	Sex         string `json:"Sex"`
	PhoneNumber string `json:"PhoneNumber"`
	SourceId    string `json:"SourceId"`
	EntityId    string `json:"EntityId"`
	Meds
}

type Medication struct {
	MedName           string `json:"MedName"`
	Dosage            string `json:"Sosage"`
	PrescriptionStart string `json:"PrescriptionStart"`
	PrescriptionEnd   string `json:"PresriptionEnd"`
}

type Meds struct {
	meds []Medication
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//	if len(args) != 1 {
	//		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	//	}
	//
	//	err := stub.PutState("hello_world", []byte(args[0]))
	//	if err != nil {
	//		return nil, err
	//	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	if function == "test" { //read a variable
		return []byte("IT WORKS!!"), nil
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var key, value string
	var err error
	fmt.Println("running write()")

	// if len(args) != 2 {
	// 	return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	// }

	var patient Patient
	key = args[0] //rename for funsies
	value = args[1]

	bytes, err := stub.GetState(key)

	if err != nil {

	} else {
		err = json.Unmarshal(bytes, &patient)
		if err != nil {
			return nil, errors.New("Corrupt Patient record")
		}
	}

	patient.FirstName = key
	patient.PhoneNumber = value

	fmt.Println("Before Patient Marshall:")
	fmt.Println(patient)

	bytes, err = json.Marshal(patient)

	fmt.Println("Marshalled Patient:")
	fmt.Println(bytes)

	var patient15 Patient
	json.Unmarshal(bytes, &patient15)

	fmt.Println("Patient Unmarshalled:")
	fmt.Println(patient15.FirstName)
	if err != nil {
		return nil, errors.New("Error creating Patient record")
	}

	err = stub.PutState(key, bytes)
	if err != nil {
		return nil, err
	}

	bytes2, err := stub.GetState(key)

	var patient2 Patient
	json.Unmarshal(bytes2, &patient2)

	fmt.Println("From GetState:")
	fmt.Println(patient2.FirstName)

	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	// if len(args) != 1 {
	// 	return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	// }

	key = args[0]
	jsonval, err := stub.GetState(key)
	var patient Patient
	json.Unmarshal(jsonval, &patient)
	valAsbytes := []byte(patient.PhoneNumber)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

// func NewPatient() interface{} {
// 	return new(Patient)
// }
