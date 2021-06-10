# Scada-LTS Telegram Bot Service

This service aims to bridge EventHandlers, HTTP Publishers in Scada-LTS, to telegram channels by the means of a separate service with an API and a telegram bot implementatiom

## Getting started


## Commands

- mute
- unmute
- subscribe

## Setting up Bot


## Testing locally

```
curl -XPOST -d 'url=<ngrok_endpoint>/webhook' 'https://api.telegram.org/bot<TOKEN>/SetWebhook'
```

## References
https://medium.com/swlh/build-a-telegram-bot-in-go-in-9-minutes-e06ad38acef1
https://github.com/profclems/go-dotenv
ngrok
https://github.com/kimrgrey/go-telegram