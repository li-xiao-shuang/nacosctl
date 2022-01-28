package namespace

import (
	"encoding/json"
	"fmt"
	"github.com/scylladb/termtables"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"nacos-cli/common"
	"nacos-cli/common/http"
	"nacos-cli/config/constant"
	"nacos-cli/namespace/model"
	"net/url"
)

// ListCommand list 命令处理逻辑
func ListCommand(cmd *cobra.Command) {
	address, port := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	resp := http.Get(prefix + constant.NamespaceUrl)
	if resp == "" {
		fmt.Println("null")
	}
	items := gjson.Get(resp, "data").String()
	namespaces := &[]model.NamespaceInfo{}
	json.Unmarshal([]byte(items), namespaces)
	t := termtables.CreateTable()
	t.AddHeaders("namespace", "showName", "quota", "configCount")
	for _, namespace := range *namespaces {
		t.AddRow(namespace.Namespace, namespace.NamespaceShowName, namespace.Quota, namespace.ConfigCount)
	}
	fmt.Println(t.Render())
}

// AddCommand add命令处理逻辑
func AddCommand(cmd *cobra.Command, namespaceName string, namespaceDesc string, namespaceId string) {
	address, port := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	payload := url.Values{"customNamespaceId": {namespaceId}, "namespaceName": {namespaceName}, "namespaceDesc": {namespaceDesc}}
	resp := http.Post(prefix+constant.NamespaceUrl, payload)
	if resp != "" {
		fmt.Println("success")
		return
	}
	fmt.Println("fail")
}

// DelCommand del命令处理逻辑
func DelCommand(cmd *cobra.Command, namespaceId string) {
	address, port := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	resp := http.Delete(prefix + constant.NamespaceUrl + "?namespaceId=" + namespaceId)
	if resp != "" {
		fmt.Println("success")
		return
	}
	fmt.Println("fail")
}

// UpdateCommand update命令处理逻辑
func UpdateCommand(cmd *cobra.Command, namespaceId string, namespaceName string, namespaceDesc string) {
	address, port := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	resp := http.Put(prefix + constant.NamespaceUrl + "?namespace=" + namespaceId + "&namespaceShowName=" + namespaceName + "&namespaceDesc=" + namespaceDesc)
	if resp != "" {
		fmt.Println("success")
		return
	}
	fmt.Println("fail")
}
