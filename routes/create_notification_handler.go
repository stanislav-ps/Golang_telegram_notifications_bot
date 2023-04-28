package routes

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartHandler обрабатывает команду /create_notification
func CreateNotificationHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Отправляем пользователю сообщение
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Здесь реализую обработку создания уведомлений")
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
