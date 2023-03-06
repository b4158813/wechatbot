package handlers

import (
	"log"
	"strings"
	"wechatbot/gpt"
	"wechatbot/utils"

	"github.com/eatmoreapple/openwechat"
)

var _ MessageHandlerInterface = (*UserMessageHandler)(nil)

// UserMessageHandler 私聊消息处理
type UserMessageHandler struct {
}

// handle 处理消息
func (g *UserMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() {
		return g.ReplyText(msg)
	}
	return nil
}

// NewUserMessageHandler 创建私聊处理器
func NewUserMessageHandler() MessageHandlerInterface {
	return &UserMessageHandler{}
}

// ReplyText 发送文本消息给朋友
func (g *UserMessageHandler) ReplyText(msg *openwechat.Message) error {
	var err error
	var reply string

	// 接收私聊消息
	sender, _ := msg.Sender()
	log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)

	requestText := strings.Trim(msg.Content, "\n")

	if "list" == requestText { // 获取功能列表
		reply, _ = utils.GetFunctionsList()
	} else if "memo" == requestText { // 获取纪念日信息
		log.Printf(sender.UserName)
		reply, _ = utils.GetMemoDataInfo()
	} else { // chatgpt聊天功能
		// 向GPT发起请求
		reply, err = gpt.CompletionsNew(requestText)

		if err != nil {
			log.Printf("gpt.CompletionsNew error: %v \n", err)
			msg.ReplyText("机器人神了，我一会发现了就去修。")
			return err
		}

		if reply == "" {
			return nil
		}

	}

	// 回复用户
	reply = strings.TrimSpace(reply)
	reply = strings.Trim(reply, "\n")
	_, err = msg.ReplyText(reply)
	if err != nil {
		log.Printf("response user error: %v \n", err)
	}
	return err
}
