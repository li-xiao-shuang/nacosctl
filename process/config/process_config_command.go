package config

import (
	"encoding/json"
	"fmt"
	"github.com/scylladb/termtables"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"nacosctl/common"
	"nacosctl/common/http"
	"nacosctl/common/logger"
	"nacosctl/config/constant"
	"nacosctl/config/model"
	"net/url"
	"strings"
)

// ParseCreateConfigCmd 解析创建配置命令
func ParseCreateConfigCmd(namespaceId string, dataId string, group string, content string, typeStr string, cmd *cobra.Command) {
	address, port := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	payload := url.Values{"namespaceId": {namespaceId}, "dataId": {dataId}, "group": {group}, "content": {content}, "type": {typeStr}}
	resp := http.Post(prefix+constant.ConfigUrl, payload)
	if resp != "" {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}

// ParseDeleteConfigCmd 解析删除配置命令
func ParseDeleteConfigCmd(namespaceId string, dataId string, group string, cmd *cobra.Command) {
	address, port := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	resp := http.Delete(prefix + constant.ConfigUrl + "?tenant=" + namespaceId + "&dataId=" + dataId + "&group=" + group)
	if resp != "" {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}

// ParseGetConfigCmd 解析查询配置命令
func ParseGetConfigCmd(namespaceId string, dataId string, group string, cmd *cobra.Command) {
	address, port := common.GetServerAddress(cmd)
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

// ParseGetConfigListCmd 解析查询配置列表命令
func ParseGetConfigListCmd(pageNo string, pageSize string, cmd *cobra.Command) {
	address, port := common.GetServerAddress(cmd)
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