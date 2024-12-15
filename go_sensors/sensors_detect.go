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
	"log"
	"unsafe"
)

func PrintSensorsInfo() []Chip {
	// Инициализация библиотеки sensors
	if ret := C.sensors_init(nil); ret != 0 {
		log.Fatal("Не удалось инициализировать библиотеку sensors.")
		return []Chip{}
	}
	defer C.sensors_cleanup() // Завершение работы с библиотекой после завершения функции

	var chipNum C.int
	var chip *C.struct_sensors_chip_name
	var goChip Chip
	chips := make([]Chip, 0)

	// Итерация по чипам
	for {
		chip = C.sensors_get_detected_chips(nil, &chipNum)
		if chip == nil {
			break
		}

		var featureNum C.int
		var feature *C.struct_sensors_feature
		goChip = Chip{
			Prefix: C.GoString(chip.prefix),
			Bus:    BusId{Type: int(chip.bus._type), Nr: int(chip.bus.nr)},
			Addr:   int(chip.addr),
			Path:   C.GoString(chip.path),
		}
		// Итерация по фичам чипа
		for {
			feature = C.sensors_get_features(chip, &featureNum)
			if feature == nil {
				break
			}

			var subfeatureNum C.int
			var subfeature *C.struct_sensors_subfeature

			// Итерация по сабфитчам фичи
			for {
				subfeature = C.sensors_get_all_subfeatures(chip, feature, &subfeatureNum)
				if subfeature == nil {
					break
				}

				var value C.double
				if C.sensors_get_value(chip, subfeature.number, &value) == 0 {
					var unit string
					// switch subfeature._type {
					// case C.SENSORS_SUBFEATURE_TEMP_INPUT:
					// 	unit = "°C"
					// case C.SENSORS_SUBFEATURE_POWER_INPUT:
					// 	unit = "W"
					// default:
					// 	unit = ""
					// }
					// C.sensors.SENSORS_SUBFEATURE_POWER_INPUT
					fmt.Println(subfeature._type)
					fmt.Printf("    Сабфитч: %s, Значение: %.2f %s\n",
						C.GoString(subfeature.name),
						float64(value),
						unit)
				} else {
					fmt.Printf("  Не удалось получить значение сабфитча: %s\n", C.GoString(subfeature.name))
				}
			}
		}
		chips = append(chips, goChip)
	}
	return chips
}

func Init() error {
	if ret := C.sensors_init(nil); ret != 0 {
		return fmt.Errorf("не удалось инициализировать библиотеку sensors")
	} else {
		return nil
	}
}

func CleanUp() {
	C.sensors_cleanup()
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
	subfeatures := make([]SubFeature, 0)
	for {
		subfeature = C.sensors_get_all_subfeatures(&cChip, &cFeature, &subfeatureNum)
		if subfeature == nil {
			break
		}
		subfeatures = append(subfeatures, cSubFeature2Go(*subfeature))
	}
	return subfeatures
}
