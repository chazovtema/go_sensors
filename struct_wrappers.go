package go_sensors

type goChip struct {
	Prefix string
	Bus    goBusId
	Addr   int
	Path   string
}

type goBusId struct {
	Type int
	Nr   int
}

type goFeature struct {
	Name            string
	Number          int
	Type            sensorsFeatureType
	FirstSubfeature int
	Padding         int
}

type goSubFeature struct {
	Name    string
	Number  int
	Type    sensorsSubfeatureType
	Mapping int
	Flags   uint
	Value   float64
}
