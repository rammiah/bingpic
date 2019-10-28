package lark

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
)

type uploadPictureRequest struct {
	Image []byte `json:"image"`
}

type uploadPictureResponse struct {
	Code int
	Data struct {
		ImageKey string `json:"image_key"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func (b *Bot) uploadPictureToServer(img []byte) (string, error) {
	const url = "https://open.feishu.cn/open-apis/image/v4/upload/"
	token, err := b.getToken()
	if err != nil {
		return "", err
	}
	header := map[string]string{
		"Authorization": "Bearer " + token,
	}
	buf := bytes.NewBuffer(nil)
	buf.WriteString(`{"image":`)
	buf.Write(img)
	buf.WriteString(`}`)
	resp, err := b.client.PostWithHeader(url, header, buf.Bytes())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var picResp uploadPictureResponse
	err = json.Unmarshal(content, &picResp)
	if err != nil {
		return "", err
	}
	if picResp.Code != successCode {
		return "", errors.Errorf("%d: %s", picResp.Code, MapError(picResp.Code))
	}

	return picResp.Data.ImageKey, nil
}

func (b *Bot) NotifyImage(email string, img []byte) error {

	return nil
}
