# wechatbot_wlx
本项目魔改自[djun/wechatbot_wlx](https://github.com/djun/wechatbot_wlx)

项目基于[openwechat](https://github.com/eatmoreapple/openwechat)
开发，目前实现了以下功能
 + 群聊@回复
 + 私聊回复
 + 添加好友自动通过并回复
 
# 注册openai
chatGPT注册可以参考[这里](https://juejin.cn/post/7173447848292253704)

# 安装使用
````
# 获取项目
git clone https://github.com/b4158813/wechatbot_wlx

# 进入项目目录
cd wechatbot

# 复制配置文件
copy config.dev.json config.json

# 启动项目
go run main.go

启动前需替换config中的api_key
