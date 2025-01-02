package go_sensors

/*
#cgo LDFLAGS: -lsensors
#include <sensors/sensors.h>
#include <stdlib.h>

*/
import (
	"C"
)

func cChip2Go(chip C.struct_sensors_chip_name) goChip {
	goChip := goChip{
		Prefix: C.GoString(chip.prefix),
		Bus:    goBusId{Type: int(chip.bus._type), Nr: int(chip.bus.nr)},
		Addr:   int(chip.addr),
		Path:   C.GoString(chip.path),
	}
	return goChip
}

func goChip2C(chip goChip) C.struct_sensors_chip_name {
	cChip := C.struct_sensors_chip_name{
		prefix: C.CString(chip.Prefix),
		bus: C.struct_sensors_bus_id{
			_type: C.short(chip.Bus.Type),
			nr:    C.short(chip.Bus.Nr),
		},
		addr: C.int(chip.Addr),
		path: C.CString(chip.Path),
	}
	return cChip
}

func cFeature2Go(feature C.struct_sensors_feature) goFeature {
	goFeature := goFeature{
		Name:            C.GoString(feature.name),
		Number:          int(feature.number),
		Type:            sensorsFeatureType(feature._type),
		FirstSubfeature: int(feature.first_subfeature),
		Padding:         int(feature.padding1),
	}
	return goFeature
}

func goFeature2C(feature goFeature) C.struct_sensors_feature {
	cFeature := C.struct_sensors_feature{
		name:             C.CString(feature.Name),
		number:           C.int(feature.Number),
		_type:            C.sensors_feature_type(feature.Type),
		first_subfeature: C.int(feature.FirstSubfeature),
		padding1:         C.int(feature.Padding),
	}
	return cFeature
}

func cSubFeature2Go(subfeature C.struct_sensors_subfeature) goSubFeature {
	goSubFeature := goSubFeature{
		Name:    C.GoString(subfeature.name),
		Number:  int(subfeature.number),
		Type:    sensorsSubfeatureType(subfeature._type),
		Mapping: int(subfeature.mapping),
		Flags:   uint(subfeature.flags),
	}
	return goSubFeature
}

func goSubFeature2C(subfeature goSubFeature) C.struct_sensors_subfeature {
	cSubFeature := C.struct_sensors_subfeature{
		name:    C.CString(subfeature.Name),
		number:  C.int(subfeature.Number),
		_type:   C.sensors_subfeature_type(subfeature.Type),
		mapping: C.int(subfeature.Mapping),
		flags:   C.uint(subfeature.Flags),
	}
	return cSubFeature
}
