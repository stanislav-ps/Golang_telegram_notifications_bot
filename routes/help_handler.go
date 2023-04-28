package routes

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartHandler обрабатывает команду /help
func HelpHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Отправляем пользователю сообщение
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список команд "+
		"Доступные команды:\n"+
		"/create_notification - Создать уведомление\n"+
		"/edit_notification - Редактировать уведомление\n"+
		"/delete_notification - Удалить уведомление\n"+
		"/list_notification - Посмотреть список уведомлений\n"+
		"/help - Помощь\n"+
		"/support - Написать в поддержку")
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
