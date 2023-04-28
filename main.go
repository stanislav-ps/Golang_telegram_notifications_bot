package main

import (
	"log"
	"obit_bot/routes"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// вызов обработчика команды
		switch update.Message.Command() {
		case "start":
			routes.StartHandler(bot, update)
		case "create_notification":
			routes.CreateNotificationHandler(bot, update)
		case "edit_notification":
			routes.EditNotificationHandler(bot, update)
		case "delete_notification":
			routes.DeleteNotificationHandler(bot, update)
		case "list_notification":
			routes.ListNotificationHandler(bot, update)
		case "help":
			routes.HelpHandler(bot, update)
		case "support":
			routes.SupportHandler(bot, update)
		default:
			log.Printf("Unknown command: %s", update.Message.Text)
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}
