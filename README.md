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
 + 添加用户自主设定定时任务的功能
   - 例如：用户输入 “提醒我 该换隐形眼镜了 25d0h0m” 代表25天后微信私聊/群聊@提醒 该换眼镜了 这件事


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

# 启动项目 (前台模式)
go run main.go

# 启动项目 (daemon模式，日志打印到./wechatbot.log)
./run.sh

启动前需替换config中的api_key
有额外信息请自行更改extra结构体、config文件以及对应代码

````