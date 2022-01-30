package namespace

import (
	"encoding/json"
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"nacosctl/common"
	"nacosctl/common/http"
	"nacosctl/printer"
	"nacosctl/process/constant"
	"nacosctl/process/model"
	"net/url"
)

// ParseCreateNamespaceCmd 解析创建命名空间命令
func ParseCreateNamespaceCmd(cmd *cobra.Command, namespaceName string, namespaceDesc string, namespaceId string) {
	address := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address)
	payload := url.Values{"customNamespaceId": {namespaceId}, "namespaceName": {namespaceName}, "namespaceDesc": {namespaceDesc}}
	resp := http.Post(prefix+constant.NamespaceUrl, payload)
	if resp != "" {
		printer.Cyan("done")
		return
	}
	printer.Red("fail")
}

// ParseDeleteNamespaceCmd 解析删除命名空间命令
func ParseDeleteNamespaceCmd(cmd *cobra.Command, namespaceId string) {
	address := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address)
	resp := http.Delete(prefix + constant.NamespaceUrl + "?namespaceId=" + namespaceId)
	if resp != "" {
		printer.Cyan("done")
		return
	}
	printer.Red("fail")
}

// ParseUpdateNamespaceCmd 解析更新命名空间命令
func ParseUpdateNamespaceCmd(cmd *cobra.Command, namespaceName string, namespaceDesc string, namespaceId string) {
	address := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address)
	resp := http.Put(prefix + constant.NamespaceUrl + "?namespace=" + namespaceId + "&namespaceShowName=" + namespaceName + "&namespaceDesc=" + namespaceDesc)
	if resp != "" {
		printer.Cyan("done")
		return
	}
	printer.Red("fail")
}

// ParseGetNamespaceListCmd 解析查询命名空间列表命令
func ParseGetNamespaceListCmd(cmd *cobra.Command) {
	address := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address)
	resp := http.Get(prefix + constant.NamespaceUrl)
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
func ParseGetNamespaceCmd(cmd *cobra.Command, namespaceId string) {
	address := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address)
	resp := http.Get(prefix + constant.NamespaceUrl + "?show=all&namespaceId=" + namespaceId)

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
