package model

type LoginInfo struct {
	AccessToken string `json:"accessToken"`
	TokenTtl    int64 `json:"tokenTtl"`
}
