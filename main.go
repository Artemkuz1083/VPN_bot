package main

import (
	database "bot/DataBase"
	botconn "bot/botConn"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Создаем соединение с ботом
	bot := botconn.GetBotConnection()

	updates := botconn.GetUpdates()

	// Обрабатываем обновления
	for update := range updates {
		if update.Message == nil { // игнорируем не сообщения
			continue
		}

		// Простой ответ на команду /start
		if update.Message.IsCommand() {
			userId := update.Message.From.ID
			switch update.Message.Command() {
			case "start":
				if database.CheckUserExists(userId) {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет старый юзер!")
					bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет новый юзер!")
					bot.Send(msg)
				}
			case "132":
			}
		}
	}
}
