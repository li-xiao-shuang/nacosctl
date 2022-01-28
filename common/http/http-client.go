package http

import (
	"io/ioutil"
	"nacosctl/common/logger"
	"net/http"
)

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		logger.Logger{}.Info("get request error,error:%s", err)
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

func Post(url string, param map[string][]string) string {
	resp, err := http.PostForm(url, param)
	if err != nil {
		logger.Logger{}.Info("post request  error,error:%s", err)
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

func Delete(url string) string {
	req, _ := http.NewRequest("DELETE", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Logger{}.Info("delete request error,error:%s", err)
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

func Put(url string) string {
	req, _ := http.NewRequest("PUT", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Logger{}.Info("put request error,error:%s", err)
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
