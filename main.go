package main

import (
	database "bot/DataBase"
	getreqforpanel "bot/GetReqForPanel"
	botconn "bot/botConn"
	"bot/handlers"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Создаем соединение с ботом
	bot := botconn.GetBotConnection()

	var userId int

	updates := botconn.GetUpdates()

	// Обрабатываем обновления
	for update := range updates {
		if update.CallbackQuery != nil {
			fmt.Println(userId)
			HandleCallback(bot, update, userId)
		} else if update.Message != nil {
			userId = update.Message.From.ID
			// Обработчик сообщений
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					if database.CheckUserExists(userId) {
						handlers.StartMenuForNewUsers(bot, update.Message.Chat.ID)
					} else {
						getreqforpanel.AddNewUser(update.Message.From.FirstName, 0, 0, 0, true, userId, false)
						handlers.StartMenuForNewUsers(bot, update.Message.Chat.ID)
					}
				}

			}
		}
	}
}

func HandleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, userId int) {
	callback := update.CallbackQuery

	var responseMessage string

	switch callback.Data {
	case "get_key":
		status, _ := database.StatusPayment(userId)
		if status {
			responseMessage = fmt.Sprintf("Ваш ключ: http://212.113.116.19:2096/sub/%d", userId)
		} else {
			responseMessage = "Сначала оплатите подписку"
		}
	case "extend_subscription":
		responseMessage = "Варианты подписок"
	case "month":
		stat, _ := database.UpdateStatusPayment(userId)
		if stat {
			responseMessage = "Подписка оформлена"
		} else {
			responseMessage = "Подписка не оформлена"
		}
	case "3_months":
		stat, _ := database.UpdateStatusPayment(userId)
		if stat {
			responseMessage = "Подписка оформлена"
		} else {
			responseMessage = "Подписка не оформлена"
		}
	case "6_months":
		stat, _ := database.UpdateStatusPayment(userId)
		if stat {
			responseMessage = "Подписка оформлена"
		} else {
			responseMessage = "Подписка не оформлена"
		}
	case "year":
		stat, _ := database.UpdateStatusPayment(userId)
		if stat {
			responseMessage = "Подписка оформлена"
		} else {
			responseMessage = "Подписка не оформлена"
		}
	default:
		responseMessage = "Неизвестная команда."
	}

	if callback.Data != "extend_subscription" {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, responseMessage)
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Ошибка при отправке сообщения: %v", err)
		}

		// Удаляем уведомление о нажатии кнопки
		if _, err := bot.AnswerCallbackQuery(tgbotapi.NewCallback(callback.ID, "")); err != nil {
			log.Printf("Ошибка при ответе на callback: %v", err)
		}
	} else {
		// Создаем кнопки для расширения подписки
		row1 := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Месяц", "month"),
		)
		row2 := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("3 Месяца", "3_months"),
		)
		row3 := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Полгода", "6_months"),
		)
		row4 := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Год", "year"),
		)

		keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4)
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, responseMessage)
		msg.ReplyMarkup = &keyboard

		if _, err := bot.Send(msg); err != nil {
			log.Printf("Ошибка при отправке сообщения с кнопками: %v", err)
		}
	}
}
