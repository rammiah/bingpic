package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDailyPictureDao_InsertDailyPicture(t *testing.T) {
	pic := &DailyPicture{
		Date: time.Now(),
		Url:  "https:/cn.bing.com/th?id=OHR.BubbleNebula_ZH-CN2787112807_1920x1080.jpg&rf=LaDigue_1920x1080.jpg&pid=hp",
	}
	err := NewDailyPictureDaoInstance().InsertDailyPicture(pic)
	assert.Equal(t, nil, err, "insert error: %v", err)
}

func TestDailyPictureDao_QueryDailyPictureByDate(t *testing.T) {
	pic, err := NewDailyPictureDaoInstance().QueryDailyPictureByDate(time.Now())
	assert.Equal(t, nil, err, "get picture error: %v", err)
	assert.NotEqual(t, nil, pic, "picture should not nil")
}
