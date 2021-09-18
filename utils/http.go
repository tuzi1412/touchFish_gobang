package utils

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// SendURL sends http request by url
func SendURL(method, url string, body io.Reader, header map[string]string) (io.ReadCloser, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(network, addr, time.Second*3) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		},
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = http.Header{}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("response is not 200")
	}
	return res.Body, nil
}

//GenHeader GenHeader
func GenHeader() map[string]string {
	header := map[string]string{"Content-Type": "application/json"}
	return header
}
