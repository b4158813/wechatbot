package utils

import (
	"fmt"
	"strings"
	"time"
)

// 打印功能菜单
func GetFunctionsList() (string, error) {
	var res string
	res += "【欢迎使用wlx专属bot】\n"
	res += "目前所支持的功能有：\n\n"
	res += "- 输入任何内容即可与chatgpt聊天（支持群聊@回复 + 私聊回复）\n\n"
	res += "- 输入 list：展示此菜单\n\n"
	res += "- 输入 memo：显示纪念日信息\n\n"
	return res, nil
}

// 打印纪念日信息
func GetMemoDataInfo() (string, error) {
	memo_data := GetMemoData("memo_day.txt")
	var res string
	now_time := time.Now()
	res += "☆最近的3个纪念日☆\n\n"
	for i := 0; i < 3; i++ {
		desc := memo_data[i].description
		ymd := strings.Split(memo_data[i].ymd.String(), " ")[0]
		rest_day := int(memo_data[i].ymd.Sub(now_time).Hours() / 24)
		res += fmt.Sprintf("⭐%s\n%10s %5d天\n\n", desc, ymd, rest_day)
	}
	return res, nil
}
