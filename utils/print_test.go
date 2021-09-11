package utils

import (
	"testing"
)

func TestPrintMap(t *testing.T) {
	var temp [15][15]uint8
	temp[6][7] = 1
	temp[6][8] = 2
	PrintMap(temp)
}
