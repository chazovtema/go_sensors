package go_sensors

/*
#cgo LDFLAGS: -lsensors
#include <sensors/sensors.h>
#include <stdlib.h>

*/
import "C"

import "math"

type sensorsFeatureType int

const (
	SensorsFeatureIn         sensorsFeatureType = C.SENSORS_FEATURE_IN
	SensorsFeatureFan        sensorsFeatureType = C.SENSORS_FEATURE_FAN
	SensorsFeatureTemp       sensorsFeatureType = C.SENSORS_FEATURE_TEMP
	SensorsFeaturePower      sensorsFeatureType = C.SENSORS_FEATURE_POWER
	SensorsFeatureEnergy     sensorsFeatureType = C.SENSORS_FEATURE_ENERGY
	SensorsFeatureCurr       sensorsFeatureType = C.SENSORS_FEATURE_CURR
	SensorsFeatureHumidity   sensorsFeatureType = C.SENSORS_FEATURE_HUMIDITY
	SensorsFeatureMaxMain    sensorsFeatureType = C.SENSORS_FEATURE_MAX_MAIN
	SensorsFeatureVid        sensorsFeatureType = C.SENSORS_FEATURE_VID
	SensorsFeatureIntrusion  sensorsFeatureType = C.SENSORS_FEATURE_INTRUSION
	SensorsFeatureMaxOther   sensorsFeatureType = C.SENSORS_FEATURE_MAX_OTHER
	SensorsFeatureBeepEnable sensorsFeatureType = C.SENSORS_FEATURE_BEEP_ENABLE
	SensorsFeatureMax        sensorsFeatureType = C.SENSORS_FEATURE_MAX
	SensorsFeatureUnknown    sensorsFeatureType = C.SENSORS_FEATURE_UNKNOWN
)

type sensorsSubfeatureType int

const (
	SensorsSubfeatureInInput = SensorsFeatureIn << 8
	SensorsSubfeatureInMin
	SensorsSubfeatureInMax
	SensorsSubfeatureInLcrit
	SensorsSubfeatureInCrit
	SensorsSubfeatureInAverage
	SensorsSubfeatureInLowest
	SensorsSubfeatureInHighest
	SensorsSubfeatureInAlarm
	SensorsSubfeatureInMinAlarm
	SensorsSubfeatureInMaxAlarm
	SensorsSubfeatureInBeep
	SensorsSubfeatureInLcritAlarm
	SensorsSubfeatureInCritAlarm

	SensorsSubfeatureFanInput
	SensorsSubfeatureFanMin
	SensorsSubfeatureFanMax
	SensorsSubfeatureFanAlarm
	SensorsSubfeatureFanFault
	SensorsSubfeatureFanDiv
	SensorsSubfeatureFanBeep
	SensorsSubfeatureFanPulses
	SensorsSubfeatureFanMinAlarm
	SensorsSubfeatureFanMaxAlarm

	SensorsSubfeatureTempInput
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
