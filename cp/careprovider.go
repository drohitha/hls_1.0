package main

import (

"fmt"

"encoding/json"
"github.com/hyperledger/fabric/protos/peer"

"github.com/hyperledger/fabric/core/chaincode/shim"
)

var EVENT_COUNTER = "event_counter"
type ManageCareProvider struct {
}
var CareProviderIndexStr = "_CareProviderindex"
type Patient struct{             // Attributes of a Patient      
  PatientID string `json:"PatientID"`
  PatientName string `json:"PatientName"`
  Address   string `json:"Address"`         
  Problems string `json:"Problems"`
  Gender string `json:"Gender"`
  PatientMobile string `json:"PatientMobile"`
  Medications string `json:"Medications"`
  Remarks string `json: "Remarks"`
  PatientEmail string `json: "PatientEmail"`
  User string `json: "User"`
  IStatus string `json: "IStatus"`
  }

  func main() {     
  err := shim.Start(new(ManageCareProvider))
  if err != nil {
    fmt.Printf("Error starting ManageCareProvider chaincode: %s", err)
  }
  }

func (t *ManageCareProvider) Init(stub shim.ChaincodeStubInterface) peer.Response {

  args := stub.GetStringArgs()
  var msg string
  var err error
  if len(args) != 1 {
    return shim.Error("Incorrect number of arguments. Expecting 1")
  }
  // Initialize the chaincode
  msg = args[0]
  fmt.Println("ManageCareProvider chaincode is deployed successfully.");
  
  // Write the state to the ledger
  err = stub.PutState("abc", []byte(msg))       //making a test var "abc", I find it handy to read/write to it right away to test the network
  if err != nil {
    return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
  }
  
  var empty []string
  jsonAsBytes, _ := json.Marshal(empty)               //marshal an emtpy array of strings to clear the index
  err = stub.PutState(CareProviderIndexStr, jsonAsBytes)
  if err != nil {
    return shim.Error(fmt.Sprintf("Failed to create asset in patientindex: %s", args[0]))
  }
  err = stub.PutState(EVENT_COUNTER, []byte("1"))
  if err != nil {
    return shim.Error(fmt.Sprintf("Failed to create asset in event counter: %s", args[0]))
  }
  return shim.Success(nil)
}

func (t *ManageCareProvider) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
   // fmt.Println("invoke is running " + function)
fn, args := stub.GetFunctionAndParameters()
  // Handle different functions
        var result string
    var err error
  // Handle different functions
  if fn == "getPatient_byID" {                         //initialize the chaincode state, used as reset
    result, err = getPatient_byID(stub, args)
  } else if fn == "cupdate_patient" {
    result, err = cupdate_patient(stub, args)
  } else if fn == "share_patient" {
    result, err = share_patient(stub,args)
  } else if fn == "get_byCareProviderID" {
    result, err = get_byCareProviderID(stub,args)
  }/* else if function == "update_istatus" {
    return t.update_istatus(stub, args)
  }*/

 if err != nil {
            return shim.Error(err.Error())
    }

   fmt.Println("invoke did not find func: " + fn)          //error
  
  return shim.Success([]byte(result))
}

func getPatient_byID(stub shim.ChaincodeStubInterface, args []string) (string, error) {
  if len(args) != 2 {
    return "", fmt.Errorf("Incorrect number of arguments. Expecting 3 args")
  }
  PatientChaincode := args[0]
  PatientID := args[1]
  f1 := "getPatient_byID"
 queryArgs1 := make([][]byte, 2)
    queryArgs1[0] = []byte(f1)
     queryArgs1[1] = []byte(PatientID)
  patientAsBytes, err := stub.InvokeChaincode(PatientChaincode, queryArgs1,"")
  if err != nil {
    errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", err.Error())
    fmt.Printf(errStr)
    return "", fmt.Errorf(errStr)
  }
  res := Patient{}
  json.Unmarshal(patientAsBytes, &res)
  fmt.Println(res)
  if res.PatientID == PatientID {
    fmt.Println("Patient found with PatientID : " + PatientID)
  } else {
    return "", fmt.Errorf("PatientID not found")
  }
  return string(patientAsBytes),nil
}

func get_byCareProviderID(stub shim.ChaincodeStubInterface, args []string) (string, error) {
  if len(args) != 2 {
    return "", fmt.Errorf("Incorrect number of arguments. Expecting 3 args")
  }
  PatientChaincode := args[0]
  CareProviderID := args[1]
  f1 := "get_byCareProviderID"
   queryArgs1 := make([][]byte, 2)
    queryArgs1[0] = []byte(f1)
     queryArgs1[1] = []byte(CareProviderID)
  patientAsBytes, err := stub.InvokeChaincode(PatientChaincode, queryArgs1,"")
  if err != nil {
    errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", err.Error())
    fmt.Printf(errStr)
    return "", fmt.Errorf(errStr)
  }
  return string(patientAsBytes),nil
}

func cupdate_patient(stub shim.ChaincodeStubInterface, args []string) (string, error) {
  if len(args) != 5 {
    return "", fmt.Errorf("Incorrect number of arguments. Expecting 3 args")
  }
  PatientChaincode := args[0]
    PatientID := args[1]
    Medications := args[2]
    Remarks := args[3]
    User := args[4]
  f1 := "cupdate_patient"
  queryArgs1 := make([][]byte, 5)
    queryArgs1[0] = []byte(f1)
     queryArgs1[1] = []byte(PatientID)
     queryArgs1[2] = []byte(Medications)
     queryArgs1[3] = []byte(Remarks)
     queryArgs1[4] = []byte(User)
  _, err := stub.InvokeChaincode(PatientChaincode, queryArgs1,"")
  if err != nil {
    errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", err.Error())
    fmt.Printf(errStr)
    return "", fmt.Errorf(errStr)
  }
  
  return string(args[1]),nil
}

func share_patient(stub shim.ChaincodeStubInterface, args []string) (string, error) {
 fmt.Println("enter share function")
 if len(args) != 3 {
    return "", fmt.Errorf("Incorrect number of arguments. Expecting 1")
  }
  PatientChaincode := args[0]
  PatientID := args[1]
  DoctorID := args[2]
  f1 := "share_patient"
  queryArgs1 := make([][]byte, 3)
    queryArgs1[0] = []byte(f1)
     queryArgs1[1] = []byte(PatientID)
     queryArgs1[2] = []byte(DoctorID)
  _, err := stub.InvokeChaincode(PatientChaincode, queryArgs1)
  if err != nil {
    errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", err.Error())
    fmt.Printf(errStr)
    return "", fmt.Errorf(errStr)
  }
  
  return string(args[1]),nil
}
