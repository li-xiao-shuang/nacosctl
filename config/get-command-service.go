package config

import (
	"encoding/json"
	"fmt"
	"nacos-cli/config/constant"
	"nacos-cli/config/model"
	"nacos-cli/http"
	"nacos-cli/logger"
)

func GetCommand(dataId string, group string, address string, port string) {
	constant.Prefix = fmt.Sprintf(constant.Prefix, address, port)
	resp := http.Get(constant.Prefix + constant.ConfigUrl + "?dataId=" + dataId + "&" + "group=" + group + "&show=all")
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
	}
}
