package picgetter

import (
	"log"
	"testing"
)

func TestGetTodayPicture(t *testing.T) {
	pic, _ := GetTodayPicture()
	log.Println("url", pic.Url)
}
