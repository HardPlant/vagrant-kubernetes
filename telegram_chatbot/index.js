const TelegramBot = require('node-telegram-bot-api');
const yaml = require('js-yaml');
const fs = require('fs');

var config = fs.readFileSync('config.yaml');
token = yaml.load(config)["token"];

const bot = new TelegramBot(token, {polling: true});

bot.on('message', (msg) => {
    bot.sendMessage(msg.chat.id, "I'll have the tuna?");
});