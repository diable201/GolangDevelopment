package main

import (
	"github.com/diable201/GolangDevelopment/tree/master/HW_04/weather"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("YOUR-KEY")
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

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "type /sayhi or /status."
			case "weather":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					weather.GetWeather(strings.ToLower(update.Message.CommandArguments())))
				bot.Send(msg)
			case "time":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					weather.GetTime(strings.ToLower(update.Message.CommandArguments())))
				bot.Send(msg)
			case "json":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					weather.SerializeWeather(strings.ToLower(update.Message.CommandArguments())))
				bot.Send(msg)
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}
	}
}
