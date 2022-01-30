package model

type ConfigItem struct {
	Id      string `json:"id"`
	DataId  string `json:"dataId"`
	Group   string `json:"group"`
	Content string `json:"content"`
	Md5     string `json:"md5"`
	Tenant  string `json:"tenant"`
	Type    string `json:"type"`
}
