package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	""
	
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var CreateNewVehicle = tx.Transaction{
	Tag: "createNewVehicle",
	Label: "Create New Vehicle",
	Description: "Create New Vehicle",
	Callers: []accesscontrol.Caller{
		{
			MSP: "orgMSP",
			OU: "admin",
		},
		{
			MSP: "org1MSP",
			OU: "admin",
		},
	},

	Args: []tx.Argument{
		{
			Required: true,
			Tag:      "id",
			Label:    "ID",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "vehicleType",
			Label:    "Vehicle Type",
			DataType: "vehicleType",
		},
		{
			Required: true,
			Tag:      "owner",
			Label:    "Owner",
			DataType: "->person",
		
		},
		{
			Required: true,
			Tag:      "model",
			Label:    "Model",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "mark",
			Label:    "Mark",
			DataType: "string",
		},
		{
			Tag:      "year",
			Label:    "Year",
			DataType: "number",
		},
		{
			Tag:      "color",
			Label:    "Color",
			DataType: "string",
		},
	},
	
	Routine: func(sw *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		id := req["id"].(string)
		vehicleType := req["vehicleType"].(vehicleType)


		
	},
}