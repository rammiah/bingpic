package lark

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func getImage() ([]byte, error) {
	file, err := os.Open("th.jpeg")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf, _ := ioutil.ReadAll(file)
	log.Println(len(buf))
	return buf, nil
}

func TestBot_NotifyImage(t *testing.T) {
	img, err := getImage()
	assert.Equal(t, nil, err)
	err = testBot.NotifyImage("rammiahcn@gmail.com", img)
	assert.Equal(t, nil, err)
}

func TestUploadImage(t *testing.T) {
	img, err := getImage()
	assert.Equal(t, nil, err)
	key, err := testBot.uploadPictureToServer(img)
	assert.Equal(t, nil, err)
	log.Printf("upload image success, key: %s", key)
}
