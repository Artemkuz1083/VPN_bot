package botconn

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

var botConn *tgbotapi.BotAPI

func GetBotConnection() *tgbotapi.BotAPI {
	if botConn == nil {
		// Загружаем переменные окружения из .env файла
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Ошибка загрузки .env файла")
		}

		// Получаем токен бота
		token := os.Getenv("TELEGRAM_BOT_TOKEN")
		if token == "" {
			log.Fatal("TELEGRAM_BOT_TOKEN не установлен")
		}

		// Создаем нового бота
		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Авторизован как %s", bot.Self.UserName)

		// Сохраняем ссылку на бота в глобальной переменной
		botConn = bot
	}

	return botConn
}

func GetUpdates() tgbotapi.UpdatesChannel {
	// Создаем соединение с ботом
	bot := GetBotConnection()

	// Создаем обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Получаем канал обновлений
	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Fatal(err)
	}

	return updates
}
