package lark

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type tokenRequest struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type tokenResponse struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}

func GetAccessToken(appID, appSecret string) (string, error) {
	// token专用url
	const url = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/"
	req := tokenRequest{
		AppID:     appID,
		AppSecret: appSecret,
	}
	// marshal 99%的情况下都不会出问题
	buf, _ := json.Marshal(req)
	resp, err := http.Post(url, "application/json", bytes.NewReader(buf))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// 反序列化
	var tokenResp tokenResponse
	err = json.Unmarshal(buf, &tokenResp)
	if err != nil {
		panic(err)
	}
	log.Printf("token resp: %+v\n", tokenResp)
	return tokenResp.TenantAccessToken, nil
}
