package namespace

import (
	"encoding/json"
	"fmt"
	"github.com/scylladb/termtables"
	"github.com/spf13/cobra"
	"nacos-cli/common"
	"nacos-cli/common/http"
	"nacos-cli/config/constant"
	"nacos-cli/namespace/model"
)

// list 命令处理逻辑
func ListCommand(cmd *cobra.Command) {
	address, port := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address, port)
	resp := http.Get(prefix + constant.NamespaceUrl)
	if resp == "" {
		fmt.Println("null")
	}
	namespaces := &[]model.NamespaceInfo{}
	json.Unmarshal([]byte(resp), namespaces)
	t := termtables.CreateTable()
	t.AddHeaders("namespace", "showName", "quota", "configCount")
	for _, namespace := range *namespaces {
		t.AddRow(namespace.Namespace, namespace.NamespaceShowName, namespace.Quota, namespace.ConfigCount)
	}
	fmt.Println(t.Render())
}
