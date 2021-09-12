package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/tuzi1412/touchFish_gobang/api"
	"github.com/tuzi1412/touchFish_gobang/config"
	"github.com/tuzi1412/touchFish_gobang/utils"
)

func main() {
	err := ipInit()
	if err != nil {
		fmt.Println(err)
		return
	}
	server := api.NewServer()
	server.Run()
	waitConnect()
	gameStart()
	// wait()
}

func ipInit() error {
	for {
		fmt.Println("please input opposite ip:")
		fmt.Scanln(&api.OppositeIP)
		if match, _ := regexp.MatchString(config.Ipv4Pattern, api.OppositeIP); !match {
			fmt.Println("please input a ipv4 address!")
			continue
		}
		break
	}
	return nil
}

func waitConnect() {
	for {
		fmt.Println("connecting...")
		res, err := utils.SendURL("GET", "http://"+api.OppositeIP+":8080/touchFish_gobang/testConnect", nil, utils.GenHeader())
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}
		var resBody []byte
		if res != nil {
			defer res.Close()
			resBody, err = ioutil.ReadAll(res)
			if err != nil {
				time.Sleep(3 * time.Second)
				continue
			}
		}
		var rsp config.HTTPRsp
		json.Unmarshal(resBody, &rsp)
		if rsp.Code != 0 {
			time.Sleep(3 * time.Second)
			continue
		}
		fmt.Println("connect success!")
		break
	}
}

func gameStart() {
	firstHand()
LOOP:
	for {
		select {
		case data := <-config.MapChan:
			utils.PrintMap(data)
			win := utils.CheckWin(data)
			if win != 0 && win != config.MyChess {
				fmt.Println("You Lose!")
				for {
					fmt.Print("Start new game?[Y/N]:")
					fmt.Scanln(&config.NewGame)
					if strings.ToLower(config.NewGame) == "n" {
						return
					} else if strings.ToLower(config.NewGame) == "y" {
						config.Map = [15][15]uint8{}
						config.MyChess = 0
						select {
						case <-config.MapChan:
							//do nothing, clean channel
						case <-time.After(time.Second):
							break
						}
						firstHand()
						continue LOOP
					} else {
						continue
					}
				}
			}
			config.Map = data
			for {
				if config.MyChess == 1 {
					fmt.Print(`Your chess is "-", `)
				} else if config.MyChess == 2 {
					fmt.Print(`Your chess is "+", `)
				}
				fmt.Println("Your move(example: 12-C):")
				fmt.Scanln(&config.Move)
				result, err := utils.ParseMove(config.Move)
				if err != nil {
					fmt.Println(err)
					continue
				}
				err = api.SendData(result)
				if err != nil {
					fmt.Println(err)
					continue
				}
				utils.PrintMap(config.Map)
				win := utils.CheckWin(config.Map)
				if win != 0 && win == config.MyChess {
					fmt.Println("You Win!")
					for {
						fmt.Print("Start new game?[Y/N]:")
						fmt.Scanln(&config.NewGame)
						if strings.ToLower(config.NewGame) == "n" {
							return
						} else if strings.ToLower(config.NewGame) == "y" {
							config.Map = [15][15]uint8{}
							config.MyChess = 0
							select {
							case <-config.MapChan:
								//do nothing, clean channel
							case <-time.After(time.Second):
								break
							}
							firstHand()
							continue LOOP
						} else {
							continue
						}
					}
				}
				break
			}
		}
	}
}

func firstHand() error {
	var msg config.HTTPRsp
	rand.Seed(time.Now().UnixNano())
	myRandomNum := rand.Int()

	msg.Code = myRandomNum
	msg.Message = "success"

	byteData, _ := json.Marshal(msg)
	body := bytes.NewReader(byteData)

	res, err := utils.SendURL("PUT", "http://"+api.OppositeIP+":8080/touchFish_gobang/firstHand", body, utils.GenHeader())
	if err != nil {
		return err
	}

	var resBody []byte
	if res != nil {
		defer res.Close()
		resBody, err = ioutil.ReadAll(res)
		if err != nil {
			return err
		}
	}
	var rsp config.HTTPRsp
	json.Unmarshal(resBody, &rsp)
	if config.MyChess == 0 {
		if myRandomNum >= msg.Code {
			config.MyChess = 1
		} else {
			config.MyChess = 1
		}
	}
	if config.MyChess == 1 {
		utils.PrintMap(config.Map)
		for {
			if config.MyChess == 1 {
				fmt.Print(`Your chess is "-", `)
			} else if config.MyChess == 2 {
				fmt.Print(`Your chess is "+", `)
			}
			fmt.Println("Your move(example: 12-C):")
			fmt.Scanln(&config.Move)
			result, err := utils.ParseMove(config.Move)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = api.SendData(result)
			if err != nil {
				fmt.Println(err)
				continue
			}
			utils.PrintMap(config.Map)
			break
		}
	} else {
		fmt.Println("waiting for  your opponent's move")
	}
	return nil
}

func wait() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	signal.Ignore(syscall.SIGPIPE)
	select {
	case <-sig:
	}
}
