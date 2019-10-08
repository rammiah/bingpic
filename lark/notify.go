package lark

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (b *Bot) NotifyImage(img string) error {

	return nil
}

type textResponse struct {
	Code int
	Msg  string
	Data struct {
		MessageID string `json:"message_id"`
	} `json:"data"`
}

func (b *Bot) NotifyText(email, msg string) error {
	/*{
	  "open_id":"ou_5ad573a6411d72b8305fda3a9c15c70e",
	  "root_id":"om_40eb06e7b84dc71c03e009ad3c754195",
	  "chat_id":"oc_5ad11d72b830411d72b836c20",
	  "user_id": "92e39a99",
	  "email":"fanlv@gmail.com",
	  "msg_type":"text",
	  "content":{
	       "text":"text content<at user_id=\"ou_88a56e7e8e9f680b682f6905cc09098e\">test</at>"
	   }}
	*/
	s := `{"email":"rammiahcn@gmail.com",
	  "msg_type":"text",
	  "content":{
	       "text":"text content test"
	   }}`
	email = "rammiahcn@gmail.com"
	const url = "https://open.feishu.cn/open-apis/message/v4/send/"
	clnt := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte(s)))
	if err != nil {
		panic(err)
	}
	token, err := b.getToken()
	if err != nil {
		panic(err)
	}
	fmt.Println("token ", token)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := clnt.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
	return nil
}
