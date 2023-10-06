# BuscaCEPTelegramBot
A Telegram bot to get address info by CEP

## Running the bot:
### Initiating
- To run the bot, you will need variable a bot token from [BotFather](https://t.me/botfather)

### Running the container
- build the image
  - "docker build -t busca-cep-bot:1.0 ."
- run the container
  - "docker container run --env BOTAPITOKEN=YOUR_BOT_TOKEN henriquegn/busca-cep-bot:1.0"

### Accessing the bot
- After the previous commands, you can just run an "/start" in your telegram bot to start it.
