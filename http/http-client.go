package http

import (
	"io/ioutil"
	"nacos-cli/logger"
	"net/http"
)

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		logger.Logger{}.Info("get response is error,error:%s", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Logger{}.Info("reads an error,error:%s ", err)
		return ""
	}
	if resp.StatusCode == 200 {
		return string(body)
	}
	return ""
}
