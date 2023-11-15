# BotManager

A Simple bot manager for securing local hosted telegram-bot-api and avoiding Linux file permissions, nginx configurations, etc.

The bot is based on the [tgo](https://github.com/haashemi/tgo) framework, and the API is based on the [gorilla mux](https://github.com/gorilla/mux) router.

### Bot Commands:

- `/add` Whitelist a new bot
- `/rem` Remove a bot from the whitelist
- `/bots` List of whitelisted bots

### API Endpoints:

There are only two endpoints, which are 1:1 to the official Telegram Bot API.

- `/{token}/{method}`
- `/file/{token}/{dir}/{file}`
  - (`{dir}/{file}` is equal to `{file_path}`)

### Usage:

1. Modify the config file however you prefer

```bash
$ cp config.example.yaml config.yaml
$ # Modify the config.yaml however you want here
```

2. Compile and run it!

```
$ go build .
$ ./BotManager
```

### Note:

I don't really recommend using this project as it's made because of my laziness to deal with Linux file permissions, Nginx configurations, and dynamic bot whitelisting.
If you can deal with the stuff I said in a better way, then for sure do it in your own way. But if you are lazy as me, just use it and worry about nothing.
