package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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
		for j := range row {
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

func CheckWin(data [15][15]uint8) uint8 {
	var tempI, tempJ, count int
	for i := range data {
		for j := range data[i] {
			if data[i][j] != 0 {
				//check right
				tempI, tempJ, count = 0, 0, 0
				for j+tempJ < 15 {
					if data[i][j+tempJ] == data[i][j] {
						count++
					} else {
						break
					}
					tempJ++
					if tempJ > 4 {
						break
					}
				}
				if count >= 5 {
					return data[i][j]
				}
				//check down
				tempI, tempJ, count = 0, 0, 0
				for i+tempI < 15 {
					if data[i+tempI][j] == data[i][j] {
						count++
					} else {
						break
					}
					tempI++
					if tempI > 4 {
						break
					}
				}
				if count >= 5 {
					return data[i][j]
				}
				//check right-down
				tempI, tempJ, count = 0, 0, 0
				for i+tempI < 15 && j+tempJ < 15 {
					if data[i+tempI][j+tempJ] == data[i][j] {
						count++
					} else {
						break
					}
					tempI++
					tempJ++
					if tempI > 4 || tempJ > 4 {
						break
					}
				}
				if count >= 5 {
					return data[i][j]
				}
				//check right-up
				tempI, tempJ, count = 0, 0, 0
				for i-tempI >= 0 && j+tempJ < 15 {
					if data[i-tempI][j+tempJ] == data[i][j] {
						count++
					} else {
						break
					}
					tempI++
					tempJ++
					if tempI > 4 || tempJ > 4 {
						break
					}
				}
				if count >= 5 {
					return data[i][j]
				}
			}
		}
	}
	return 0
}

func ParseMove(move string) ([15][15]uint8, error) {
	if match, _ := regexp.MatchString(config.MovePattern, move); !match {
		fmt.Println("Input error! example: 12-C")
		return config.Map, fmt.Errorf("input error")
	}
	move = strings.ToUpper(move)
	sp := strings.Split(move, "-")
	num, _ := strconv.Atoi(sp[0])
	cha := sp[1][0]
	if config.Map[num-1][cha-'A'] != 0 {
		return config.Map, fmt.Errorf("chess exist")
	} else {
		config.Map[num-1][cha-'A'] = config.MyChess
	}
	return config.Map, nil
}
