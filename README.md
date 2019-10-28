# 必应每日图片项目

## 功能

- 每日图片url获取，持久化到数据库
- 获取失败lark机器人提示
- 获取成功发送图片到lark进行提示

## 进度
- [x] url获取与持久化
- [x] 失败时的消息提醒
- [x] 成功时的图片发送

## 相关库

- [goquery](https://github.com/PuerkitoBio/goquery)：html解析库
- [gorm](https://github.com/jinzhu/gorm)：ORM库
- [logrus](https://github.com/sirupsen/logrus)：日志打印相关
- [testify](https://github.com/stretchr/testify)：测试相关，断言使用
- [mysql](https://github.com/go-sql-driver/mysql)：数据库驱动
