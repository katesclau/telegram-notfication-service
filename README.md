# Telegram Bot Service

This service aims to bridge any issuer to telegram channels by the means of a separate service with an API and a telegram bot implementation.

This service was initially invisioned to work as a notification service for ScadaLTS[scada-lts.org]

## Getting started


## Commands

Main goal of this bot handler is to enable users to subscribe to topics configured here.

- mute - Disable notifcations for subscribed topics
- unmute - Enable notifications for subscribed topics
- subscribe - Subscribe chat to a named topic
- list - List topics available in this bot

This service uses a similar implementation to [2].

## API


## Setting up Bot
Check 
## Testing locally
Local testing can be done by using ngrok to tunnel the local service to a valid https endpoint and pushing the created endpoint to Telegram's webhook configuration.

```
curl -XPOST -d 'url=<ngrok_endpoint>/webhook' 'https://api.telegram.org/bot<TOKEN>/SetWebhook'
```

## References
1. https://medium.com/swlh/build-a-telegram-bot-in-go-in-9-minutes-e06ad38acef1
2. https://github.com/profclems/go-dotenv
3. ngrok
4. https://github.com/kimrgrey/go-telegram