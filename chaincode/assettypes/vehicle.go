package assettypes

import (
	"fmt"
	"github.com/hyperledger-labs/cc-tools/assets"
	"math"
	"strings"
)

var Vehicle = assets.AssetType{
	Tag:         "vehicle",
	Label:       "Vehicle",
	Description: "A Car With Your Features and Owner",
	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "ID",
			DataType: "string",
			Validate: func(id interface{}) error {
				if strings.TrimSpace(id.(string)) == "" {
					return fmt.Errorf("the id cannot be empty")
				}
				return nil
			},
			Writers: []string{"org1MSP", "orgMSP"},
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
			Validate: func(model interface{}) error {
				if strings.TrimSpace(model.(string)) == "" {
					return fmt.Errorf("the model cannot be empty")
				}
				return nil
			},
		},
		{
			Required: true,
			Tag:      "mark",
			Label:    "Mark",
			DataType: "string",
			Validate: func(mark interface{}) error {
				if strings.TrimSpace(mark.(string)) == "" {
					return fmt.Errorf("the mark cannot be empty")
				}
				return nil
			},
		},
		{
			Tag:      "year",
			Label:    "Year",
			DataType: "number",
			Validate: func(year interface{}) error {
				if math.Ceil(year.(float64)) != year.(float64) {
					return fmt.Errorf("the year must be an integer")
				}
				return nil
			},
		},
		{
			Tag:      "color",
			Label:    "Color",
			DataType: "string",
			Validate: func(color interface{}) error {
				if strings.TrimSpace(color.(string)) == "" {
					return fmt.Errorf("the color cannot be empty")
				}
				return nil
			},
		},
	},
}
