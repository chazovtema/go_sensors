package go_sensors

type Chip struct {
	Name     string
	Features []Feature

	chip goChip
}

func (c *Chip) Refresh() {
	wrappedFeatures := getFeatures(c.chip)
	c.Features = make([]Feature, len(wrappedFeatures))
	for i, feature := range wrappedFeatures {
		c.Features[i] = Feature{
			Name:        feature.Name,
			Type:        feature.Type,
			SubFeatures: c.getSubfeatures(feature),
		}

	}
}

func (c *Chip) getSubfeatures(feature goFeature) []SubFeature {
	wrappedSubFeatures := getSubfeatures(c.chip, feature)
	subFeatures := make([]SubFeature, len(wrappedSubFeatures))

	for i, subFeature := range wrappedSubFeatures {
		subFeatures[i] = SubFeature{
			Name:  subFeature.Name,
			Type:  subFeature.Type,
			Value: subFeature.Value,
		}
	}
	return subFeatures
}

type Feature struct {
	Name        string
	Type        sensorsFeatureType
	SubFeatures []SubFeature
}

type SubFeature struct {
	Name  string
	Type  sensorsSubfeatureType
	Value float64
}
