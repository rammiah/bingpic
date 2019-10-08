package model

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type DailyPicture struct {
	Date time.Time `gorm:"column:date" json:"date"`
	Url  string    `gorm:"column:url" json:"url"`
}

func (*DailyPicture) TableName() string {
	return "daily_picture"
}

type DailyPictureDao struct {
}

func NewDailyPictureDaoInstance() *DailyPictureDao {
	// 没有内部变量的话返回空指针也一样
	return nil
}

func (*DailyPictureDao) InsertDailyPicture(pic *DailyPicture) error {
	err := db.Save(pic).Error
	if err != nil {
		log.Errorf("save picture into db failed, error: %v", err)
		return err
	}
	return nil
}

func (*DailyPictureDao) QueryDailyPictureByDate(date time.Time) (*DailyPicture, error) {
	var pic DailyPicture
	err := db.Where("date = ?", date.Format("2006-01-02")).First(&pic).Error
	if err != nil {
		log.Errorf("get picture from db failed, error: %v", err)
		return nil, err
	}

	return &pic, nil
}
