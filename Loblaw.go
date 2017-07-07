package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)



 
// ABC is a high level smart contract that ABCs' together business artifact based smart contracts
type ABC struct {

}

// UserDetails is for storing User Details

type UserDetails struct{	
	FfId string `json:"ffId"`
	Title string `json:"title"`
	Gender string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Dob string `json:"dob"`
	Email string `json:"email"`
	Country string `json:"country"`
	Address string `json:"address"`
	City string `json:"city"`
	Zip string `json:"zip"`
	CreatedBy string `json:"createdBy"`
	TotalPoint string `json:"totalPoint"`
}

// to return the verify result
type VerifyU struct{	
	Result string `json:"result"`
}

// Transaction is for storing transaction Details

type Transaction struct{	
	TrxId string `json:"trxId"`
	TimeStamp string `json:"timeStamp"`
	FfId string `json:"ffId"`
	Source string `json:"source"`
	Points string `json:"points"`
	Trxntype string `json:"trxntype"`
	TrxnSubType string `json:"trxnSubType"`
	Remarks string `json:"remarks"`
}


// GetMile is for storing retreived Get the total Points

type GetMile struct{	
	TotalPoint string `json:"totalPoint"`
}



// Asn is for storing ASN details

type Asn struct{	
	AsnNumber string `json:"asnNumber"`
	AsnUniqueid string `json:"asnUniqueid"`
	CreateTimestamp string `json:"createTimestamp"`
	UpdateTimestamp string `json:"updateTimestamp"`
	UpdatedBy string `json:"updatedBy"`
	Status string `json:"status"`
	CreatedBy string `json:"createdBy"`
}



// Item is for storing item details

type Item struct{	
	LineItemId string `json:"lineItemId"`
	ItemId string `json:"itemId"`
	Description string `json:"description"`
	Qty string `json:"qty"`
	Unit string `json:"unit"`
	Status string `json:"status"`
	QtyReceivedAtMedturn string `json:"qtyReceivedAtMedturn"`
	QtyReceivedAtWarehouse string `json:"qtyReceivedAtWarehouse"`
	QtyReceivedAtDisposal string `json:"qtyReceivedAtDisposal"`
	QtyReceivedAtManufacturer string `json:"qtyReceivedAtManufacturer"`
	CreateTs string `json:"createTs"`
	UpdateTs string `json:"updateTs"`
	UpdatedBy string `json:"updatedBy"`
	Remarks string `json:"remarks"`
	BoxBarcodeNumber string `json:"boxBarcodeNumber"`
	DebitMemo string `json:"debitMemo"`
	LotNumber string `json:"lotNumber"`
	Dc string `json:"dc"`
	Ndc string `json:"ndc"`
	ExpDate string `json:"expDate"`
	PurchageOrderNumber string `json:"purchageOrderNumber"`
	AsnNumber string `json:"asnNumber"`
	MrrRequestNumber string `json:"mrrRequestNumber"`
}


//to store Asn with item
type AsnItem struct{	
	AsnDetail Asn  `json:"asnDetail"`
	ItemDetail []Item `json:"itemDetail"`
}

//to store Asn with item
type ItemArray struct{	
	ItemDetail []Item `json:"itemDetail"`
}

// ITEMTrxnHistory is for storing item history transaction Details
type ItemTrxnHistory struct{	
	TrxId string `json:"trxId"`
	TimeStamp string `json:"timeStamp"`
	LineItemId string `json:"lineItemId"`
	Source string `json:"source"`
	Status string `json:"status"`
	Trxntype string `json:"trxntype"`
	TrxnSubType string `json:"trxnSubType"`
	Remarks string `json:"remarks"`
}

// ItemStatus is for storing item status Details
type ItemStatus struct{	
	Status_sht_med string `json:"STATUS_SHT_MED"`
	Status_rcvd_med string `json:"STATUS_RCVD_MED"`
	Status_inp_med string `json:"STATUS_INP_MED"`
	Status_sht_mf string `json:"STATUS_SHT_MF"`
	Status_rcvd_mf string `json:"STATUS_RCVD_MF"`
	Status_sht_dd string `json:"STATUS_SHT_DD"`
	Status_recv_dd string `json:"STATUS_RCVD_DD"`
	Status_destroy string `json:"STATUS_DESTORY"`
	Status_ret_wh string `json:"STATUS_RET_WH"`
	Status_rcvd_wh string `json:"STATUS_RCVD_WH"`
}

//to store item with item details
type ItemTrxnHistoryDetail struct{	
	Trxitemhistory []ItemTrxnHistory `json:"trxitemhistory"`
}


// Mrr is for storing Mrr details

type Mrr struct{	
	RequestNumber string `json:"requestNumber"`
	MrrUniqueid string `json:"mrrUniqueid"`
	CreateTimestamp string `json:"createTimestamp"`
	UpdateTimestamp string `json:"updateTimestamp"`
	UpdatedBy string `json:"updatedBy"`
	Status string `json:"status"`
	CreatedBy string `json:"createdBy"`
}


//to store Mrr with item
type MrrItem struct{	
	MrrDetail Mrr  `json:"mrrDetail"`
	ItemDetail []Item `json:"itemDetail"`
}





// Init initializes the smart contracts
func (t *ABC) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {



		// Check if table already exists
	_, err := stub.GetTable("UserDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("UserDetails", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ffId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "title", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "email", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "country", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "zip", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "totalPoint", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating UserDetails.")
	}

	
	// Check if table already exists
	_, err = stub.GetTable("Transaction")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("Transaction", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "trxId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "timeStamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ffId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "points", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "trxntype", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "trxnSubType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "remarks", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
	
	
	// Check if table already exists
	_, err = stub.GetTable("ASN")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}
	

	// Create application Table
	err = stub.CreateTable("ASN", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "asnNumber", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "asnUniqueid", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createTimestamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "updateTimestamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "updatedBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "payload", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating UserDetails.")
	}
	

		// Check if table already exists
	_, err = stub.GetTable("ITEM")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("ITEM", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "lineItemId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "itemId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "description", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "qty", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "unit", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "qtyReceivedAtMedturn", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "qtyReceivedAtWarehouse", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "qtyReceivedAtDisposal", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "qtyReceivedAtManufacturer", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createTs", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "updateTs", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "updatedBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "remarks", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "boxBarcodeNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "debitMemo", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lotNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dc", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ndc", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "expDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "purchageOrderNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "asnUniqueid", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ddrUniqueid", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "grmUniqueid", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "shUniqueid", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "asnNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "mrrRequestNumber", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}


	// Check if table already exists
	_, err = stub.GetTable("ASNTrxnHistory")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("ASNTrxnHistory", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "trxId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "timeStamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "asnUniqueid", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "trxntype", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "trxnSubType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "remarks", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
	
		
		// Check if table already exists
	_, err = stub.GetTable("ITEMTrxnHistory")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("ITEMTrxnHistory", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "trxId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "timeStamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lineItemId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "trxntype", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "trxnSubType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "remarks", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
	
	
	// Check if table already exists
	_, err = stub.GetTable("MRR")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("MRR", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "requestNumber", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "mmrUniqueid", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createTimestamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "updateTimestamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "updatedBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "payload", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating UserDetails.")
	}
	
	
	
	
		
		// Check if table already exists
	_, err = stub.GetTable("ITEMTempHumidTrxnHistory")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("ITEMTempHumidTrxnHistory", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "trxId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "timeStamp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lineItemId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "trxntype", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "trxnSubType", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "remarks", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
	
	
		// setting up the increment 
	stub.PutState("ASNincrement", []byte("1"))
	stub.PutState("DDRincrement", []byte("1"))
	stub.PutState("GRMincrement", []byte("1"))
	stub.PutState("SHUincrement", []byte("1"))	

	
		
	// setting up the users role
	stub.PutState("user_type1_1", []byte("WARE_HOUSE"))
	stub.PutState("user_type1_2", []byte("MEDTURN"))
	stub.PutState("user_type1_3", []byte("DISPOSAL"))
	stub.PutState("user_type1_4", []byte("MANUFACTURER"))


	// setting up the users role
	stub.PutState("STATUS_SHT_MED", []byte("Check-In"))
	stub.PutState("STATUS_RCVD_MED", []byte("Bag Tagging"))
	stub.PutState("STATUS_INP_MED", []byte("On Boarding"))
	stub.PutState("STATUS_SHT_MF", []byte("Off Boarding and Transfer"))
	stub.PutState("STATUS_RCVD_MF", []byte("On-Boarding"))
	stub.PutState("STATUS_SHT_DD", []byte("Off-Boarding"))
	stub.PutState("STATUS_RCVD_DD", []byte("Claim"))
	stub.PutState("STATUS_DESTORY", []byte("Destroyed or Donated"))	
	stub.PutState("STATUS_RET_WH", []byte("Returned to Warehouse"))
	stub.PutState("STATUS_RCVD_WH", []byte("Received at Warehouse"))	
	
	
	
	return nil, nil
}



//registerUser to register a user
func (t *ABC) registerUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 12 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 12. Got: %d.", len(args))
		}
		
		ffId:=args[0]
		title:=args[1]
		gender:=args[2]
		firstName:=args[3]
		lastName:=args[4]
		dob:=args[5]
		email:=args[6]
		country:=args[7]
		address:=args[8]
		city:=args[9]
		zip:=args[10]
		
		assignerOrg1, err := stub.GetState(args[11])
		assignerOrg := string(assignerOrg1)
		
		createdBy:=assignerOrg
		totalPoint:="0"


		// Insert a row
		ok, err := stub.InsertRow("UserDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: ffId}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: email}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: totalPoint}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}



// to get the deatils of a user against ffid (for internal testing, irrespective of org)
func (t *ABC) getUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	ffId := args[0]
	

	// Get the row pertaining to this ffId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: ffId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("UserDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	res2E := UserDetails{}
	
	res2E.FfId = row.Columns[0].GetString_()
	res2E.Title = row.Columns[1].GetString_()
	res2E.Gender = row.Columns[2].GetString_()
	res2E.FirstName = row.Columns[3].GetString_()
	res2E.LastName = row.Columns[4].GetString_()
	res2E.Dob = row.Columns[5].GetString_()
	res2E.Email = row.Columns[6].GetString_()
	res2E.Country = row.Columns[7].GetString_()
	res2E.Address = row.Columns[8].GetString_()
	res2E.City = row.Columns[9].GetString_()
	res2E.Zip = row.Columns[10].GetString_()
	res2E.CreatedBy = row.Columns[11].GetString_()
	res2E.TotalPoint = row.Columns[12].GetString_()
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


// verify the user is present or not (for internal testing, irrespective of org)
func (t *ABC) verifyUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	ffId := args[0]
	dob := args[1]
	

	// Get the row pertaining to this ffId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: ffId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("UserDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}

	userDob := row.Columns[5].GetString_()
	
	res2E := VerifyU{}
	
	if dob == userDob{
		res2E.Result="success"
	}else{
		res2E.Result="failed"
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}




// add or delete points and insert the transaction(irrespective of org)
func (t *ABC) addDeleteMile(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	trxId := args[0]
	timeStamp:=args[1]
	ffId := args[2]
	
	assignerOrg1, err := stub.GetState(args[3])
	assignerOrg := string(assignerOrg1)
	
	source := assignerOrg
	points := args[4]
	trxntype := args[5]
	trxnSubType := args[6]
	remarks := args[7]
	
	newPoints, _ := strconv.ParseInt(points, 10, 0)
	
	//whether ADD_PENDING, DELETE_PENDING 
	if trxnSubType == "ADD_PENDING" || trxnSubType == "DELETE_PENDING"{
		newPoints = 0
	}
	

	// Get the row pertaining to this ffid
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: ffId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("UserDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving user with ffid %s. Error %s", ffId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}

	newRoyaltyPoint := row.Columns[12].GetString_()
	
	if trxntype=="add"{
		earlierMile:=row.Columns[12].GetString_()
		earlierRoyalty, _:=strconv.ParseInt(earlierMile, 10, 0)
		newRoyaltyPoint = strconv.Itoa(int(earlierRoyalty) + int(newPoints))
	}else if trxntype=="delete"{
	
		earlierMile:=row.Columns[12].GetString_()
		earlierRoyalty, _:=strconv.ParseInt(earlierMile, 10, 0)
		newRoyaltiPointtoTest := int(earlierRoyalty) - int(newPoints)
		
		if newRoyaltiPointtoTest < 0 {
			return nil, errors.New("can't deduct as the resulting royalty becoming less than zero.")
		}
		newRoyaltyPoint = strconv.Itoa(int(earlierRoyalty) - int(newPoints))
	}else{
		return nil, fmt.Errorf("Error: Failed retrieving user with ffid %s. Error %s", ffId, err.Error())
	}
	
	
	//End- Check that the currentStatus to newStatus transition is accurate
	// Delete the row pertaining to this ffid
	err = stub.DeleteRow(
		"UserDetails",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	
	//ffId := row.Columns[0].GetString_()
	
	title := row.Columns[1].GetString_()
	gender := row.Columns[2].GetString_()
	firstName := row.Columns[3].GetString_()
	lastName := row.Columns[4].GetString_()
	dob := row.Columns[5].GetString_()
	email := row.Columns[6].GetString_()
	country := row.Columns[7].GetString_()
	address := row.Columns[8].GetString_()
	city := row.Columns[9].GetString_()
	zip := row.Columns[10].GetString_()
	createdBy := row.Columns[11].GetString_()
	totalPoint := newRoyaltyPoint


		// Insert a row
		ok, err := stub.InsertRow("UserDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: ffId}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: email}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: totalPoint}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}

		
		//inserting the transaction
		
		// Insert a row
		ok, err = stub.InsertRow("Transaction", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: trxId}},
				&shim.Column{Value: &shim.Column_String_{String_: timeStamp}},
				&shim.Column{Value: &shim.Column_String_{String_: ffId}},
				&shim.Column{Value: &shim.Column_String_{String_: source}},
				&shim.Column{Value: &shim.Column_String_{String_: points}},
				&shim.Column{Value: &shim.Column_String_{String_: trxntype}},
				&shim.Column{Value: &shim.Column_String_{String_: trxnSubType}},
				&shim.Column{Value: &shim.Column_String_{String_: remarks}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}		
	return nil, nil

}


//get the miles against the ffid (irrespective of org)
func (t *ABC) getMile(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	ffId := args[0]
	

	// Get the row pertaining to this ffId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: ffId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("UserDetails", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the ffId " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the ffId " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	
	res2E := GetMile{}
	
	res2E.TotalPoint = row.Columns[12].GetString_()
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}



//get all transaction against the ffid (depends on org)
func (t *ABC) getTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	ffId := args[0]
	assignerRole := args[1]

	var columns []shim.Column

	rows, err := stub.GetRows("Transaction", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	assignerOrg1, err := stub.GetState(assignerRole)
	assignerOrg := string(assignerOrg1)
	
		
	res2E:= []*Transaction{}	
	
	for row := range rows {		
		newApp:= new(Transaction)
		newApp.TrxId = row.Columns[0].GetString_()
		newApp.TimeStamp = row.Columns[1].GetString_()
		newApp.FfId = row.Columns[2].GetString_()
		newApp.Source = row.Columns[3].GetString_()
		newApp.Points = row.Columns[4].GetString_()
		newApp.Trxntype = row.Columns[5].GetString_()
		newApp.TrxnSubType = row.Columns[6].GetString_()
		newApp.Remarks = row.Columns[7].GetString_()
		
		if newApp.FfId == ffId && newApp.Source == assignerOrg{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}




//get All transaction against ffid (irrespective of org)
func (t *ABC) getAllTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting ffId to query")
	}

	ffId := args[0]
	//assignerRole := args[1]

	var columns []shim.Column

	rows, err := stub.GetRows("Transaction", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	//assignerOrg1, err := stub.GetState(assignerRole)
	//assignerOrg := string(assignerOrg1)
	
		
	res2E:= []*Transaction{}	
	
	for row := range rows {		
		newApp:= new(Transaction)
		newApp.TrxId = row.Columns[0].GetString_()
		newApp.TimeStamp = row.Columns[1].GetString_()
		newApp.FfId = row.Columns[2].GetString_()
		newApp.Source = row.Columns[3].GetString_()
		newApp.Points = row.Columns[4].GetString_()
		newApp.Trxntype = row.Columns[5].GetString_()
		newApp.TrxnSubType = row.Columns[6].GetString_()
		newApp.Remarks = row.Columns[7].GetString_()
		
		if newApp.FfId == ffId{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


//To check safty
func safeValue(input string) string{ 
if(input == "" ){
	return "NA"
	}else{
	return input
	} 
}

//simple check for chaincode deploy
func (t *ABC) probe(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	retValue := " { \"probe\" : \"Success\" } "
	return []byte(retValue), nil
} 


//createNewASN to register a user
func (t *ABC) createNewASN(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) < 8 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting <8. Got: %d.", len(args))
		}
				
		asnNumber:=args[0]
		
		//getting ASN incrementer
		Avalbytes, err := stub.GetState("ASNincrement")
		Aval, _ := strconv.ParseInt(string(Avalbytes), 10, 0)
		newAval:=int(Aval) + 1
		newASNincrement:= strconv.Itoa(newAval)
		stub.PutState("ASNincrement", []byte(newASNincrement))
		
		asnUniqueid:=string(Avalbytes)
		
		createTimestamp:=args[1]
		updateTimestamp:=args[2]
		updatedBy:=args[3]
		status:=args[4]
		payload1:=args[5]
		
		//tempPay:=payload
		
		assignerOrg1, err := stub.GetState(args[6])
		assignerOrg := string(assignerOrg1)
		
		createdBy:=assignerOrg
		payload:=args[7]
		//Privacy
		if createdBy != "WARE_HOUSE" {
			return nil, fmt.Errorf("You are not authorized to createNewASN")
		}
		
		// Inserting ASN details
		ok, err := stub.InsertRow("ASN", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: asnNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: asnUniqueid}},
				&shim.Column{Value: &shim.Column_String_{String_: createTimestamp}},
				&shim.Column{Value: &shim.Column_String_{String_: updateTimestamp}},
				&shim.Column{Value: &shim.Column_String_{String_: updatedBy}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: payload}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
		
		//inserting item details
	
		payload=payload1
		var itemArray []Item
	
		json.Unmarshal([]byte(payload), &itemArray)
		
		
		for row := range itemArray {		
					
			lineItemId := safeValue(itemArray[row].LineItemId)
			itemId := safeValue(itemArray[row].ItemId)
			description := safeValue(itemArray[row].Description)
			qty := safeValue(itemArray[row].Qty)
			unit := safeValue(itemArray[row].Unit)
			status := safeValue(itemArray[row].Status)
			qtyReceivedAtMedturn := safeValue(itemArray[row].QtyReceivedAtMedturn)
			qtyReceivedAtWarehouse := safeValue(itemArray[row].QtyReceivedAtWarehouse)
			qtyReceivedAtDisposal := safeValue(itemArray[row].QtyReceivedAtDisposal)
			qtyReceivedAtManufacturer := safeValue(itemArray[row].QtyReceivedAtManufacturer)
			createTs := safeValue(itemArray[row].CreateTs)
			updateTs := safeValue(itemArray[row].UpdateTs)
			updatedBy := safeValue(itemArray[row].UpdatedBy)
			remarks := safeValue(itemArray[row].Remarks)
			boxBarcodeNumber := safeValue(itemArray[row].BoxBarcodeNumber)
			debitMemo := safeValue(itemArray[row].DebitMemo)
			lotNumber := safeValue(itemArray[row].LotNumber)
			dc := safeValue(itemArray[row].Dc)
			ndc := safeValue(itemArray[row].Ndc)
			expDate := safeValue(itemArray[row].ExpDate)
			purchageOrderNumber := safeValue(itemArray[row].PurchageOrderNumber)
			asnUniqueid = asnUniqueid
			ddrUniqueid := "NA"
			grmUniqueid := "NA"
			shUniqueid := "NA"
			asnNumber =	asnNumber
			mrrRequestNumber := "NA"
					
					
			// Insert a row 
			ok, err := stub.InsertRow("ITEM", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: lineItemId}},
				&shim.Column{Value: &shim.Column_String_{String_: itemId}},
				&shim.Column{Value: &shim.Column_String_{String_: description}},
				&shim.Column{Value: &shim.Column_String_{String_: qty}},
				&shim.Column{Value: &shim.Column_String_{String_: unit}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtMedturn}},
				&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtWarehouse}},
				&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtDisposal}},
				&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtManufacturer}},
				&shim.Column{Value: &shim.Column_String_{String_: createTs}},
				&shim.Column{Value: &shim.Column_String_{String_: updateTs}},
				&shim.Column{Value: &shim.Column_String_{String_: updatedBy}},
				&shim.Column{Value: &shim.Column_String_{String_: remarks}},
				&shim.Column{Value: &shim.Column_String_{String_: boxBarcodeNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: debitMemo}},
				&shim.Column{Value: &shim.Column_String_{String_: lotNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: dc}},
				&shim.Column{Value: &shim.Column_String_{String_: ndc}},
				&shim.Column{Value: &shim.Column_String_{String_: expDate}},
				&shim.Column{Value: &shim.Column_String_{String_: purchageOrderNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: asnUniqueid}},
				&shim.Column{Value: &shim.Column_String_{String_: ddrUniqueid}},
				&shim.Column{Value: &shim.Column_String_{String_: grmUniqueid}},
				&shim.Column{Value: &shim.Column_String_{String_: shUniqueid}},
				&shim.Column{Value: &shim.Column_String_{String_: asnNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: mrrRequestNumber}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
					
	}
				
		return nil, nil

}



// update LineItem status and record the transaction
func (t *ABC) updateLineItemTempHumid(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) < 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting <4.")
	}

	
	trxId := args[0]
	timeStamp := args[1]
	lineItemId := args[2]
	source := args[3]
	status := args[4]
	trxntype := args[5]
	trxnSubType := args[6]
	remarks := args[7]
	//qty := args[8] //need to remove
	//payload:= args[8]
	trxId=trxId+lineItemId
	
	assignerOrg1, err := stub.GetState(source)
	assignerOrg := string(assignerOrg1)
	
	source = assignerOrg

	//Privacy
	/*if source != "MEDTURN" {
			return nil, fmt.Errorf("You are not authorized to updateLineItem")
	}*/
	
	// Insert a row 
			ok, err := stub.InsertRow("ITEMTempHumidTrxnHistory", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: trxId}},
				&shim.Column{Value: &shim.Column_String_{String_: timeStamp}},
				&shim.Column{Value: &shim.Column_String_{String_: lineItemId}},
				&shim.Column{Value: &shim.Column_String_{String_: source}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: trxntype}},
				&shim.Column{Value: &shim.Column_String_{String_: trxnSubType}},
				&shim.Column{Value: &shim.Column_String_{String_: remarks}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
	
	return nil, nil

}
	


	
//get getLineItemWithHistoryTempHumid(irrespective of organization)
func (t *ABC) getLineItemWithHistoryTempHumid(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 argument to query")
	}

	lineItemId := args[0]
	createdBy := args[1]
	
	fmt.Println(createdBy)
	// Get the row pertaining to this ASN
	var columns []shim.Column
	

	rows, err := stub.GetRows("ITEMTempHumidTrxnHistory", columns)
	
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the lineItemId " + lineItemId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	var itemTrxnHistoryDetail ItemTrxnHistoryDetail
	var itemTrxnHistory ItemTrxnHistory
	itemTrxnHistoryDetail.Trxitemhistory = make([]ItemTrxnHistory, 0)
	
	for row := range rows {		
		fetchedLineItemId := row.Columns[2].GetString_()
		
		if fetchedLineItemId == lineItemId{
			itemTrxnHistory.TrxId = row.Columns[0].GetString_()
			itemTrxnHistory.TimeStamp = row.Columns[1].GetString_()
			itemTrxnHistory.LineItemId = row.Columns[2].GetString_()
			itemTrxnHistory.Source = row.Columns[3].GetString_()
			itemTrxnHistory.Status = row.Columns[4].GetString_()
			itemTrxnHistory.Trxntype = row.Columns[5].GetString_()
			itemTrxnHistory.TrxnSubType = row.Columns[6].GetString_()
			itemTrxnHistory.Remarks = row.Columns[7].GetString_()
			
			itemTrxnHistoryDetail.Trxitemhistory = append(itemTrxnHistoryDetail.Trxitemhistory, itemTrxnHistory)
		}
	
				
	}
		
    mapB, _ := json.Marshal(itemTrxnHistoryDetail)
    fmt.Println(string(mapB))
	
	return mapB, nil
}
	
	



// update LineItem status and record the transaction
func (t *ABC) updateLineItem(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) < 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting <4.")
	}

	
	trxId := args[0]
	timeStamp := args[1]
	lineItemId := args[2]
	source := args[3]
	status := args[4]
	trxntype := args[5]
	trxnSubType := args[6]
	remarks := args[7]
	//qty := args[8] //need to remove
	//payload:= args[8]
	trxId=trxId+lineItemId
	
	assignerOrg1, err := stub.GetState(source)
	assignerOrg := string(assignerOrg1)
	
	source = assignerOrg

	//Privacy
	/*if source != "MEDTURN" {
			return nil, fmt.Errorf("You are not authorized to updateLineItem")
	}*/
	
	// Insert a row 
			ok, err := stub.InsertRow("ITEMTrxnHistory", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: trxId}},
				&shim.Column{Value: &shim.Column_String_{String_: timeStamp}},
				&shim.Column{Value: &shim.Column_String_{String_: lineItemId}},
				&shim.Column{Value: &shim.Column_String_{String_: source}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: trxntype}},
				&shim.Column{Value: &shim.Column_String_{String_: trxnSubType}},
				&shim.Column{Value: &shim.Column_String_{String_: remarks}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
		
		
		// Get the row pertaining to this lineItemId
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: lineItemId}}
		columns = append(columns, col1)

		row, err := stub.GetRow("ITEM", columns)
		if err != nil {
			return nil, fmt.Errorf("Error: Failed retrieving application with lineItemId %s. Error %s", lineItemId, err.Error())
		}

		// GetRows returns empty message if key does not exist
		if len(row.Columns) == 0 {
			return nil, nil
		}
		
		//End- Check that the currentStatus to newStatus transition is accurate
		// Delete the row pertaining to this applicationId
		err = stub.DeleteRow(
			"ITEM",
			columns,
		)
		if err != nil {
			return nil, errors.New("Failed deleting row.")
		}
		
		
			//lineItemId := row.Columns[0].GetString_()
			itemId := row.Columns[1].GetString_()
			description := row.Columns[2].GetString_()
			qty := row.Columns[3].GetString_()
			unit := row.Columns[4].GetString_()
			//status := row.Columns[5].GetString_()
			qtyReceivedAtMedturn := row.Columns[6].GetString_()
			qtyReceivedAtWarehouse := row.Columns[7].GetString_()
			qtyReceivedAtDisposal := row.Columns[8].GetString_()
			qtyReceivedAtManufacturer := row.Columns[9].GetString_()
			createTs := row.Columns[10].GetString_()
			updateTs := row.Columns[11].GetString_()
			updatedBy := row.Columns[12].GetString_()
			//remarks := row.Columns[13].GetString_()
			boxBarcodeNumber := row.Columns[14].GetString_()
			debitMemo := row.Columns[15].GetString_()
			lotNumber := row.Columns[16].GetString_()
			dc := row.Columns[17].GetString_()
			ndc := row.Columns[18].GetString_()
			expDate := row.Columns[19].GetString_()
			purchageOrderNumber := row.Columns[20].GetString_()
			asnUniqueid := row.Columns[21].GetString_()
			ddrUniqueid := row.Columns[22].GetString_()
			grmUniqueid := row.Columns[23].GetString_()
			shUniqueid := row.Columns[24].GetString_()
			asnNumber := row.Columns[25].GetString_()
			mrrRequestNumber := row.Columns[26].GetString_()
					
					
					
			// Insert a row 
			ok, err = stub.InsertRow("ITEM", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: lineItemId}},
				&shim.Column{Value: &shim.Column_String_{String_: itemId}},
				&shim.Column{Value: &shim.Column_String_{String_: description}},
				&shim.Column{Value: &shim.Column_String_{String_: qty}},
				&shim.Column{Value: &shim.Column_String_{String_: unit}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtMedturn}},
				&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtWarehouse}},
				&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtDisposal}},
				&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtManufacturer}},
				&shim.Column{Value: &shim.Column_String_{String_: createTs}},
				&shim.Column{Value: &shim.Column_String_{String_: updateTs}},
				&shim.Column{Value: &shim.Column_String_{String_: updatedBy}},
				&shim.Column{Value: &shim.Column_String_{String_: remarks}},
				&shim.Column{Value: &shim.Column_String_{String_: boxBarcodeNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: debitMemo}},
				&shim.Column{Value: &shim.Column_String_{String_: lotNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: dc}},
				&shim.Column{Value: &shim.Column_String_{String_: ndc}},
				&shim.Column{Value: &shim.Column_String_{String_: expDate}},
				&shim.Column{Value: &shim.Column_String_{String_: purchageOrderNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: asnUniqueid}},
				&shim.Column{Value: &shim.Column_String_{String_: ddrUniqueid}},
				&shim.Column{Value: &shim.Column_String_{String_: grmUniqueid}},
				&shim.Column{Value: &shim.Column_String_{String_: shUniqueid}},
				&shim.Column{Value: &shim.Column_String_{String_: asnNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: mrrRequestNumber}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
		
		
	return nil, nil

}
	

	
	// update LineItem status and record the transaction
func (t *ABC) updateASN(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) < 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting <4.")
	}

	
	trxId := args[0]
	timeStamp := args[1]
	//lineItemId := args[2] 
	updatedBy := args[2]
	status := args[3]
	trxntype := args[4]
	trxnSubType := args[5]
	remarks := args[6]
	asnNumber := args[7]
	payload := args[8]
	source := args[9]
	tempPayload:=payload
	
	
	// Get the row pertaining to this asnNumber
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: asnNumber}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ASN", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving application with asnNumber %s. Error %s", asnNumber, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}
	
	asnUniqueid := row.Columns[1].GetString_()

	
	assignerOrg1, err := stub.GetState(source)
	assignerOrg := string(assignerOrg1)
	
	source = assignerOrg

	//Privacy
	/*if source != "MEDTURN" {
			return nil, fmt.Errorf("You are not authorized to updateLineItem")
	}*/
	
	
	// Insert a row 
			ok, err := stub.InsertRow("ASNTrxnHistory", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: trxId}},
				&shim.Column{Value: &shim.Column_String_{String_: timeStamp}},
				&shim.Column{Value: &shim.Column_String_{String_: asnUniqueid}},
				&shim.Column{Value: &shim.Column_String_{String_: source}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: trxntype}},
				&shim.Column{Value: &shim.Column_String_{String_: trxnSubType}},
				&shim.Column{Value: &shim.Column_String_{String_: remarks}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
	
	
		// Get the row pertaining to this asnNumber
		var columns1 []shim.Column
		col1 = shim.Column{Value: &shim.Column_String_{String_: asnNumber}}
		columns1 = append(columns1, col1)

		row, err = stub.GetRow("ASN", columns1)
		if err != nil {
			return nil, fmt.Errorf("Error: Failed retrieving application with asnNumber %s. Error %s", asnNumber, err.Error())
		}

		// GetRows returns empty message if key does not exist
		if len(row.Columns) == 0 {
			return nil, nil
		}
		
		//End- Check that the currentStatus to newStatus transition is accurate
		// Delete the row pertaining to this applicationId
		err = stub.DeleteRow(
			"ASN",
			columns,
		)
		if err != nil {
			return nil, errors.New("Failed deleting row.")
		}
		
	
	
		//asnNumber:=row.Columns[0].GetString_()
			
		//asnUniqueid:=row.Columns[1].GetString_()
		
		createTimestamp:=row.Columns[2].GetString_()
		updateTimestamp:=row.Columns[3].GetString_()
		//updatedBy:=source
		//status:=row.Columns[5].GetString_()
		createdBy:=row.Columns[6].GetString_()
		payload=row.Columns[7].GetString_()
		
		// Inserting ASN details
		ok, err = stub.InsertRow("ASN", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: asnNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: asnUniqueid}},
				&shim.Column{Value: &shim.Column_String_{String_: createTimestamp}},
				&shim.Column{Value: &shim.Column_String_{String_: updateTimestamp}},
				&shim.Column{Value: &shim.Column_String_{String_: updatedBy}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: payload}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
	
		payload=tempPayload
		var items []string
		json.Unmarshal([]byte(payload), &items)	
		
	
	
		for row := range items {
	
			lineItemId:=items[row]
			trxId=trxId+lineItemId
	
			// Insert a row 
					ok, err = stub.InsertRow("ITEMTrxnHistory", shim.Row{
					Columns: []*shim.Column{
						&shim.Column{Value: &shim.Column_String_{String_: trxId}},
						&shim.Column{Value: &shim.Column_String_{String_: timeStamp}},
						&shim.Column{Value: &shim.Column_String_{String_: lineItemId}},
						&shim.Column{Value: &shim.Column_String_{String_: source}},
						&shim.Column{Value: &shim.Column_String_{String_: status}},
						&shim.Column{Value: &shim.Column_String_{String_: trxntype}},
						&shim.Column{Value: &shim.Column_String_{String_: trxnSubType}},
						&shim.Column{Value: &shim.Column_String_{String_: remarks}},
					}})

				if err != nil {
					return nil, err 
				}
				if !ok && err == nil {
					return nil, errors.New("Row already exists.")
				}
		
		
				// Get the row pertaining to this lineItemId
				var columns []shim.Column
				col1 := shim.Column{Value: &shim.Column_String_{String_: lineItemId}}
				columns = append(columns, col1)

				row, err := stub.GetRow("ITEM", columns)
				if err != nil {
					return nil, fmt.Errorf("Error: Failed retrieving application with lineItemId %s. Error %s", lineItemId, err.Error())
				}

				// GetRows returns empty message if key does not exist
				if len(row.Columns) == 0 {
					return nil, nil
				}
		
				//End- Check that the currentStatus to newStatus transition is accurate
				// Delete the row pertaining to this applicationId
				err = stub.DeleteRow(
					"ITEM",
					columns,
				)
				if err != nil {
					return nil, errors.New("Failed deleting row.")
				}
		
		
				//lineItemId := row.Columns[0].GetString_()
				itemId := row.Columns[1].GetString_()
				description := row.Columns[2].GetString_()
				qty := row.Columns[3].GetString_()
				unit := row.Columns[4].GetString_()
				//status := row.Columns[5].GetString_()
				qtyReceivedAtMedturn := row.Columns[6].GetString_()
				qtyReceivedAtWarehouse := row.Columns[7].GetString_()
				qtyReceivedAtDisposal := row.Columns[8].GetString_()
				qtyReceivedAtManufacturer := row.Columns[9].GetString_()
				createTs := row.Columns[10].GetString_()
				updateTs := row.Columns[11].GetString_()
				updatedBy := source
				//remarks := row.Columns[13].GetString_()
				boxBarcodeNumber := row.Columns[14].GetString_()
				debitMemo := row.Columns[15].GetString_()
				lotNumber := row.Columns[16].GetString_()
				dc := row.Columns[17].GetString_()
				ndc := row.Columns[18].GetString_()
				expDate := row.Columns[19].GetString_()
				purchageOrderNumber := row.Columns[20].GetString_()
				asnUniqueid = row.Columns[21].GetString_()
				ddrUniqueid := row.Columns[22].GetString_()
				grmUniqueid := row.Columns[23].GetString_()
				shUniqueid := row.Columns[24].GetString_()
				asnNumber = row.Columns[25].GetString_()
				mrrRequestNumber := row.Columns[26].GetString_()	
					
					
				// Insert a row 
				ok, err = stub.InsertRow("ITEM", shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: lineItemId}},
					&shim.Column{Value: &shim.Column_String_{String_: itemId}},
					&shim.Column{Value: &shim.Column_String_{String_: description}},
					&shim.Column{Value: &shim.Column_String_{String_: qty}},
					&shim.Column{Value: &shim.Column_String_{String_: unit}},
					&shim.Column{Value: &shim.Column_String_{String_: status}},
					&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtMedturn}},
					&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtWarehouse}},
					&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtDisposal}},
					&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtManufacturer}},
					&shim.Column{Value: &shim.Column_String_{String_: createTs}},
					&shim.Column{Value: &shim.Column_String_{String_: updateTs}},
					&shim.Column{Value: &shim.Column_String_{String_: updatedBy}},
					&shim.Column{Value: &shim.Column_String_{String_: remarks}},
					&shim.Column{Value: &shim.Column_String_{String_: boxBarcodeNumber}},
					&shim.Column{Value: &shim.Column_String_{String_: debitMemo}},
					&shim.Column{Value: &shim.Column_String_{String_: lotNumber}},
					&shim.Column{Value: &shim.Column_String_{String_: dc}},
					&shim.Column{Value: &shim.Column_String_{String_: ndc}},
					&shim.Column{Value: &shim.Column_String_{String_: expDate}},
					&shim.Column{Value: &shim.Column_String_{String_: purchageOrderNumber}},
					&shim.Column{Value: &shim.Column_String_{String_: asnUniqueid}},
					&shim.Column{Value: &shim.Column_String_{String_: ddrUniqueid}},
					&shim.Column{Value: &shim.Column_String_{String_: grmUniqueid}},
					&shim.Column{Value: &shim.Column_String_{String_: shUniqueid}},
					&shim.Column{Value: &shim.Column_String_{String_: asnNumber}},
					&shim.Column{Value: &shim.Column_String_{String_: mrrRequestNumber}},
				}})

			if err != nil {
				return nil, err 
			}
			if !ok && err == nil {
				return nil, errors.New("Row already exists.")
			}
			
		}
	return nil, nil

}


//get getASNDetails(irrespective of organization)
func (t *ABC) getASNDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 argument to query")
	}

	asnNumber := args[0]
	createdBy := args[1]
	
	fmt.Println(createdBy)
	
	// Get the row pertaining to this ASN
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: asnNumber}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ASN", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the asnNumber " + asnNumber + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the asnNumber " + asnNumber + "\"}"
		return nil, errors.New(jsonResp)
	}

	//preparing ASN
	var asnitem AsnItem
	var itemdetails Item
	asnitem.ItemDetail = make([]Item, 0)
	
	asnitem.AsnDetail.AsnNumber = row.Columns[0].GetString_()
	asnitem.AsnDetail.AsnUniqueid = row.Columns[1].GetString_()
	asnitem.AsnDetail.CreateTimestamp = row.Columns[2].GetString_()
	asnitem.AsnDetail.UpdateTimestamp = row.Columns[3].GetString_()
	asnitem.AsnDetail.UpdatedBy = row.Columns[4].GetString_()
	asnitem.AsnDetail.Status = row.Columns[5].GetString_()
	asnitem.AsnDetail.CreatedBy = row.Columns[6].GetString_()
	
	payload:=row.Columns[7].GetString_()
	
	var items []string
	json.Unmarshal([]byte(payload), &items)	
	
	
	for row1 := range items {	

		lineItemId:=items[row1]
	
		// Get the row pertaining to this ASN
		var columns1 []shim.Column
		col2 := shim.Column{Value: &shim.Column_String_{String_: lineItemId}}
		columns1 = append(columns1, col2)

		row, err = stub.GetRow("ITEM", columns1)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get the data for the lineItemId " + lineItemId + "\"}"
			return nil, errors.New(jsonResp)
		}

		// GetRows returns empty message if key does not exist
		if len(row.Columns) == 0 {
			jsonResp := "{\"Error\":\"Failed to get the data for the lineItemId " + lineItemId + "\"}"
			return nil, errors.New(jsonResp)
		}
	
		itemdetails.LineItemId = row.Columns[0].GetString_()
		itemdetails.ItemId = row.Columns[1].GetString_()
		itemdetails.Description = row.Columns[2].GetString_()
		itemdetails.Qty = row.Columns[3].GetString_()
		itemdetails.Unit = row.Columns[4].GetString_()
		itemdetails.Status = row.Columns[5].GetString_()
		itemdetails.QtyReceivedAtMedturn = row.Columns[6].GetString_()
		itemdetails.QtyReceivedAtWarehouse = row.Columns[7].GetString_()
		itemdetails.QtyReceivedAtDisposal = row.Columns[8].GetString_()
		itemdetails.QtyReceivedAtManufacturer = row.Columns[9].GetString_()
		itemdetails.CreateTs = row.Columns[10].GetString_()
		itemdetails.UpdateTs = row.Columns[11].GetString_()
		itemdetails.UpdatedBy = row.Columns[12].GetString_()
		itemdetails.Remarks = row.Columns[13].GetString_()
		itemdetails.BoxBarcodeNumber = row.Columns[14].GetString_()
		itemdetails.DebitMemo = row.Columns[15].GetString_()
		itemdetails.LotNumber = row.Columns[16].GetString_()
		itemdetails.Dc = row.Columns[17].GetString_()
		itemdetails.Ndc = row.Columns[18].GetString_()
		itemdetails.ExpDate = row.Columns[19].GetString_()
		itemdetails.PurchageOrderNumber = row.Columns[20].GetString_()
		itemdetails.AsnNumber = row.Columns[25].GetString_()
		itemdetails.MrrRequestNumber = row.Columns[26].GetString_()
		
		asnitem.ItemDetail = append(asnitem.ItemDetail, itemdetails)		
	}
		
    mapB, _ := json.Marshal(asnitem)
    fmt.Println(string(mapB))
	
	return mapB, nil
}


//get getLineItemWithHistory(irrespective of organization)
func (t *ABC) getLineItemWithHistory(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 argument to query")
	}

	lineItemId := args[0]
	createdBy := args[1]
	
	fmt.Println(createdBy)
	// Get the row pertaining to this ASN
	var columns []shim.Column
	

	rows, err := stub.GetRows("ITEMTrxnHistory", columns)
	
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the lineItemId " + lineItemId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	var itemTrxnHistoryDetail ItemTrxnHistoryDetail
	var itemTrxnHistory ItemTrxnHistory
	itemTrxnHistoryDetail.Trxitemhistory = make([]ItemTrxnHistory, 0)
	
	for row := range rows {		
		fetchedLineItemId := row.Columns[2].GetString_()
		
		if fetchedLineItemId == lineItemId{
			itemTrxnHistory.TrxId = row.Columns[0].GetString_()
			itemTrxnHistory.TimeStamp = row.Columns[1].GetString_()
			itemTrxnHistory.LineItemId = row.Columns[2].GetString_()
			itemTrxnHistory.Source = row.Columns[3].GetString_()
			itemTrxnHistory.Status = row.Columns[4].GetString_()
			itemTrxnHistory.Trxntype = row.Columns[5].GetString_()
			itemTrxnHistory.TrxnSubType = row.Columns[6].GetString_()
			itemTrxnHistory.Remarks = row.Columns[7].GetString_()
			
			itemTrxnHistoryDetail.Trxitemhistory = append(itemTrxnHistoryDetail.Trxitemhistory, itemTrxnHistory)
		}
	
				
	}
		
    mapB, _ := json.Marshal(itemTrxnHistoryDetail)
    fmt.Println(string(mapB))
	
	return mapB, nil
}

//get getLineitem(irrespective of organization)
func (t *ABC) getLineitem(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 argument to query")
	}

	lineItemId := args[0]
	createdBy := args[1]
	
	fmt.Println(createdBy)
	
	
	// Get the row pertaining to this ASN
		var columns1 []shim.Column
		col2 := shim.Column{Value: &shim.Column_String_{String_: lineItemId}}
		columns1 = append(columns1, col2)

		row, err := stub.GetRow("ITEM", columns1)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get the data for the lineItemId " + lineItemId + "\"}"
			return nil, errors.New(jsonResp)
		}

		// GetRows returns empty message if key does not exist
		if len(row.Columns) == 0 {
			jsonResp := "{\"Error\":\"Failed to get the data for the lineItemId " + lineItemId + "\"}"
			return nil, errors.New(jsonResp)
		}
		var itemdetails Item
		
		itemdetails.LineItemId = row.Columns[0].GetString_()
		itemdetails.ItemId = row.Columns[1].GetString_()
		itemdetails.Description = row.Columns[2].GetString_()
		itemdetails.Qty = row.Columns[3].GetString_()
		itemdetails.Unit = row.Columns[4].GetString_()
		itemdetails.Status = row.Columns[5].GetString_()
		itemdetails.QtyReceivedAtMedturn = row.Columns[6].GetString_()
		itemdetails.QtyReceivedAtWarehouse = row.Columns[7].GetString_()
		itemdetails.QtyReceivedAtDisposal = row.Columns[8].GetString_()
		itemdetails.QtyReceivedAtManufacturer = row.Columns[9].GetString_()
		itemdetails.CreateTs = row.Columns[10].GetString_()
		itemdetails.UpdateTs = row.Columns[11].GetString_()
		itemdetails.UpdatedBy = row.Columns[12].GetString_()
		itemdetails.Remarks = row.Columns[13].GetString_()
		itemdetails.BoxBarcodeNumber = row.Columns[14].GetString_()
		itemdetails.DebitMemo = row.Columns[15].GetString_()
		itemdetails.LotNumber = row.Columns[16].GetString_()
		itemdetails.Dc = row.Columns[17].GetString_()
		itemdetails.Ndc = row.Columns[18].GetString_()
		itemdetails.ExpDate = row.Columns[19].GetString_()
		itemdetails.PurchageOrderNumber = row.Columns[20].GetString_()
		itemdetails.AsnNumber = row.Columns[25].GetString_()
		itemdetails.MrrRequestNumber = row.Columns[26].GetString_()
	
	
    mapB, _ := json.Marshal(itemdetails)
    fmt.Println(string(mapB))
	
	return mapB, nil
}

//get getLineitemByStatus(irrespective of organization)
func (t *ABC) getLineitemByStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 argument to query")
	}

	status := args[0]
	createdBy := args[1]
	
	fmt.Println(createdBy)
	
	
	
	
	// Get the row pertaining to this ASN
	var columns []shim.Column
	

	rows, err := stub.GetRows("ITEM", columns)
	
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the status " + status + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	var itemArray ItemArray
	var itemdetails Item
	itemArray.ItemDetail = make([]Item, 0)
	
	for row := range rows {		
		fetchedLineItemStatus := row.Columns[5].GetString_()
		
		if fetchedLineItemStatus == status{
			
			
			itemdetails.LineItemId = row.Columns[0].GetString_()
			itemdetails.ItemId = row.Columns[1].GetString_()
			itemdetails.Description = row.Columns[2].GetString_()
			itemdetails.Qty = row.Columns[3].GetString_()
			itemdetails.Unit = row.Columns[4].GetString_()
			itemdetails.Status = row.Columns[5].GetString_()
			itemdetails.QtyReceivedAtMedturn = row.Columns[6].GetString_()
			itemdetails.QtyReceivedAtWarehouse = row.Columns[7].GetString_()
			itemdetails.QtyReceivedAtDisposal = row.Columns[8].GetString_()
			itemdetails.QtyReceivedAtManufacturer = row.Columns[9].GetString_()
			itemdetails.CreateTs = row.Columns[10].GetString_()
			itemdetails.UpdateTs = row.Columns[11].GetString_()
			itemdetails.UpdatedBy = row.Columns[12].GetString_()
			itemdetails.Remarks = row.Columns[13].GetString_()
			itemdetails.BoxBarcodeNumber = row.Columns[14].GetString_()
			itemdetails.DebitMemo = row.Columns[15].GetString_()
			itemdetails.LotNumber = row.Columns[16].GetString_()
			itemdetails.Dc = row.Columns[17].GetString_()
			itemdetails.Ndc = row.Columns[18].GetString_()
			itemdetails.ExpDate = row.Columns[19].GetString_()
			itemdetails.PurchageOrderNumber = row.Columns[20].GetString_()
			itemdetails.AsnNumber = row.Columns[25].GetString_()
			itemdetails.MrrRequestNumber = row.Columns[26].GetString_()
			
			itemArray.ItemDetail = append(itemArray.ItemDetail, itemdetails)
		}
	}
		
	
    mapB, _ := json.Marshal(itemArray)
    fmt.Println(string(mapB))
	
	return mapB, nil
}


//get getLineitemCountByStatus(irrespective of organization)
func (t *ABC) getLineitemCountByStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 argument to query")
	}

	createdBy := args[0]
	
	fmt.Println(createdBy)	
	
	
	// Get the row pertaining to this ASN
	var columns []shim.Column
	

	rows, err := stub.GetRows("ITEM", columns)
	
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the createdBy " + createdBy + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	var itemstatus ItemStatus

	status1, _ := stub.GetState("STATUS_SHT_MED")
	status2, _ := stub.GetState("STATUS_RCVD_MED")
	status3, _ := stub.GetState("STATUS_INP_MED")
	status4, _ := stub.GetState("STATUS_SHT_MF")
	status5, _ := stub.GetState("STATUS_RCVD_MF")
	status6, _ := stub.GetState("STATUS_SHT_DD")
	status7, _ := stub.GetState("STATUS_RCVD_DD")
	status8, _ := stub.GetState("STATUS_DESTORY")
	status9, _ := stub.GetState("STATUS_RET_WH")
	status10, _ := stub.GetState("STATUS_RCVD_WH")

	stat:= [10]int{0,0,0,0,0,0,0,0,0,0}
	
	
	for row := range rows {		
		fetchedLineItemStatus := row.Columns[5].GetString_()
		
		if fetchedLineItemStatus == string(status1){
				stat[0]=stat[0]+1

				}else if fetchedLineItemStatus == string(status2){
					stat[1]=stat[1]+1

				}else if fetchedLineItemStatus == string(status3){
					stat[2]=stat[2]+1

				}else if fetchedLineItemStatus == string(status4){
					stat[3]=stat[3]+1

				}else if fetchedLineItemStatus == string(status5){
					stat[4]=stat[4]+1

				}else if fetchedLineItemStatus == string(status6){
					stat[5]=stat[5]+1

				}else if fetchedLineItemStatus == string(status7){
					stat[6]=stat[6]+1

				}else if fetchedLineItemStatus == string(status8){
					stat[7]=stat[7]+1

				}else if fetchedLineItemStatus == string(status9){
					stat[8]=stat[8]+1

				}else if fetchedLineItemStatus == string(status10){
					stat[9]=stat[9]+1

				}
	}

	itemstatus.Status_sht_med=strconv.Itoa(stat[0])
	itemstatus.Status_rcvd_med=strconv.Itoa(stat[1])
	itemstatus.Status_inp_med=strconv.Itoa(stat[2])
	itemstatus.Status_sht_mf=strconv.Itoa(stat[3])
	itemstatus.Status_rcvd_mf=strconv.Itoa(stat[4])
	itemstatus.Status_sht_dd=strconv.Itoa(stat[5])
	itemstatus.Status_recv_dd=strconv.Itoa(stat[6])
	itemstatus.Status_destroy=strconv.Itoa(stat[7])
	itemstatus.Status_ret_wh=strconv.Itoa(stat[8])
	itemstatus.Status_rcvd_wh=strconv.Itoa(stat[9])
	
    mapB, _ := json.Marshal(itemstatus)
    fmt.Println(string(mapB))
	
	return mapB, nil
}




//createMRR to register a user
func (t *ABC) createMRR(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) < 8 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting <8. Got: %d.", len(args))
		}
				
		requestNumber:=args[0]
		
		//getting MMR incrementer
		Avalbytes, err := stub.GetState("GRMincrement")
		Aval, _ := strconv.ParseInt(string(Avalbytes), 10, 0)
		newAval:=int(Aval) + 1
		newMMRincrement:= strconv.Itoa(newAval)
		stub.PutState("GRMincrement", []byte(newMMRincrement))
		
		mrrUniqueid:=string(Avalbytes)
		
		createTimestamp:=args[1]
		updateTimestamp:=args[2]
		updatedBy:=args[3]
		status:=args[4]
		payload:=args[5]
		trxId:=args[6]
		remarks := args[7]
		
		//tempPay:=payload
		
		assignerOrg1, err := stub.GetState(args[8])
		assignerOrg := string(assignerOrg1)
		
		createdBy:=assignerOrg
		
		
		// Inserting MRR details
		ok, err := stub.InsertRow("MRR", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: requestNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: mrrUniqueid}},
				&shim.Column{Value: &shim.Column_String_{String_: createTimestamp}},
				&shim.Column{Value: &shim.Column_String_{String_: updateTimestamp}},
				&shim.Column{Value: &shim.Column_String_{String_: updatedBy}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: payload}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
		
		//trxId := args[0]
		timeStamp := createTimestamp
		//lineItemId := args[2] 
		//updatedBy := args[2]
		//status := args[3]
		trxntype := "NA"
		trxnSubType := "NA"
		//remarks := args[6]
		//asnNumber := args[7]
		//payload := args[8]
		source := createdBy
		//tempPayload:=payload
		
		var items []string
		json.Unmarshal([]byte(payload), &items)	
		
	
	
		for row := range items {
	
			lineItemId:=items[row]
			trxId=trxId+lineItemId
	
			// Insert a row 
					ok, err = stub.InsertRow("ITEMTrxnHistory", shim.Row{
					Columns: []*shim.Column{
						&shim.Column{Value: &shim.Column_String_{String_: trxId}},
						&shim.Column{Value: &shim.Column_String_{String_: timeStamp}},
						&shim.Column{Value: &shim.Column_String_{String_: lineItemId}},
						&shim.Column{Value: &shim.Column_String_{String_: source}},
						&shim.Column{Value: &shim.Column_String_{String_: status}},
						&shim.Column{Value: &shim.Column_String_{String_: trxntype}},
						&shim.Column{Value: &shim.Column_String_{String_: trxnSubType}},
						&shim.Column{Value: &shim.Column_String_{String_: remarks}},
					}})

				if err != nil {
					return nil, err 
				}
				if !ok && err == nil {
					return nil, errors.New("Row already exists.")
				}
		
		
				// Get the row pertaining to this lineItemId
				var columns []shim.Column
				col1 := shim.Column{Value: &shim.Column_String_{String_: lineItemId}}
				columns = append(columns, col1)

				row, err := stub.GetRow("ITEM", columns)
				if err != nil {
					return nil, fmt.Errorf("Error: Failed retrieving application with lineItemId %s. Error %s", lineItemId, err.Error())
				}

				// GetRows returns empty message if key does not exist
				if len(row.Columns) == 0 {
					return nil, nil
				}
		
				//End- Check that the currentStatus to newStatus transition is accurate
				// Delete the row pertaining to this applicationId
				err = stub.DeleteRow(
					"ITEM",
					columns,
				)
				if err != nil {
					return nil, errors.New("Failed deleting row.")
				}
		
		
				//lineItemId := row.Columns[0].GetString_()
				itemId := row.Columns[1].GetString_()
				description := row.Columns[2].GetString_()
				qty := row.Columns[3].GetString_()
				unit := row.Columns[4].GetString_()
				//status := row.Columns[5].GetString_()
				qtyReceivedAtMedturn := row.Columns[6].GetString_()
				qtyReceivedAtWarehouse := row.Columns[7].GetString_()
				qtyReceivedAtDisposal := row.Columns[8].GetString_()
				qtyReceivedAtManufacturer := row.Columns[9].GetString_()
				createTs := row.Columns[10].GetString_()
				updateTs := row.Columns[11].GetString_()
				updatedBy := source
				//remarks := row.Columns[13].GetString_()
				boxBarcodeNumber := row.Columns[14].GetString_()
				debitMemo := row.Columns[15].GetString_()
				lotNumber := row.Columns[16].GetString_()
				dc := row.Columns[17].GetString_()
				ndc := row.Columns[18].GetString_()
				expDate := row.Columns[19].GetString_()
				purchageOrderNumber := row.Columns[20].GetString_()
				asnUniqueid := row.Columns[21].GetString_()
				ddrUniqueid := row.Columns[22].GetString_()
				grmUniqueid := row.Columns[23].GetString_()
				shUniqueid := row.Columns[24].GetString_()
				asnNumber := row.Columns[25].GetString_()
				mrrRequestNumber := requestNumber	
					
					
				// Insert a row 
				ok, err = stub.InsertRow("ITEM", shim.Row{
				Columns: []*shim.Column{
					&shim.Column{Value: &shim.Column_String_{String_: lineItemId}},
					&shim.Column{Value: &shim.Column_String_{String_: itemId}},
					&shim.Column{Value: &shim.Column_String_{String_: description}},
					&shim.Column{Value: &shim.Column_String_{String_: qty}},
					&shim.Column{Value: &shim.Column_String_{String_: unit}},
					&shim.Column{Value: &shim.Column_String_{String_: status}},
					&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtMedturn}},
					&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtWarehouse}},
					&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtDisposal}},
					&shim.Column{Value: &shim.Column_String_{String_: qtyReceivedAtManufacturer}},
					&shim.Column{Value: &shim.Column_String_{String_: createTs}},
					&shim.Column{Value: &shim.Column_String_{String_: updateTs}},
					&shim.Column{Value: &shim.Column_String_{String_: updatedBy}},
					&shim.Column{Value: &shim.Column_String_{String_: remarks}},
					&shim.Column{Value: &shim.Column_String_{String_: boxBarcodeNumber}},
					&shim.Column{Value: &shim.Column_String_{String_: debitMemo}},
					&shim.Column{Value: &shim.Column_String_{String_: lotNumber}},
					&shim.Column{Value: &shim.Column_String_{String_: dc}},
					&shim.Column{Value: &shim.Column_String_{String_: ndc}},
					&shim.Column{Value: &shim.Column_String_{String_: expDate}},
					&shim.Column{Value: &shim.Column_String_{String_: purchageOrderNumber}},
					&shim.Column{Value: &shim.Column_String_{String_: asnUniqueid}},
					&shim.Column{Value: &shim.Column_String_{String_: ddrUniqueid}},
					&shim.Column{Value: &shim.Column_String_{String_: grmUniqueid}},
					&shim.Column{Value: &shim.Column_String_{String_: shUniqueid}},
					&shim.Column{Value: &shim.Column_String_{String_: asnNumber}},
					&shim.Column{Value: &shim.Column_String_{String_: mrrRequestNumber}},
				}})

			if err != nil {
				return nil, err 
			}
			if !ok && err == nil {
				return nil, errors.New("Row already exists.")
			}
			
		}
		
		
		
		
		return nil, nil

}


//get getMRRDetails(irrespective of organization)
func (t *ABC) getMRRDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2 argument to query")
	}

	requestNumber := args[0]
	createdBy := args[1]
	
	fmt.Println(createdBy)
	
	// Get the row pertaining to this MRR
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: requestNumber}}
	columns = append(columns, col1)

	row, err := stub.GetRow("MRR", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the requestNumber " + requestNumber + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the requestNumber " + requestNumber + "\"}"
		return nil, errors.New(jsonResp)
	}

	//preparing MRR
	var mrritem MrrItem
	var itemdetails Item
	mrritem.ItemDetail = make([]Item, 0)
	
	mrritem.MrrDetail.RequestNumber = row.Columns[0].GetString_()
	mrritem.MrrDetail.MrrUniqueid = row.Columns[1].GetString_()
	mrritem.MrrDetail.CreateTimestamp = row.Columns[2].GetString_()
	mrritem.MrrDetail.UpdateTimestamp = row.Columns[3].GetString_()
	mrritem.MrrDetail.UpdatedBy = row.Columns[4].GetString_()
	mrritem.MrrDetail.Status = row.Columns[5].GetString_()
	mrritem.MrrDetail.CreatedBy = row.Columns[6].GetString_()
	
	payload:=row.Columns[7].GetString_()
	
	var items []string
	json.Unmarshal([]byte(payload), &items)	
	
	
	for row1 := range items {	

		lineItemId:=items[row1]
	
		// Get the row pertaining to this ASN
		var columns1 []shim.Column
		col2 := shim.Column{Value: &shim.Column_String_{String_: lineItemId}}
		columns1 = append(columns1, col2)

		row, err = stub.GetRow("ITEM", columns1)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get the data for the lineItemId " + lineItemId + "\"}"
			return nil, errors.New(jsonResp)
		}

		// GetRows returns empty message if key does not exist
		if len(row.Columns) == 0 {
			jsonResp := "{\"Error\":\"Failed to get the data for the lineItemId " + lineItemId + "\"}"
			return nil, errors.New(jsonResp)
		}
	
		itemdetails.LineItemId = row.Columns[0].GetString_()
		itemdetails.ItemId = row.Columns[1].GetString_()
		itemdetails.Description = row.Columns[2].GetString_()
		itemdetails.Qty = row.Columns[3].GetString_()
		itemdetails.Unit = row.Columns[4].GetString_()
		itemdetails.Status = row.Columns[5].GetString_()
		itemdetails.QtyReceivedAtMedturn = row.Columns[6].GetString_()
		itemdetails.QtyReceivedAtWarehouse = row.Columns[7].GetString_()
		itemdetails.QtyReceivedAtDisposal = row.Columns[8].GetString_()
		itemdetails.QtyReceivedAtManufacturer = row.Columns[9].GetString_()
		itemdetails.CreateTs = row.Columns[10].GetString_()
		itemdetails.UpdateTs = row.Columns[11].GetString_()
		itemdetails.UpdatedBy = row.Columns[12].GetString_()
		itemdetails.Remarks = row.Columns[13].GetString_()
		itemdetails.BoxBarcodeNumber = row.Columns[14].GetString_()
		itemdetails.DebitMemo = row.Columns[15].GetString_()
		itemdetails.LotNumber = row.Columns[16].GetString_()
		itemdetails.Dc = row.Columns[17].GetString_()
		itemdetails.Ndc = row.Columns[18].GetString_()
		itemdetails.ExpDate = row.Columns[19].GetString_()
		itemdetails.PurchageOrderNumber = row.Columns[20].GetString_()
		itemdetails.AsnNumber = row.Columns[25].GetString_()
		itemdetails.MrrRequestNumber = row.Columns[26].GetString_()
		
		mrritem.ItemDetail = append(mrritem.ItemDetail, itemdetails)		
	}
		
    mapB, _ := json.Marshal(mrritem)
    fmt.Println(string(mapB))
	
	return mapB, nil
}





// Invoke invokes the chaincode
func (t *ABC) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "createNewASN" {
		t := ABC{}
		return t.createNewASN(stub, args)	
	}else if function == "updateLineItem" { 
		t := ABC{}
		return t.updateLineItem(stub, args)
	}else if function == "updateASN" { 
		t := ABC{}
		return t.updateASN(stub, args)
	}else if function == "createMRR" { 
		t := ABC{}
		return t.createMRR(stub, args)
	}else if function == "registerUser" { 
		t := ABC{}
		return t.registerUser(stub, args)
	}else if function == "addDeleteMile" { 
		t := ABC{}
		return t.addDeleteMile(stub, args)
	}else if function == "updateLineItemTempHumid" { 
		t := ABC{}
		return t.updateLineItemTempHumid(stub, args)
	}
	
	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *ABC) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getASNDetails" {
		t := ABC{}
		return t.getASNDetails(stub, args)		
	}else if function == "getLineItemWithHistory" { 
		t := ABC{}
		return t.getLineItemWithHistory(stub, args)
	}else if function == "probe" { 
		t := ABC{}
		return t.probe(stub, args)
	}else if function == "getLineitem" { 
		t := ABC{}
		return t.getLineitem(stub, args)
	}else if function == "getLineitemByStatus" { 
		t := ABC{}
		return t.getLineitemByStatus(stub, args)
	}else if function == "getLineitemCountByStatus" { 
		t := ABC{}
		return t.getLineitemCountByStatus(stub, args)
	}else if function == "getMRRDetails" { 
		t := ABC{}
		return t.getMRRDetails(stub, args)
	}else if function == "getMile" {
		t := ABC{}
		return t.getMile(stub, args)		
	} else if function == "getTransaction" { 
		t := ABC{}
		return t.getTransaction(stub, args)
	}else if function == "getAllTransaction" { 
		t := ABC{}
		return t.getAllTransaction(stub, args)
	} else if function == "getUser" { 
		t := ABC{}
		return t.getUser(stub, args)
	}else if function == "verifyUser" { 
		t := ABC{}
		return t.verifyUser(stub, args)
	}else if function == "getLineItemWithHistoryTempHumid" { 
		t := ABC{}
		return t.getLineItemWithHistoryTempHumid(stub, args)
	}
	
	return nil, nil
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(ABC))
	if err != nil {
		fmt.Printf("Error starting ABC: %s", err)
	}
}
