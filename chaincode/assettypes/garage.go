package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

var Garage = assets.AssetType{
	Tag:         "garage",
	Label:       "Garage",
	Description: "Garage for storing vehicles",

	Props: []assets.AssetProp{

		{
			Required: true,
			IsKey:    true,
			Tag:      "name",
			Label:    "Name",
			DataType: "string",
		},
		{
			Tag:      "vehicles",
			Label:    "Vehicles",
			DataType: "[]->vehicle",
		},
	},
}
