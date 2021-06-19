# Telegram Bot Notification Service

This service aims to bridge any issuer to telegram channels by the means of a separate service with an API and a telegram bot implementation.

This service was initially invisioned to work as a notification service for ScadaLTS[github.com/katesclau]

![Architecture](/assets/Telegram_Notifier.png)

## Getting started
<!-- TODO: improve  -->
Copy `.env.example` to `.env` - edit with your bot's token
```
# Telegram access
TELEGRAM_TOKEN=bot_token
TELEGRAM_BOT_NAME=bot_name
```

```
git clone https://github.com/katesclau/telegram-notfication-service
cd telegram-notfication-service
make
```

In another terminal, start ngrok
```
ngrok http 8088
```

Issue the endpoint to Telegram API to attach your webhook to the bot's configuration
```
curl -XPOST -d 'url=<ngrok_https_endpoint>/webhook' 'https://api.telegram.org/bot<TELEGRAM_TOKEN>/SetWebhook'
```

## Bot Commands

Main goal of this bot handler is to enable users to subscribe to topics configured here.

- mute - Disable notifcations for subscribed topics
- unmute - Enable notifications for subscribed topics
- subscribe - Subscribe chat to a named topic
- topics - List topics available in this bot

This service uses a similar implementation to [2].

## API
The REST API implements
- Topic
- Subscribers 

### /topic
Allows 
- GET: retrieve the existing topics
- POST: create a new topic

### /topic/<topic_name>
- GET: retrieve information on the topic
- DELETE: removes topic

### /topic/<topic_name>/event
- POST: push event to subscribers

### /topic/<topic_name>/subscribers
- GET: retrieve subscribers information, for a given topic

## API Authentication
Consists in a single Bearer token, defined in the .env file. Please provide it in each request as such...

```
curl -XGET -H 'Authorization: Bearer <API_TOKEN defined in .env>' -H "Content-type: application/json" 'http://localhost:8088/topic/some_topic/subscribers'
```

## References
1. https://medium.com/swlh/build-a-telegram-bot-in-go-in-9-minutes-e06ad38acef1
2. https://github.com/profclems/go-dotenv
3. ngrok
4. https://github.com/kimrgrey/go-telegram
5. https://gorm.io
6. https://medium.com/rate-engineering/go-test-your-code-an-introduction-to-effective-testing-in-go-6e4f66f2c259
7. https://github.com/stretchr/testify
8. https://www.rodrigoaraujo.me/posts/golang-pattern-graceful-shutdown-of-concurrent-events/
9. https://www.callicoder.com/deploy-containerized-go-app-kubernetes/

## TODO
- Telegram commands
- Test Routes
- Helm package
- Docs
