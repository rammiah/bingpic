package main

import (
	"bingpic/lark"
	"bingpic/model"
	"bingpic/picgetter"
	log "github.com/sirupsen/logrus"
)

const (
	appId     = "cli_9d168c288fb29107"
	appSecret = "2NGzXhlFGSzzyfqU2Wumyd104Gfi4K8A"
)

func main() {
	const email = "rammiahcn@gmail.com"
	// 使用执行一次的形式，让crontab每天调用吧
	bot, err := lark.NewBot(appId, appSecret)
	if err != nil {
		panic(err)
	}
	pic, err := picgetter.GetTodayPicture()
	if err != nil {
		bot.NotifyText(email, "get picture failed, error: "+err.Error())
		return
	}
	err = model.NewDailyPictureDaoInstance().InsertDailyPicture(pic)
	if err != nil {
		log.Errorf("insert into db failed, error: %v", err)
	}
	// 发送图片
	img, err := Download(pic.Url)
	if err != nil {
		bot.NotifyText(email, "download image failed, error: "+err.Error())
		return
	}
	err = bot.NotifyImage("rammiahcn@gmail.com", img)
	if err != nil {
		bot.NotifyText(email, "notify text failed, error: "+err.Error())
	}
}
