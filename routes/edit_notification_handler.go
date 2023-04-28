package routes

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartHandler обрабатывает команду /edit_notification
func EditNotificationHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Отправляем пользователю сообщение
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите уведомление, которое требуется изменить")

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
