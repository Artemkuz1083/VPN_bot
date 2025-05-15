package handlers

import (
	getreqforpanel "bot/GetReqForPanel"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
)

func HandleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	callback := update.CallbackQuery

	// Проверка на nil
	if callback == nil || callback.Message == nil {
		log.Println("Ошибка: callback или callback.Message равны nil")
		return
	}

	var responseMessage string

	//Говнокод
	switch callback.Data {
	case "get_key":
		responseMessage = "Вы выбрали 'Получить ключ'. Пожалуйста, следуйте инструкциям для получения вашего ключа."
	case "extend_subscription":
		responseMessage = "Варианты подписок"
	case "month":
		responseMessage = "Подписка на месяц оформлена"
		getreqforpanel.AddNewUser(update.Message.From.FirstName, 0, 0, 0, true, update.Message.From.ID, true)
	case "3_months":
		responseMessage = "Подписка на 3 месяца оформлена"
		getreqforpanel.AddNewUser(update.Message.From.FirstName, 0, 0, 0, true, update.Message.From.ID, true)
	case "6_months":
		responseMessage = "Подписка на полгода оформлена"
		getreqforpanel.AddNewUser(update.Message.From.FirstName, 0, 0, 0, true, update.Message.From.ID, true)
	case "year":
		responseMessage = "Подписка на год оформлена"
		getreqforpanel.AddNewUser(update.Message.From.FirstName, 0, 0, 0, true, update.Message.From.ID, true)
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
