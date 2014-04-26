package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func curl(url string) ([]byte, error) {
	ts := &http.Transport{}
	client := http.Client{
		Transport: ts,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("referer", "http://yun.baidu.com/share/home?uk=2822032057&view=share")
	rsp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rsp.Body.Close()
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return b, nil
}
