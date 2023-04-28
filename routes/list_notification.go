package routes

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartHandler обрабатывает команду /list_notification
func ListNotificationHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Отправляем пользователю сообщение
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Все ваши уведомления")

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
