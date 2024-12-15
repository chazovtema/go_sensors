package go_sensors

import "C"

func cChip2Go(chip C.struct_sensors_chip_name) Chip {
	goChip := Chip{
		Prefix: C.GoString(chip.prefix),
		Bus:    BusId{Type: int(chip.bus._type), Nr: int(chip.bus.nr)},
		Addr:   int(chip.addr),
		Path:   C.GoString(chip.path),
	}
	return goChip
}

func goChip2C(chip Chip) C.struct_sensors_chip_name {
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

func cFeature2Go(feature C.struct_sensors_feature) Feature {
	goFeature := Feature{
		Name:   C.GoString(feature.name),
		Number: int(feature.number),
		Type:   int(feature._type),
	}
	return goFeature
}

func goFeature2C(feature Feature) C.struct_sensors_feature {
	// _type := uint32(feature.Type)
	// var sensorType C.sensors_feature_type = (C.sensors_feature_type)(_type)
	cFeature := C.struct_sensors_feature{
		name:   C.CString(feature.Name),
		number: C.int(feature.Number),
		// _type:  sensorType,
	}
	return cFeature
}
