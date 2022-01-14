package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type commandFunc func(chatID int64, misc ...interface{}) tgbotapi.MessageConfig

var COMMANDS = map[string]commandFunc{
	"clear": handleClear,
	"help":  handleHelp,
	"stats": handleStats,
}

func handleClear(chatID int64, misc ...interface{}) tgbotapi.MessageConfig {
	rsCounts := misc[0].(map[string]int)
	for k := range rsCounts {
		delete(rsCounts, k)
	}

	return tgbotapi.NewMessage(chatID, "Статистика по РС очищена!")
}

func handleHelp(chatID int64, misc ...interface{}) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(
		chatID,
		"Бот для подсчета сообщений людей, которые лучше бы пошли на РС.\n"+
			"/stats -- показывает текущую глобальную статистику по РС\n"+
			"/clear -- очищает глобальную статистику по РС\n"+
			"/help -- показывает это сообщение",
	)
}

func handleStats(chatID int64, misc ...interface{}) tgbotapi.MessageConfig {
	rsCounts := misc[0].(map[string]int)
	return tgbotapi.NewMessage(chatID, formatReply(rsCounts))
}
