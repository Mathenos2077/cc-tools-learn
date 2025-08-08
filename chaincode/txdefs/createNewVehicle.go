package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools-demo/chaincode/datatypes"
	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
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
		vehicleType := req["vehicleType"].(datatypes.VehicleType)
		ownerID := req["owner"].(string)
		model := req["model"].(string)
		mark := req["mark"].(string)
		year := req["year"].(string)
		color := req["color"].(string)

		vehicleMap := make(map[string]interface{})
		vehicleMap["@assetType"] = "vehicle"
		vehicleMap["id"] = id
		vehicleMap["vehicleType"] = vehicleType
		vehicleMap["ownerID"] = ownerID
		vehicleMap["model"] = model
		vehicleMap["mark"] = mark
		vehicleMap["year"] = year
		vehicleMap["color"] = color

		vehicleAsset, err := assets.NewAsset(vehicleMap)
		
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		_, err = vehicleAsset.PutNew(sw)

		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "Error saving asset on blockchain", err.Status())
		}

		vehicleJSON, nerr := json.Marshal(vehicleAsset)

		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return vehicleJSON, nil
	},
}