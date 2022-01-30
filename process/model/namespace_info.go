package model

type NamespaceInfo struct {
	Namespace         string
	NamespaceShowName string
	Quota             int
	ConfigCount       int
}

func NewNamespaceInfo() *NamespaceInfo {
	return &NamespaceInfo{}
}
