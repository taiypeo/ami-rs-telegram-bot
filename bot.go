package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panicf("Bot API creation error: %s", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	rsCounts := make(map[string]int, 0)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if handler, ok := COMMANDS[update.Message.Command()]; ok {
				msg := handler(update.Message.Chat.ID, rsCounts)
				if _, err := bot.Send(msg); err != nil {
					log.Printf("Failed to send a message to %d: %s", update.Message.Chat.ID, err)
				}

				continue
			}
		}

		if checkRsMessage(update.Message.Text) {
			rsCounts[update.Message.From.UserName] += 1

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, formatReply(rsCounts))
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Failed to send a message to %d: %s", update.Message.Chat.ID, err)
			}
		}
	}
}
