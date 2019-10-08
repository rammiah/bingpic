package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDailyPictureDao_InsertDailyPicture(t *testing.T) {
	pic := DailyPicture{
		Date: time.Now(),
		Url:  "https:/cn.bing.com/th?id=OHR.WorldOctopus_ZH-CN2670477302_UHD.jpg&rf=LaDigue_UHD.jpg&pid=hp&w=1920&h=1080&rs=1&c=4",
	}
	err := NewDailyPictureDaoInstance().InsertDailyPicture(pic)
	assert.Equal(t, nil, err, "insert error: %v", err)
}

func TestDailyPictureDao_QueryDailyPictureByDate(t *testing.T) {
	pic, err := NewDailyPictureDaoInstance().QueryDailyPictureByDate(time.Now())
	assert.Equal(t, nil, err, "get picture error: %v", err)
	assert.NotEqual(t, nil, pic, "picture should not nil")
}
