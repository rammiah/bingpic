package lark

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBot_NotifyText(t *testing.T) {
	const (
		appId     = "cli_9d168c288fb29107"
		appSecret = "2NGzXhlFGSzzyfqU2Wumyd104Gfi4K8A"
	)

	bot, err := NewBot(appId, appSecret)
	assert.Equal(t, nil, err)
	err = bot.NotifyText("", "hello world")
	assert.Equal(t, nil, err)
}
