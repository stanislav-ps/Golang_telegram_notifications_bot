package routes

import (
	"obit_bot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartHandler обрабатывает команду /list_notification
func SupportHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	utils.InfoLogger.Println("Received /support command")

	// Отправляем пользователю сообщение
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Подробно опишите вашу проблему и ждите ответа от поддержки")

	_, err := bot.Send(msg)
	if err != nil {
		utils.ErrorLogger.Println(err)
	}
}
