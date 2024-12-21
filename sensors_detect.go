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

func Init() error {
	if ret := C.sensors_init(nil); ret != 0 {
		return fmt.Errorf("can not init sensors package")
	} else {
		return nil
	}
}

func CleanUp() {
	C.sensors_cleanup()
}

func DetectChips() []Chip {
	var chips []Chip = make([]Chip, 0)
	for i := 0; ; i++ {
		chip, err := DetectChip(i)
		if err != nil {
			break
		}
		chips = append(chips, chip)
	}
	return chips
}

func DetectChip(chipNumber int) (Chip, error) {
	var chip *C.struct_sensors_chip_name
	cChipNumber := C.int(chipNumber)
	cChipNumberPtr := &cChipNumber
	chip = C.sensors_get_detected_chips(nil, cChipNumberPtr)
	if chip != nil {
		goChip := cChip2Go(*chip)
		return goChip, nil
	} else {
		return Chip{}, fmt.Errorf("chip %d not found", chipNumber)
	}
}

func GetFeatures(chip Chip) []Feature {
	cChip := goChip2C(chip)
	defer C.free(unsafe.Pointer(cChip.path))
	defer C.free(unsafe.Pointer(cChip.prefix))
	var featureNum C.int
	var feature *C.struct_sensors_feature
	features := make([]Feature, 0)
	for {
		feature = C.sensors_get_features(&cChip, &featureNum)
		if feature == nil {
			break
		}
		features = append(features, cFeature2Go(*feature))
	}

	return features
}

func GetSubfeatures(chip Chip, feature Feature) []SubFeature {
	cChip := goChip2C(chip)
	defer C.free(unsafe.Pointer(cChip.path))
	defer C.free(unsafe.Pointer(cChip.prefix))
	cFeature := goFeature2C(feature)
	defer C.free(unsafe.Pointer(cFeature.name))
	var subfeatureNum C.int
	var subfeature *C.struct_sensors_subfeature
	var subfeatureValue C.double
	subfeatures := make([]SubFeature, 0)
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
