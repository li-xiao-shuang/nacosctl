package constant

const (
	Prefix = "http://%s"

	base = "/nacos/v1"

	// ConfigUrl 配置url
	ConfigUrl = base + "/cs/configs"

	// VersionUrl 配置版本url
	VersionUrl = base + "/cs/history/previous"

	// NamespaceUrl 命名空间url
	NamespaceUrl = base + "/console/namespaces"

	// LoginUrl 登录url
	LoginUrl = base + "/auth/login"

	// UserUrl 用户url
	UserUrl = base + "/auth/users"
)
