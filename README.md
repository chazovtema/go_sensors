# Go sensors

Low level [libsensors](https://packages.debian.org/sid/libsensors-dev) wrapper. Unlike other packages, this one uses a dynamic library instead of calling lmsensors from the shell.

# Dependencies

```
sudo apt install libsensors-dev
```

# Usage

```go
package main

import (
	"fmt"
	"github.com/chazovtema/go_sensors/v2"
)


func main() {
	detector, err := go_sensors.NewDetector()
	if err != nil {
		panic(err)
	}
	defer detector.Close()
	chips := detector.Detect()
	chip := chips[0]
	fmt.Println(chip.Features[0])
}
```
