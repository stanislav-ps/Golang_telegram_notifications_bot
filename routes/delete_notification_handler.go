package routes

import (
	"obit_bot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartHandler обрабатывает команду /delete_notification
func DeleteNotificationHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	utils.InfoLogger.Println("Received /delete_notification command")

	// Отправляем пользователю сообщение
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите уведомление, которое требуется удалить")

	_, err := bot.Send(msg)
	if err != nil {
		utils.ErrorLogger.Println(err)
	}
}
