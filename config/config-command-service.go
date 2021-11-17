package config

import (
	"encoding/json"
	"fmt"
	"github.com/scylladb/termtables"
	"github.com/tidwall/gjson"
	"nacos-cli/common/http"
	"nacos-cli/common/logger"
	"nacos-cli/config/constant"
	"nacos-cli/config/model"
	"net/url"
	"strings"
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
		jsonByte, _ := json.MarshalIndent(configItem, "", "    ")
		fmt.Println(string(jsonByte))
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

// DelCommand del 命令处理逻辑
func DelCommand(namespaceId string, dataId string, group string, address string, port string) {
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	resp := http.Delete(prefix + constant.ConfigUrl + "?tenant=" + namespaceId + "&dataId=" + dataId + "&group=" + group)
	if resp != "" {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}

// ListCommand list 命令处理逻辑
func ListCommand(pageNo string, pageSize string, address string, port string) {
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	resp := http.Get(prefix + constant.ConfigUrl + "?pageNo=" + pageNo + "&pageSize=" + pageSize + "&dataId=&group=&search=blur")
	if resp == "" {
		fmt.Println("fail")
	} else {
		items := gjson.Get(resp, "pageItems").String()
		configItems := &[]model.ConfigItem{}
		err := json.Unmarshal([]byte(items), configItems)
		if err != nil {
			logger.Logger{}.Info("json parse error,configItem:%s", configItems)
		}
		//创建表格
		t := termtables.CreateTable()
		t.AddHeaders("id", "dataId", "group", "content")
		for _, item := range *configItems {
			subStr := ""
			if strings.Contains(item.Content, "\n") {
				str := strings.Split(item.Content, "\n")
				subStr = str[0] + "..."
			} else {
				subStr = item.Content
			}
			if len(subStr) > 80 {
				subStr = item.Content[0:80] + "..."
			}
			t.AddRow(item.Id, item.DataId, item.Group, subStr)
		}
		fmt.Println(t.Render())
	}
}

// VersionCommand version 命令处理逻辑
func VersionCommand(id string, namespaceId string, dataId string, group string, address string, port string) {
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	resp := http.Get(prefix + constant.VersionUrl + "?id=" + id + "&tenant=" + namespaceId + "&dataId=" + dataId + "&group=" + group)
	if resp != "" {
		configItem := &model.ConfigItem{}
		json.Unmarshal([]byte(resp), configItem)
		jsonByte, _ := json.MarshalIndent(configItem, "", "    ")
		fmt.Println(string(jsonByte))
	} else {
		fmt.Println("null")
	}
}
