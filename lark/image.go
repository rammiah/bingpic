package lark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"mime/multipart"
)

type uploadPictureResponse struct {
	Code int `json:"code"`
	Data struct {
		ImageKey string `json:"image_key"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func (b *Bot) uploadPictureToServer(img []byte) (string, error) {
	const uploadUrl = "https://open.feishu.cn/open-apis/image/v4/put/"
	// 创建一个form
	body := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", "img.jpeg")
	if err != nil {
		return "", err
	}
	// 写入图片
	_, err = part.Write(img)
	//log.Println(n)
	if err != nil {
		return "", err
	}
	// 图片类型
	err = writer.WriteField("image_type", "message")
	if err != nil {
		return "", err
	}
	// 关闭writer，不关闭会得到unexpected EOF错误
	err = writer.Close()
	if err != nil {
		return "", err
	}
	// 得到头部
	token, err := b.getToken()
	if err != nil {
		return "", err
	}
	header := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  writer.FormDataContentType(),
	}
	// 发送请求
	resp, err := b.client.PostWithHeader(uploadUrl, header, body.Bytes())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// 解析判断返回
	var picResp uploadPictureResponse
	err = json.Unmarshal(content, &picResp)
	if err != nil {
		return "", err
	}
	if picResp.Code != successCode {
		return "", errors.Errorf("%d %s: %s", picResp.Code, picResp.Msg, MapError(picResp.Code))
	}

	return picResp.Data.ImageKey, nil
}

func (b *Bot) NotifyImage(email string, img []byte) error {
	const sendUrl = "https://open.feishu.cn/open-apis/message/v4/send/"
	imgKey, err := b.uploadPictureToServer(img)
	if err != nil {
		return err
	}
	req := messageRequest{
		Email:   email,
		MsgType: "image",
		Content: map[string]string{"image_key": imgKey},
	}
	// 发送图片消息
	buf, err := json.Marshal(req)
	if err != nil {
		return err
	}
	// 得到头部
	token, err := b.getToken()
	if err != nil {
		return err
	}
	header := map[string]string{
		"Authorization": "Bearer " + token,
	}
	resp, err := b.client.PostWithHeader(sendUrl, header, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var msgResp messageResponse
	err = json.Unmarshal(buf, &msgResp)
	if err != nil {
		return err
	}
	// 判断结果
	if msgResp.Code != successCode {
		return fmt.Errorf("%d %s: %s", msgResp.Code, msgResp.Msg, MapError(msgResp.Code))
	}

	return nil
}
