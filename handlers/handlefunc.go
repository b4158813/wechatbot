package handlers

import (
	"strings"
	"wechatbot/config"
	"wechatbot/gpt"
	"wechatbot/utils"
)

func HandleRequestText(requestText string) (string, error) {
	var reply string
	var err error
	if strings.Count(requestText, "list") > 0 { // 获取功能列表
		reply, err = utils.GetFunctionsList()
		if err != nil {
			return "", err
		}
	} else if strings.Count(requestText, "memo") > 0 { // 获取纪念日信息
		reply, err = utils.GetMemoDataInfo()
		if err != nil {
			return "", err
		}
	} else if config.Config.GptChat { // chatgpt聊天功能
		// 向GPT发起请求
		reply, err = gpt.Completions(requestText)
		if err != nil {
			return "", err
		}
	} else {
		reply = "未开启gpt聊天功能！"
	}

	return reply, nil
}
