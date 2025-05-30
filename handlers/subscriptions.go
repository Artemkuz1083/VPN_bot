package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Subscription(bot *tgbotapi.BotAPI, chatID int64) {
	msgText := "Ниже представлены варианты подписок"

	// Создаем кнопки
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

	msg := tgbotapi.NewMessage(chatID, msgText)
	msg.ReplyMarkup = &keyboard

	_, err := bot.Send(msg)
	if err != nil {
		log.Fatalf("Ошибка при редактировании сообщения: %v", err)
	}
}
