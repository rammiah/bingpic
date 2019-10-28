package lark

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestBot_NotifyImage(t *testing.T) {

}

func TestUploadImage(t *testing.T) {
	file, err := os.Open("th.jpeg")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf, _ := ioutil.ReadAll(file)
	log.Println(len(buf))
	key, err := testBot.uploadPictureToServer(buf)
	if err != nil {
		panic(err)
	}
	log.Printf("upload image success, key: %s", key)
}
