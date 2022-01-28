package namespace

import (
	"fmt"
	"github.com/spf13/cobra"
	"nacos-cli/common"
	"nacos-cli/common/http"
	"nacos-cli/config/constant"
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
