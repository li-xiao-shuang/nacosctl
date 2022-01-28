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

// ParseCreateNamespaceCmd 解析创建命名空间命令
func ParseCreateNamespaceCmd(cmd *cobra.Command, namespaceName string, namespaceDesc string, namespaceId string) {
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

// ParseDeleteNamespaceCmd 解析删除命名空间命令
func ParseDeleteNamespaceCmd(cmd *cobra.Command, namespaceId string) {
	address, port := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	resp := http.Delete(prefix + constant.NamespaceUrl + "?namespaceId=" + namespaceId)
	if resp != "" {
		fmt.Println("success")
		return
	}
	fmt.Println("fail")
}

// ParseUpdateNamespaceCmd 解析更新命名空间命令
func ParseUpdateNamespaceCmd(cmd *cobra.Command, namespaceName string, namespaceDesc string, namespaceId string) {
	address, port := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	resp := http.Put(prefix + constant.NamespaceUrl + "?namespace=" + namespaceId + "&namespaceShowName=" + namespaceName + "&namespaceDesc=" + namespaceDesc)
	if resp != "" {
		fmt.Println("success")
		return
	}
	fmt.Println("fail")
}

// ParseGetNamespaceListCmd 解析查询命名空间列表命令
func ParseGetNamespaceListCmd(cmd *cobra.Command) {
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
