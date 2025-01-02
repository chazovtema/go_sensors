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
	"sync"
)

var (
	instance *Detector
	once     sync.Once
	initErr  error = nil
)

type Detector struct {
}

func NewDetector() (*Detector, error) {
	once.Do(func() {
		if ret := C.sensors_init(nil); ret != 0 {
			initErr = fmt.Errorf("can not init sensors package")
		} else {
			initErr = nil
		}
		instance = &Detector{}
	})
	return instance, initErr
}

func (d *Detector) Close() {
	C.sensors_cleanup()
}

func (d *Detector) Detect() []Chip {
	wrappedChips := detectChips()
	chips := make([]Chip, len(wrappedChips))

	for i, chip := range wrappedChips {
		chip := Chip{
			Name: chip.Prefix,
			chip: chip,
		}
		chip.Refresh()
		chips[i] = chip
	}
	return chips
}
