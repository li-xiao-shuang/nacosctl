package constant

const (
	Prefix = "http://%s:%s"

	base = "/nacos"

	// ConfigUrl 配置url
	ConfigUrl = base + "/v1/cs/configs"

	// VersionUrl 配置版本url
	VersionUrl = base + "/v1/cs/history/previous"

	// NamespaceUrl 命名空间url
	NamespaceUrl = base + "/v1/console/namespaces"
)
