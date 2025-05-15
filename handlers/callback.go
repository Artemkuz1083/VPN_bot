package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
)

func HandleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	callback := update.CallbackQuery
	var responseMessage string

	switch callback.Data {
	case "get_key":
		responseMessage = "Вы выбрали 'Получить ключ'. Пожалуйста, следуйте инструкциям для получения вашего ключа."
	case "extend_subscription":
		responseMessage = "Вы выбрали 'Продлить подписку'. Пожалуйста, следуйте инструкциям для продления вашей подписки."
	case "month":
		//TODO привязать оплату за определенный период
	case "3_months":
		//TODO привязать оплату за определенный период
	case "6_months":
		//TODO привязать оплату за определенный период
	case "year":
		//TODO привязать оплату за определенный период
	default:
		responseMessage = "Неизвестная команда."
	}

	// Отправляем ответ на callback
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, responseMessage)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}

	// Удаляем уведомление о нажатии кнопки
	if _, err := bot.AnswerCallbackQuery(tgbotapi.NewCallback(callback.ID, "")); err != nil {
		log.Printf("Ошибка при ответе на callback: %v", err)
	}
}
