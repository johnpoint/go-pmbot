package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	var adminID int64 = 1
	var botKey string = ""
	bot, err := tgbotapi.NewBotAPI(botKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			if update.Message.From.ID == int(adminID) {
				msg.Text = "欢迎使用PM-bot,当有人受限于帐号限制无法私聊时可以通过我转达~\n 选中要回复的消息使用 /reply + 回复内容 可以回复"
			} else {
				msg.Text = "欢迎使用PM-bot,受限于帐号限制无法私聊时可以通过我转达~\n /say 开始使用"
			}
		case "say":
			NewMsg := tgbotapi.NewMessage(adminID, "New message~")
			NewMsgText := tgbotapi.NewForward(adminID,
				update.Message.Chat.ID, update.Message.MessageID)
			bot.Send(NewMsg)
			bot.Send(NewMsgText)
			msg.Text = "sended!"
		case "status":
			msg.Text = "alive~"
		case "reply":
			NewMsg := tgbotapi.NewForward(int64(update.Message.ReplyToMessage.ForwardFrom.ID),
				update.Message.Chat.ID, update.Message.MessageID)
			bot.Send(NewMsg)
			msg.Text = "done!"
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
