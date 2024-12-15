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
	Name            string
	Number          int
	Type            SensorsFeatureType
	FirstSubfeature int
	Padding         int
}

type SubFeature struct {
	Name    string
	Number  int
	Type    SensorsSubfeatureType
	Mapping int
	Flags   uint
}
