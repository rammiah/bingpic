package lark

import (
	"bingpic/lark/rttp"
	"time"
)

type Bot struct {
	appID           string
	appSecret       string
	token           string
	tokenExpireTime time.Time
	client          *rttp.Client
}

func NewBot(appID string, appSecret string) (*Bot, error) {
	token, err := GetAccessToken(appID, appSecret)
	if err != nil {
		return nil, err
	}
	bot := &Bot{appID: appID,
		appSecret:       appSecret,
		token:           token,
		tokenExpireTime: time.Now().Add(2 * time.Hour),
		client:          rttp.NewClient()}
	return bot, nil
}

func (b *Bot) getToken() (string, error) {
	if time.Now().After(b.tokenExpireTime) {
		// 获取一下token
		token, err := GetAccessToken(b.appID, b.appSecret)
		if err != nil {
			return "", err
		}
		b.token = token
		b.tokenExpireTime = time.Now().Add(2 * time.Hour)
	}
	return b.token, nil
}
