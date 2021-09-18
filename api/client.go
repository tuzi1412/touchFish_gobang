package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/tuzi1412/touchFish_gobang/config"
	"github.com/tuzi1412/touchFish_gobang/utils"
)

var OppositeIP string

func SendData(data [15][15]uint8) error {
	var msg config.HTTPRsp
	msg.Code = 0
	msg.Message = "success"
	msg.Data = data

	byteData, _ := json.Marshal(msg)
	body := bytes.NewReader(byteData)

	res, err := utils.SendURL("PUT", "http://"+OppositeIP+":22333/touchFish_gobang/", body, utils.GenHeader())
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
	if rsp.Code == 0 {
		config.Map = rsp.Data
	}
	return nil
}
