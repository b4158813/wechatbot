package utils

import (
	"fmt"
	"log"
	"strings"
	"time"
	"wechatbot/config"

	"github.com/eatmoreapple/openwechat"
	"github.com/robfig/cron/v3"
)

var bot *openwechat.Bot

func StartTimerTask(bot_param *openwechat.Bot) {

	bot = bot_param

	c := cron.New(cron.WithSeconds())
	c.Start()
	// second minute hour day month week
	c.AddFunc("0 0 19 * * ?", QQReminderTask)   // qq互发消息提醒
	c.AddFunc("0 0 19 * * ?", MemoReminderTask) // 纪念日提醒

	select {}
}

func MemoReminderTask() {

	self, _ := bot.GetCurrentUser()
	groups, _ := self.Groups()
	// group := groups.GetByNickName("bot测试")
	group := groups.GetByNickName(config.Config.Extra.GroupName)

	memo_data := GetMemoData(config.Config.Extra.MemoDayFile)
	now_time := time.Now()
	desc := memo_data[0].description
	ymd := strings.Split(memo_data[0].ymd.String(), " ")[0]
	rest_day := int(memo_data[0].ymd.Sub(now_time).Hours() / 24)
	if rest_day == 1 || rest_day == 3 || rest_day == 7 {
		s := fmt.Sprintf("⏰叮~下面这个纪念日快到啦\n\n⭐%s\n%10s %5d天\n", desc, ymd, rest_day)
		_, err := group.SendText(s)
		if err != nil {
			log.Printf("MemoReminderTask group.SendText error")
			return
		}
		log.Printf("MemoReminderTask send ok.")
	}
}

func QQReminderTask() {

	self, _ := bot.GetCurrentUser()
	groups, _ := self.Groups()
	// group := groups.GetByNickName("bot测试")
	group := groups.GetByNickName(config.Config.Extra.GroupName)
	_, err := group.SendText("⏰ 记得QQ互发消息哦~")
	if err != nil {
		log.Printf("QQReminderTask group.SendText error")
		return
	}
	log.Printf("QQReminderTask send ok.")
}
