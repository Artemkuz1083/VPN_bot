package main

import (
	database "bot/DataBase"
	getreqforpanel "bot/GetReqForPanel"
	botconn "bot/botConn"
	"bot/handlers"

	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Создаем соединение с ботом
	bot := botconn.GetBotConnection()

	updates := botconn.GetUpdates()

	// Обрабатываем обновления
	for update := range updates {
		if update.CallbackQuery != nil {
			handlers.HandleCallback(bot, update)
		} else if update.Message != nil {
			// Обработчик сообщений
			if update.Message.IsCommand() {
				userId := update.Message.From.ID
				switch update.Message.Command() {
				case "start":
					if database.CheckUserExists(userId) {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет старый юзер!")
						if _, err := bot.Send(msg); err != nil {
							log.Printf("Ошибка при отправке сообщения: %v", err)
						}
					} else {
						handlers.StartMenuForNewUsers(bot, update.Message.Chat.ID)
					}
				case "132":
					getreqforpanel.Authenticate()
				case "sub":
					getreqforpanel.Authenticate()
				}
			}

		}
	}
}
