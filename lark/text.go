package lark

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type messageRequest struct {
	Email   string            `json:"email"`
	MsgType string            `json:"msg_type"`
	Content map[string]string `json:"content"`
}

type messageResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		MessageID string `json:"message_id"`
	} `json:"data"`
}

func (b *Bot) NotifyText(email, msg string) error {
	const sendUrl = "https://open.feishu.cn/open-apis/message/v4/send/"
	// 构造请求结构体
	req := messageRequest{
		Email:   email,
		MsgType: "text",
		Content: map[string]string{"text": msg},
	}
	buf, err := json.Marshal(req)
	if err != nil {
		return err
	}
	token, err := b.getToken()
	if err != nil {
		panic(err)
	}
	header := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}
	resp, err := b.client.PostWithHeader(sendUrl, header, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	var textResp messageResponse
	err = json.Unmarshal(content, &textResp)
	if err != nil {
		return err
	}
	// 结果判断
	if textResp.Code != successCode {
		return fmt.Errorf("%d %s: %s", textResp.Code, textResp.Msg, MapError(textResp.Code))
	}
	return nil
}
