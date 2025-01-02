package go_sensors

/*
#cgo LDFLAGS: -lsensors
#include <sensors/sensors.h>
#include <stdlib.h>

*/
import (
	"C"
)

import (
	"fmt"
	"unsafe"
)

func detectChips() []goChip {
	var chips []goChip = make([]goChip, 0)
	for i := 0; ; i++ {
		chip, err := detectChip(i)
		if err != nil {
			break
		}
		chips = append(chips, chip)
	}
	return chips
}

func detectChip(chipNumber int) (goChip, error) {
	var chip *C.struct_sensors_chip_name
	cChipNumber := C.int(chipNumber)
	cChipNumberPtr := &cChipNumber
	chip = C.sensors_get_detected_chips(nil, cChipNumberPtr)
	if chip != nil {
		goChip := cChip2Go(*chip)
		return goChip, nil
	} else {
		return goChip{}, fmt.Errorf("chip %d not found", chipNumber)
	}
}

func getFeatures(chip goChip) []goFeature {
	cChip := goChip2C(chip)
	defer C.free(unsafe.Pointer(cChip.path))
	defer C.free(unsafe.Pointer(cChip.prefix))
	var featureNum C.int
	var feature *C.struct_sensors_feature
	features := make([]goFeature, 0)
	for {
		feature = C.sensors_get_features(&cChip, &featureNum)
		if feature == nil {
			break
		}
		features = append(features, cFeature2Go(*feature))
	}

	return features
}

func getSubfeatures(chip goChip, feature goFeature) []goSubFeature {
	cChip := goChip2C(chip)
	defer C.free(unsafe.Pointer(cChip.path))
	defer C.free(unsafe.Pointer(cChip.prefix))
	cFeature := goFeature2C(feature)
	defer C.free(unsafe.Pointer(cFeature.name))
	var subfeatureNum C.int
	var subfeature *C.struct_sensors_subfeature
	var subfeatureValue C.double
	subfeatures := make([]goSubFeature, 0)
	for {
		subfeature = C.sensors_get_all_subfeatures(&cChip, &cFeature, &subfeatureNum)
		if subfeature == nil {
			break
		}
		execute_code := C.sensors_get_value(&cChip, subfeature.number, &subfeatureValue)
		if execute_code != 0 {
			panic("Can't get value")
		}
		goSubFeature := cSubFeature2Go(*subfeature)
		goSubFeature.Value = float64(subfeatureValue)
		subfeatures = append(subfeatures, goSubFeature)
	}
	return subfeatures
}
