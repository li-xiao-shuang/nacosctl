package config

import (
	"github.com/spf13/cobra"
	"nacosctl/process/config"
	"nacosctl/process/constant"
)

var DeleteConfigCmd = &cobra.Command{
	Use:     "config [dataId]",
	Short:   "Delete a configuration",
	Long:    "Delete a configuration",
	Example: "nacosctl delete config [dataId]",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config.ParseDeleteConfigCmd(cmd, args[0], configNamespaceId, group)
	},
}

func init() {
	DeleteConfigCmd.Flags().StringVarP(&configNamespaceId, "namespaceId", "i", "", "namespace id")
	DeleteConfigCmd.Flags().StringVarP(&group, "group", "g", constant.DefaultGroup, "config group")
}
