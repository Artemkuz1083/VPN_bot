package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
)

// Сообщение для нового пользователя
func StartMenuForNewUsers(bot *tgbotapi.BotAPI, chatID int64) {
	// Создаем сообщение
	welcomeMessage := "Добро пожаловать в kuzzbass VPN!\n" +
		"Мы рады, что вы присоединились к нам. Здесь вы сможете воспользоваться нашими услугами и получить доступ к множеству полезных функций.\n\n" +
		"Мы выдаем vless ключи, которые позволят вам безопасно и анонимно пользоваться интернетом. Просто введите ваш vless ключ в настройках VPN-клиента, и вы сможете начать пользоваться нашим сервисом.\n\n" +
		"Если у вас возникнут какие-либо вопросы или проблемы, не стесняйтесь обращаться к нашей службе поддержки. Мы всегда готовы помочь!\n\n" +
		"Желаем вам приятного использования нашего VPN-сервиса."

	// Создаем кнопки
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Получить ключ", "get_key"),
			tgbotapi.NewInlineKeyboardButtonData("Продлить подписку", "extend_subscription"),
		),
	)

	// Создаем сообщение с кнопками
	msg := tgbotapi.NewMessage(chatID, welcomeMessage)
	msg.ReplyMarkup = keyboard

	// Отправляем сообщение
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatalf("Ошибка при отправке сообщения: %v", err)
	}
}
