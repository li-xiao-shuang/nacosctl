package namespace

import (
	"encoding/json"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/tidwall/gjson"
	"nacosctl/common/http"
	"nacosctl/printer"
	"nacosctl/process/model"
	"net/url"
)

// ParseCreateNamespaceCmd 解析创建命名空间命令
func ParseCreateNamespaceCmd(namespaceName string, namespaceDesc string, namespaceId string) {
	serverUrl := http.GetNamespaceUrl()
	payload := url.Values{"customNamespaceId": {namespaceId}, "namespaceName": {namespaceName}, "namespaceDesc": {namespaceDesc}}
	resp := http.Post(serverUrl, payload)
	if resp != "" {
		printer.Cyan("done")
		return
	}
	printer.Red("fail")
}

// ParseDeleteNamespaceCmd 解析删除命名空间命令
func ParseDeleteNamespaceCmd(namespaceId string) {
	url := http.GetNamespaceUrl()
	resp := http.Delete(url + "&namespaceId=" + namespaceId)
	if resp != "" {
		printer.Cyan("done")
		return
	}
	printer.Red("fail")
}

// ParseUpdateNamespaceCmd 解析更新命名空间命令
func ParseUpdateNamespaceCmd(namespaceName string, namespaceDesc string, namespaceId string) {
	url := http.GetNamespaceUrl()
	resp := http.Put(url + "&namespace=" + namespaceId + "&namespaceShowName=" + namespaceName + "&namespaceDesc=" + namespaceDesc)
	if resp != "" {
		printer.Cyan("done")
		return
	}
	printer.Red("fail")
}

// ParseGetNamespaceListCmd 解析查询命名空间列表命令
func ParseGetNamespaceListCmd() {
	url := http.GetNamespaceUrl()
	resp := http.Get(url)
	table := printer.NewTable(100)
	table.AddRow(aurora.Magenta("命名空间ID"), aurora.Magenta("命名空间名称"), aurora.Magenta("配置额度"), aurora.Magenta("使用数量"))
	if resp == "" {
		fmt.Println(table)
		return
	}
	items := gjson.Get(resp, "data").String()
	namespaces := &[]model.NamespaceInfo{}
	json.Unmarshal([]byte(items), namespaces)
	for _, namespace := range *namespaces {
		table.AddRow(namespace.Namespace, namespace.NamespaceShowName, namespace.Quota, namespace.ConfigCount)
	}
	fmt.Println(table)
}

// ParseGetNamespaceCmd 解析查询命名空间命令
func ParseGetNamespaceCmd(namespaceId string) {
	url := http.GetNamespaceUrl()
	resp := http.Get(url + "&show=all&namespaceId=" + namespaceId)

	table := printer.NewTable(100)
	table.AddRow(aurora.Magenta("命名空间ID"), aurora.Magenta("命名空间名称"), aurora.Magenta("配置额度"), aurora.Magenta("使用数量"))

	if resp == "" {
		fmt.Println(table)
		return
	}
	namespace := model.NewNamespaceInfo()
	json.Unmarshal([]byte(resp), namespace)
	table.AddRow(namespace.Namespace, namespace.NamespaceShowName, namespace.Quota, namespace.ConfigCount)
	fmt.Println(table)
}
