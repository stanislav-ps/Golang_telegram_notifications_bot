package routes

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartHandler обрабатывает команду /start
func DeleteNotificationHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Отправляем пользователю сообщение
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите уведомление, которое требуется удалить")

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
