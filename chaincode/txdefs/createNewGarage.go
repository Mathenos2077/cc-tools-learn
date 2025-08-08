package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var CreateNewGarage = tx.Transaction{
	Tag:         "createNewGarage",
	Label:       "Create New Garage",
	Description: "Create a New Garage",
	Method:      "POST",
	Callers: []accesscontrol.Caller{
		{
			MSP: "org3MSP",
			OU:  "admin",
		},
		{
			MSP: "orgMSP",
			OU:  "admin",
		},
	},

	Args: []tx.Argument{
		{
			Required:    true,
			Tag:         "name",
			Label:       "Name",
			Description: "The Name Of Garage",
			DataType:    "string",
		},
	},

	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		name := req["name"].(string)

		garageMap := make(map[string]interface{})
		garageMap["@assetType"] = "garage"
		garageMap["name"] = name

		garageAsset, err := assets.NewAsset(garageMap)

		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		_, err = garageAsset.PutNew(stub)

		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "Error saving asset on blockchain", err.Status())
		}

		garageJSON, nerr := json.Marshal(garageAsset)

		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return garageJSON, nil

	},
}
