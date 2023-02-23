package handlers

import (
	"log"
	"strings"
	"wechatbot/gpt"
	"wechatbot/utils"

	"github.com/eatmoreapple/openwechat"
)

var _ MessageHandlerInterface = (*GroupMessageHandler)(nil)

// GroupMessageHandler 群消息处理
type GroupMessageHandler struct {
}

// handle 处理消息
func (g *GroupMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() {
		return g.ReplyText(msg)
	}
	return nil
}

// NewGroupMessageHandler 创建群消息处理器
func NewGroupMessageHandler() MessageHandlerInterface {
	return &GroupMessageHandler{}
}

// ReplyText 发送文本消息到群
func (g *GroupMessageHandler) ReplyText(msg *openwechat.Message) error {
	var err error
	var reply string

	// 接收群消息
	sender, _ := msg.Sender()
	group := openwechat.Group{sender}
	log.Printf("Received Group %v Text Msg : %v", group.NickName, msg.Content)

	// 不是@的不处理
	if !msg.IsAt() {
		return nil
	}

	// 替换掉@文本
	replaceText := "@" + sender.NickName
	requestText := strings.TrimSpace(strings.ReplaceAll(msg.Content, replaceText, ""))

	// 获取@我的用户
	groupSender, err := msg.SenderInGroup()
	if err != nil {
		log.Printf("get sender in group error :%v \n", err)
		return err
	}

	if strings.Count(requestText, "list") > 0 { // 获取功能列表
		reply, _ = utils.GetFunctionsList()
	} else if strings.Count(requestText, "memo") > 0 { // 获取纪念日信息
		reply, _ = utils.GetMemoDataInfo()
	} else { // chatgpt聊天功能
		// 向GPT发起请求
		reply, err = gpt.Completions(requestText)

		if err != nil {
			log.Printf("gpt request error: %v \n", err)
			msg.ReplyText("机器人神了，我一会发现了就去修。")
			return err
		}

		if reply == "" {
			return nil
		}

	}

	// 回复@我的用户
	reply = strings.TrimSpace(reply)
	reply = strings.Trim(reply, "\n")
	atText := "@" + groupSender.NickName
	replyText := atText + "\n" + reply
	_, err = msg.ReplyText(replyText)

	if err != nil {
		log.Printf("response group error: %v \n", err)
	}

	return err
}
