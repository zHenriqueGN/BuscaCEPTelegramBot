# BuscaCEPTelegramBot
A Telegram bot to get address info by CEP

## Running the bot:
### Configuring the environment
- create a .env file and set the variable BOTAPITOKEN with your bot token from [BotFather](https://t.me/botfather)

### Running the container
- build the image
  - "docker build -t busca-cep-bot:1.0 ."
- run the container
  - "docker container run busca-cep-bot:1.0"

### Accessing the bot
- After the previous commands, you can just run an "/start" in your telegram bot to start it.
