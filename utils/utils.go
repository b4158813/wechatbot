package utils

func get_functions_list() string {
	var res string
	res += "【欢迎使用wlx专属微信聊天机器人】\n"
	res += "目前所支持的功能有：\n\n"
	res += "输入任何内容即可与chatgpt聊天（支持群聊@回复 + 私聊回复）\n"
	res += "输入 list:：展示此菜单\n"
	res += "输入 commemoration：显示纪念日信息\n"
	return res
}
