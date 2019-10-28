package lark

var (
	testBot *Bot
)

func init() {
	var err error
	const (
		appId     = "cli_9d168c288fb29107"
		appSecret = "2NGzXhlFGSzzyfqU2Wumyd104Gfi4K8A"
	)
	testBot, err = NewBot(appId, appSecret)
	if err != nil {
		panic(err)
	}
}
