package config

import (
	"encoding/json"
	"fmt"
	"nacos-cli/config/constant"
	"nacos-cli/config/model"
	"nacos-cli/http"
	"nacos-cli/logger"
	"net/url"
)

//GetCommand get命令处理逻辑
func GetCommand(namespaceId string, dataId string, group string, address string, port string) {
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	resp := http.Get(prefix + constant.ConfigUrl + "?tenant=" + namespaceId + "&dataId=" + dataId + "&group=" + group + "&show=all")
	configItem := &model.ConfigItem{}
	if resp != "" {
		err := json.Unmarshal([]byte(resp), configItem)
		if err != nil {
			logger.Logger{}.Info("json parse error,configItem:%s", configItem)
		}
		fmt.Println("id: " + configItem.Id)
		fmt.Println("dataId: " + configItem.DataId)
		fmt.Println("group: " + configItem.Group)
		fmt.Println("content: " + configItem.Content)
		fmt.Println("md5: " + configItem.Md5)
		fmt.Println("tenant: " + configItem.Tenant)
		fmt.Println("type: " + configItem.Type)
	} else {
		fmt.Println("null")
	}
}

// AddCommand add 命令处理逻辑
func AddCommand(namespaceId string, dataId string, group string, content string, typeStr string, address string, port string) {
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	payload := url.Values{"namespaceId": {namespaceId}, "dataId": {dataId}, "group": {group}, "content": {content}, "type": {typeStr}}
	resp := http.Post(prefix+constant.ConfigUrl, payload)
	if resp != "" {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}
