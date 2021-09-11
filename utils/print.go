package utils

import (
	"fmt"
	"regexp"

	"github.com/tuzi1412/touchFish_gobang/config"
)

func PrintMap(data [15][15]uint8) {
	CallClear()
}

func CheckWin(data [15][15]uint8) bool {
	return false
}

func ParseMove(move string) ([15][15]uint8, error) {
	if match, _ := regexp.MatchString(config.MovePattern, move); !match {
		fmt.Println("Input error! example: 12-C")
		return config.Map, fmt.Errorf("Input error!")
	}
	return config.Map, nil
}
