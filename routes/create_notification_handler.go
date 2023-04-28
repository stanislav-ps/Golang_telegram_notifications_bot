package routes

import (
	"strconv"

	"obit_bot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// CreateNotificationHandler обрабатывает команду /create_notification
func CreateNotificationHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	utils.InfoLogger.Println("Received /create_notification command")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите время и дату уведомления (например, 13:30 01.05.2023):")
	bot.Send(msg)

	// Получаем ответ пользователя
	updates, err := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
		Offset: update.UpdateID + 1,
	})
	if err != nil {
		utils.ErrorLogger.Println(err)
		return
	}

	var dateTime string
	var frequency int // перемещаем объявление переменной сюда
	for resp := range updates {
		if resp.Message == nil || resp.Message.Chat.ID != update.Message.Chat.ID {
			continue
		}
		dateTime = resp.Message.Text
		break
	}

	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Введите частоту оповещения (в минутах):")
	bot.Send(msg)

	// Получаем ответ пользователя
	for resp := range updates {
		if resp.Message == nil || resp.Message.Chat.ID != update.Message.Chat.ID {
			continue
		}
		var err error // добавляем объявление переменной err
		frequency, err = strconv.Atoi(resp.Message.Text)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректное значение частоты оповещения")
			bot.Send(msg)
			utils.ErrorLogger.Println(err)
			return
		}
		break
	}

	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Уведомление создано:\nВремя: "+dateTime+"\nЧастота: "+strconv.Itoa(frequency))
	bot.Send(msg)
}
