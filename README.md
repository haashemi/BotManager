# BotManagerBot

A Simple bot manager for securing local hosted telegram-bot-api and avoiding Linux file permissions, nginx configurations, etc.

The bot is based on the [tgo](https://github.com/haashemi/tgo) framework, and the API is based on the [gorilla mux](https://github.com/gorilla/mux) router.

### Bot Commands:

- `/add` Whitelist a new bot
- `/rem` Remove a bot from the whitelist
- `/bots` List of whitelisted bots

### API Endpoints:

There are only two endpoints, which are 1:1 to the official Telegram Bot API.

1. `/{token}/{method}`
2. `/file/{token}/{dir}/{file}` (`{dir}/{file}` is equal to `{file_path}`)

### Usage:

1. Modify the config file however you prefer

```bash
$ cp config.example.yaml config.yaml
$ # Modify the config.yaml however you want here
```

2. Compile and run it!

```
$ go build .
$ ./BotManagerBot
```
