package lark

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

func (b *Bot) UploadPicture(img []byte) (string, error) {
	const url = "https://open.feishu.cn/open-apis/image/v4/upload/"



	return "", nil
}
