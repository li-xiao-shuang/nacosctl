package http

import (
	"io/ioutil"
	"nacosctl/common/logger"
	"nacosctl/printer"
	"net/http"
	"os"
)

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		printer.Yellow(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		printer.Yellow(err)
		return ""
	}
	if resp.StatusCode == 403 {
		printer.Red("[ERROR] Please log in your account first")
		os.Exit(0)
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
		printer.Yellow(err)
		return ""
	}
	if resp.StatusCode == 403 {
		printer.Red("[ERROR] Please log in your account first")
		os.Exit(0)
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
		printer.Yellow(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		printer.Yellow(err)
		return ""
	}
	if resp.StatusCode == 403 {
		printer.Red("[ERROR] Please log in your account first")
		os.Exit(0)
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
		printer.Yellow(err)
		return ""
	}
	if resp.StatusCode == 403 {
		printer.Red("[ERROR] Please log in your account first")
		os.Exit(0)
	}
	if resp.StatusCode == 200 {
		return string(body)
	}
	return ""
}
