package adaptor

import (
	"trx_tracker/adaptor/utils"

	"gopkg.in/telebot.v3"
)

type TelegramAdaptor struct {
	bot *telebot.Bot
}

func NewTelegramBotAdaptor(bot *telebot.Bot) TelegramBot {
	return &TelegramAdaptor{
		bot: bot,
	}
}

func (b *TelegramAdaptor) SendMessage(message string) {

	value := utils.Recipient{
		ChatId: "7040234590",
	}
	b.bot.Send(&value, message)
}
