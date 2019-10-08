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
	// 使用执行一次的形式，让crontab每天调用吧
	pic, err := picgetter.GetTodayPicture()
	if err != nil {
		// TODO: lark报警
		return
	}
	err = model.NewDailyPictureDaoInstance().InsertDailyPicture(pic)
	if err != nil {
		log.Errorf("insert into db failed, error: %v", err)
	}
	// end
	// TODO: lark提示
	//bot := lark.NewBot(appId, appSecret)
	err = bot.NotifyImage("")
	if err != nil {
		panic(err)
	}
}
