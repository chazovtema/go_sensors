package go_sensors

type Chip struct {
	Prefix string
	Bus    BusId
	Addr   int
	Path   string
}

type BusId struct {
	Type int
	Nr   int
}

type Feature struct {
	Name   string
	Number int
	Type   int
}
