package config

import (
	"encoding/json"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"nacosctl/common/http"
	"nacosctl/common/logger"
	"nacosctl/printer"
	"nacosctl/process/model"
	"net/url"
)

// ParseCreateConfigCmd 解析创建配置命令
func ParseCreateConfigCmd(cmd *cobra.Command, namespaceId string, dataId string, group string, content string, typeStr string) {
	serverUrl := http.GetConfigUrl(cmd)
	payload := url.Values{"tenant": {namespaceId}, "dataId": {dataId}, "group": {group}, "content": {content}, "type": {typeStr}}
	resp := http.Post(serverUrl, payload)
	if resp != "" {
		printer.Cyan("done")
	} else {
		printer.Red("fail")
	}
}

// ParseDeleteConfigCmd 解析删除配置命令
func ParseDeleteConfigCmd(cmd *cobra.Command, dataId string, namespaceId string, group string) {
	serverUrl := http.GetConfigUrl(cmd)
	resp := http.Delete(serverUrl + "&dataId=" + dataId + "&group=" + group + "&tenant=" + namespaceId)
	if resp != "" {
		printer.Cyan("done")
	} else {
		printer.Red("fail")
	}
}

// ParseGetConfigCmd 解析查询配置命令
func ParseGetConfigCmd(cmd *cobra.Command, dataId string, namespaceId string, group string) {
	serverUrl := http.GetConfigUrl(cmd)
	resp := http.Get(serverUrl + "&tenant=" + namespaceId + "&dataId=" + dataId + "&group=" + group + "&show=all")

	table := printer.NewTableWrap(300, true)
	if resp == "" {
		table.AddRow(aurora.Magenta("ID:"))
		table.AddRow(aurora.Magenta("DataId:"))
		table.AddRow(aurora.Magenta("命名空间:"))
		table.AddRow(aurora.Magenta("组名:"))
		table.AddRow(aurora.Magenta("MD5:"))
		table.AddRow(aurora.Magenta("配置类型:"))
		table.AddRow(aurora.Magenta("配置内容:"))
		fmt.Println(table)
		return
	}
	configItem := &model.ConfigItem{}
	err := json.Unmarshal([]byte(resp), configItem)
	if err != nil {
		logger.Logger{}.Info("json parse error,configItem:%s", configItem)
	}
	table.AddRow(aurora.Magenta("ID:"), configItem.Id)
	table.AddRow(aurora.Magenta("DataId:"), configItem.DataId)
	table.AddRow(aurora.Magenta("命名空间:"), configItem.Tenant)
	table.AddRow(aurora.Magenta("组名:"), configItem.Group)
	table.AddRow(aurora.Magenta("MD5:"), configItem.Md5)
	table.AddRow(aurora.Magenta("配置类型:"), configItem.Type)
	table.AddRow(aurora.Magenta("配置内容:"), configItem.Content)
	fmt.Println(table)
}

// ParseGetConfigListCmd 解析查询配置列表命令
func ParseGetConfigListCmd(cmd *cobra.Command, pageNo string, pageSize string, namespaceId string, group string) {
	serverUrl := http.GetConfigUrl(cmd)
	resp := http.Get(serverUrl + "&pageNo=" + pageNo + "&pageSize=" + pageSize + "&tenant=" + namespaceId + "&group=" + group + "&dataId=&search=blur")
	table := printer.NewTableWrap(100, false)
	table.AddRow(aurora.Magenta("ID:"), aurora.Magenta("DataId:"), aurora.Magenta("命名空间:"),
		aurora.Magenta("组名:"), aurora.Magenta("配置类型:"), aurora.Magenta("配置内容:"))
	if resp == "" {
		fmt.Println(table)
		return
	}
	items := gjson.Get(resp, "pageItems").String()
	configItems := &[]model.ConfigItem{}
	err := json.Unmarshal([]byte(items), configItems)
	if err != nil {
		logger.Logger{}.Info("json parse error,configItem:%s", configItems)
	}
	for _, item := range *configItems {
		table.AddRow(item.Id, item.DataId, item.Tenant, item.Group, item.Type, item.Content)
	}
	fmt.Println(table)
}

// ParseConfigVersionCmd 解析配置版本命令
func ParseConfigVersionCmd(cmd *cobra.Command, id string, namespaceId string, dataId string, group string) {
	url := http.GetConfigVersionUrl(cmd)
	resp := http.Get(url + "&id=" + id + "&tenant=" + namespaceId + "&dataId=" + dataId + "&group=" + group)

	table := printer.NewTableWrap(300, true)
	if resp == "" {
		table.AddRow(aurora.Magenta("ID:"))
		table.AddRow(aurora.Magenta("DataId:"))
		table.AddRow(aurora.Magenta("命名空间:"))
		table.AddRow(aurora.Magenta("组名:"))
		table.AddRow(aurora.Magenta("MD5:"))
		table.AddRow(aurora.Magenta("配置类型:"))
		table.AddRow(aurora.Magenta("配置内容:"))
		fmt.Println(table)
		return
	}
	configItem := &model.ConfigItem{}
	err := json.Unmarshal([]byte(resp), configItem)
	if err != nil {
		logger.Logger{}.Info("json parse error,configItem:%s", configItem)
	}
	table.AddRow(aurora.Magenta("ID:"), configItem.Id)
	table.AddRow(aurora.Magenta("DataId:"), configItem.DataId)
	table.AddRow(aurora.Magenta("命名空间:"), configItem.Tenant)
	table.AddRow(aurora.Magenta("组名:"), configItem.Group)
	table.AddRow(aurora.Magenta("MD5:"), configItem.Md5)
	table.AddRow(aurora.Magenta("配置类型:"), configItem.Type)
	table.AddRow(aurora.Magenta("配置内容:"), configItem.Content)
	fmt.Println(table)
}
