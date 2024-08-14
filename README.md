
Annona bot

处理 keyworld 请求
反馈 keyworld msg
上报授权请求
查询 user info 信息


## 环境设置
```env
# bot webhook url
BOT_TELEGRAM_WEBHOOK_URL=https://xxx/webhook/tele
```

## 设置 webhook

post  https://xxx/webhook/tele/botToken

or

get/post  https://xxx/webhook/set/botToken

```
curl -d '' -X POST  https://xxx/webhook/set/botToken
```

## 启用机器人，同步机器人信息到core
```
/botenable
```