package lark

import (
	"log"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	appID := "cli_9d168c288fb29107"
	appSecret := "2NGzXhlFGSzzyfqU2Wumyd104Gfi4K8A"
	token, err := GetAccessToken(appID, appSecret)
	log.Println(token, err)
}
