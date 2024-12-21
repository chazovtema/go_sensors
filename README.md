# Go sensors

Low level [libsensors](https://packages.debian.org/sid/libsensors-dev) wrapper. Unlike other packages, this one uses a dynamic library instead of calling lmsensors in the shell.

# Dependencies

```
sudo apt install libsensors-dev
```

# Usage

```go
package main

import (
	"fmt"
	"time"
	"github.com/chazovtema/go_sensors"
)


func main() {
	go_sensors.Init()
	defer go_sensors.CleanUp()
	chip, _ := go_sensors.DetectChip(0)
	var lastTemp float64 = 0
	for {
		features := go_sensors.GetFeatures(chip)
		feature := features[0]
		subfeatures := go_sensors.GetSubfeatures(chip, feature)
		subfeature := subfeatures[0]
		if subfeature.Value != lastTemp {
			fmt.Println(subfeature.Value)
			lastTemp = subfeature.Value
		}
		time.Sleep(time.Second) 
	}

}
```
