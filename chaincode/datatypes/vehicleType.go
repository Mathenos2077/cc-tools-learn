package datatypes

import (
	"fmt"
	"strconv"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type VehicleType float64

const (
	VehicleTypeCar VehicleType = iota + 1
	VehicleTypeMotorcycle
	VehicleTypeBike
	VehicleTypeTruck
	VehicleTypePlane
	VehicleTypeHelicopter
	VehicleTypeBoat
	VehicleTypeSubmarine
)

func (v VehicleType) CheckType() errors.ICCError {
	switch v {
	case VehicleTypeCar:
		return nil
	case VehicleTypeMotorcycle:
		return nil
	case VehicleTypeBike:
		return nil
	case VehicleTypeTruck:
		return nil
	case VehicleTypePlane:
		return nil
	case VehicleTypeHelicopter:
		return nil
	case VehicleTypeBoat:
		return nil
	case VehicleTypeSubmarine:
		return nil
	default:
		return errors.NewCCError("invalid type", 400)
	}

}

var vehicleType = assets.DataType{
	AcceptedFormats: []string{"number"},
	Description:     "The category into which a vehicle falls",
	DropDownValues: map[string]interface{}{
		"Car":        VehicleTypeCar,
		"Motorcycle": VehicleTypeMotorcycle,
		"Bike":       VehicleTypeBike,
		"Truck":      VehicleTypeTruck,
		"Plane":      VehicleTypePlane,
		"Helicopter": VehicleTypeHelicopter,
		"Boat":       VehicleTypeBoat,
		"Submarine":  VehicleTypeSubmarine,
	},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		var dataVal float64

		switch v := data.(type) {
		case float64:
			dataVal = v
		case int:
			dataVal = float64(v)
		case VehicleType:
			dataVal = float64(v)
		case string:
			var err error
			dataVal, err = strconv.ParseFloat(v, 64)
			if err != nil {
				return "", nil, errors.WrapErrorWithStatus(err, "asset property must be an integer, is %t", 400)
			}
		default:
			return "", nil, errors.NewCCError("asset property must be an integer, is %t", 400)
		}

		retVal := VehicleType(dataVal)
		err := retVal.CheckType()
		return fmt.Sprint(retVal), retVal, err
	},
}
