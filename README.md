# wechatbot_wlx
本项目魔改自[djun/wechatbot_wlx](https://github.com/djun/wechatbot)

项目基于 [openwechat](https://github.com/eatmoreapple/openwechat) 开发，目前实现了以下功能：
 + chatgpt-api的 群聊@回复
 + chatgpt-api的 私聊回复
 + 添加好友自动通过并回复

npy专属功能：
 + list: 列出功能菜单
 + memo: 显示出最近的几个纪念日
 + 每日提醒 qq互发消息
 + 离最近的纪念日仅剩7/3/1天时向专属群聊发送提醒消息

TODO:
 + 设置daemon模式
 + 后台日志输出到文件中
 + 将群聊名称等信息加入到config环境变量中
 + 部署到云服务器上 (done)

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

````