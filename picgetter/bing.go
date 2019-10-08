package picgetter

import (
	"bingpic/model"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func GetTodayPicture() (*model.DailyPicture, error) {
	// 必应的网址
	const url = "https://cn.bing.com/"
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("get response from %s failed, error: %v", url, err)
		return nil, err
	}
	defer resp.Body.Close()
	// 解析url
	imgUrl, err := parseBingUrl(resp.Body)
	if err != nil {
		log.Errorf("parse url from body failed, error: %v", err)
		return nil, err
	}
	// 把前面的主机名加上吧
	imgUrl = resp.Request.URL.String() + imgUrl
	return &model.DailyPicture{
		Date: time.Now(),
		Url:  imgUrl,
	}, nil
}

func parseBingUrl(reader io.Reader) (string, error) {
	// goquery库，还行吧
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}
	/// 一条语句直接得到url
	url, ok := doc.Find("#bgLink").First().Attr("href")
	if !ok {
		return "", nil
	}
	log.Infof("parse url result: %v", url)

	return url, nil
}
