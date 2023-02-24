package utils

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/eatmoreapple/openwechat"
)

func StartTimerTask(bot *openwechat.Bot) {
	ticker := time.NewTicker(time.Duration(24) * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			go MemoReminderTask(bot)
			go QQReminderTask(bot)
		}
	}
}

func MemoReminderTask(bot *openwechat.Bot) error {

	self, _ := bot.GetCurrentUser()
	groups, _ := self.Groups()
	// group := groups.GetByNickName("bot测试")
	group := groups.GetByNickName("🥰")

	memo_data := GetMemoData("memo_day.txt")
	now_time := time.Now()
	desc := memo_data[0].description
	ymd := strings.Split(memo_data[0].ymd.String(), " ")[0]
	rest_day := int(memo_data[0].ymd.Sub(now_time).Hours() / 24)
	if rest_day == 1 || rest_day == 3 || rest_day == 7 {
		s := fmt.Sprintf("⏰叮~下面这个纪念日快到啦\n\n⭐%s\n%10s %5d天\n", desc, ymd, rest_day)
		_, err := group.SendText(s)
		if err != nil {
			log.Printf("group.SendText()")
			return err
		}
	}
	return nil
}

func QQReminderTask(bot *openwechat.Bot) error {

	self, _ := bot.GetCurrentUser()
	groups, _ := self.Groups()
	// group := groups.GetByNickName("bot测试")
	group := groups.GetByNickName("🥰")
	_, err := group.SendText("⏰ 记得QQ互发消息哦~")
	if err != nil {
		log.Printf("group.SendText()")
		return err
	}
	return nil
}
