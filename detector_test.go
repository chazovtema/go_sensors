package go_sensors

import (
	"fmt"
	"testing"
)

func TestNewDetector(t *testing.T) {
	detector, err := NewDetector()
	if err != nil {
		t.Fatal(err)
	}
	defer detector.Close()
	chips := detector.Detect()
	chip := chips[0]
	fmt.Println(chip.Features[0])
}
