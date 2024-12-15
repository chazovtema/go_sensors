package go_sensors

import "math"

type SensorsFeatureType int

const (
	SensorsFeatureIn SensorsFeatureType = iota
	SensorsFeatureFan
	SensorsFeatureTemp
	SensorsFeaturePower
	SensorsFeatureEnergy
	SensorsFeatureCurr
	SensorsFeatureHumidity
	SensorsFeatureMaxMain
	SensorsFeatureVid = 0x10
	SensorsFeatureIntrusion
	SensorsFeatureMaxOther
	SensorsFeatureBeepEnable = 0x18
	SensorsFeatureMax
	SensorsFeatureUnknown = math.MaxInt32
)

type SensorsSubfeatureType int

const (
	SensorsSubfeatureInInput = SensorsFeatureIn << 8
	SensorsSubfeatureInMin
	SensorsSubfeatureInMax
	SensorsSubfeatureInLcrit
	SensorsSubfeatureInCrit
	SensorsSubfeatureInAverage
	SensorsSubfeatureInLowest
	SensorsSubfeatureInHighest
	SensorsSubfeatureInAlarm = (SensorsFeatureIn << 8) | 0x80
	SensorsSubfeatureInMinAlarm
	SensorsSubfeatureInMaxAlarm
	SensorsSubfeatureInBeep
	SensorsSubfeatureInLcritAlarm
	SensorsSubfeatureInCritAlarm

	SensorsSubfeatureFanInput = SensorsFeatureFan << 8
	SensorsSubfeatureFanMin
	SensorsSubfeatureFanMax
	SensorsSubfeatureFanAlarm = (SensorsFeatureFan << 8) | 0x80
	SensorsSubfeatureFanFault
	SensorsSubfeatureFanDiv
	SensorsSubfeatureFanBeep
	SensorsSubfeatureFanPulses
	SensorsSubfeatureFanMinAlarm
	SensorsSubfeatureFanMaxAlarm

	SensorsSubfeatureTempInput = SensorsFeatureTemp << 8
	SensorsSubfeatureTempMax
	SensorsSubfeatureTempMaxHyst
	SensorsSubfeatureTempMin
	SensorsSubfeatureTempCrit
	SensorsSubfeatureTempCritHyst
	SensorsSubfeatureTempLcrit
	SensorsSubfeatureTempEmergency
	SensorsSubfeatureTempEmergencyHyst
	SensorsSubfeatureTempLowest
	SensorsSubfeatureTempHighest
	SensorsSubfeatureTempMinHyst
	SensorsSubfeatureTempLcritHyst
	SensorsSubfeatureTempAlarm = (SensorsFeatureTemp << 8) | 0x80
	SensorsSubfeatureTempMaxAlarm
	SensorsSubfeatureTempMinAlarm
	SensorsSubfeatureTempCritAlarm
	SensorsSubfeatureTempFault
	SensorsSubfeatureTempType
	SensorsSubfeatureTempOffset
	SensorsSubfeatureTempBeep
	SensorsSubfeatureTempEmergencyAlarm
	SensorsSubfeatureTempLcritAlarm

	SensorsSubfeaturePowerAverage = SensorsFeaturePower << 8
	SensorsSubfeaturePowerAverageHighest
	SensorsSubfeaturePowerAverageLowest
	SensorsSubfeaturePowerInput
	SensorsSubfeaturePowerInputHighest
	SensorsSubfeaturePowerInputLowest
	SensorsSubfeaturePowerCap
	SensorsSubfeaturePowerCapHyst
	SensorsSubfeaturePowerMax
	SensorsSubfeaturePowerCrit
	SensorsSubfeaturePowerMin
	SensorsSubfeaturePowerLcrit
	SensorsSubfeaturePowerAverageInterval = (SensorsFeaturePower << 8) | 0x80
	SensorsSubfeaturePowerAlarm
	SensorsSubfeaturePowerCapAlarm
	SensorsSubfeaturePowerMaxAlarm
	SensorsSubfeaturePowerCritAlarm
	SensorsSubfeaturePowerMinAlarm
	SensorsSubfeaturePowerLcritAlarm

	SensorsSubfeatureEnergyInput = SensorsFeatureEnergy << 8

	SensorsSubfeatureCurrInput = SensorsFeatureCurr << 8
	SensorsSubfeatureCurrMin
	SensorsSubfeatureCurrMax
	SensorsSubfeatureCurrLcrit
	SensorsSubfeatureCurrCrit
	SensorsSubfeatureCurrAverage
	SensorsSubfeatureCurrLowest
	SensorsSubfeatureCurrHighest
	SensorsSubfeatureCurrAlarm = (SensorsFeatureCurr << 8) | 0x80
	SensorsSubfeatureCurrMinAlarm
	SensorsSubfeatureCurrMaxAlarm
	SensorsSubfeatureCurrBeep
	SensorsSubfeatureCurrLcritAlarm
	SensorsSubfeatureCurrCritAlarm

	SensorsSubfeatureHumidityInput = SensorsFeatureHumidity << 8

	SensorsSubfeatureVid = SensorsFeatureVid << 8

	SensorsSubfeatureIntrusionAlarm = SensorsFeatureIntrusion << 8
	SensorsSubfeatureIntrusionBeep

	SensorsSubfeatureBeepEnable = SensorsFeatureBeepEnable << 8

	SensorsSubfeatureUnknown = math.MaxInt32
)
