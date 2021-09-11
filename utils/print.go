package utils

import (
	"fmt"
	"regexp"

	"github.com/tuzi1412/touchFish_gobang/config"
)

func PrintMap(data [15][15]uint8) {
	CallClear()
	fmt.Println("    A  B  C  D  E  F  G  H  I  J  K  L  M  N  O  ")
	for i, row := range data {
		if i+1 < 10 {
			fmt.Printf("%d   ", i+1)
		} else {
			fmt.Printf("%d  ", i+1)
		}
		for j, _ := range row {
			if data[i][j] == 0 {
				fmt.Print("   ")
			} else if data[i][j] == 1 {
				fmt.Print("-  ")
			} else if data[i][j] == 2 {
				fmt.Print("+  ")
			}
			if j == 14 {
				fmt.Println("")
			}
		}
	}
	fmt.Println("")
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
