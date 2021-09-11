package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		fmt.Println("请输入对方ip:")
		fmt.Scanln(&api.OppositeIP)
		if match, _ := regexp.MatchString(config.Ipv4Pattern, api.OppositeIP); !match {
			fmt.Println("请输入标准ipv4地址！")
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
LOOP:
	for {
		select {
		case data := <-config.MapChan:
			utils.PrintMap(data)
			if utils.CheckWin(data) {
				fmt.Println("You Lose!")
				for {
					fmt.Print("Start new game?[Y/N]:")
					fmt.Scanln(&config.NewGame)
					if strings.ToLower(config.NewGame) == "n" {
						return
					} else if strings.ToLower(config.NewGame) == "y" {
						continue LOOP
					} else {
						continue
					}
				}
			}
			for {
				fmt.Println("Your move:")
				fmt.Scanln(&config.Move)
				result, err := utils.ParseMove(config.Move)
				if err != nil {
					fmt.Println(err)
					continue
				}
				err = api.SendData(result)
				if err != nil {

				}
			}
		}
	}
}

func wait() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	signal.Ignore(syscall.SIGPIPE)
	select {
	case <-sig:
	}
}
