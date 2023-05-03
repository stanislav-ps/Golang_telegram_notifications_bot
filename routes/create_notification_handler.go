package routes

import (
	"obit_bot/keyboards"
	"obit_bot/services"
	"obit_bot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// CreateNotificationHandler обработчик команды /create_notification
func CreateNotificationHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Запрашиваем у пользователя время и дату уведомления
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите время и дату уведомления (например, 13:30 01.05.2023):")
	if _, err := bot.Send(msg); err != nil {
		utils.ErrorLogger.Fatal(err)
	}

	updates := bot.GetUpdatesChan(tgbotapi.UpdateConfig{Timeout: 30})
	responseChan := make(chan tgbotapi.Update, 1)
	updates = append(updates, responseChan...)
	response := <-responseChan

	// Запрашиваем у пользователя текст уведомления
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Введите текст уведомления:")
	if _, err := bot.Send(msg); err != nil {
		utils.ErrorLogger.Println("Failed to send message:", err)
		return
	}

	// Ждем ответа от пользователя
	response = <-responseChan

	// Получаем текст сообщения с уведомлением
	text := response.Message.Text

	// Показываем пользователю клавиатуру для выбора частоты уведомления
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите частоту уведомления:")
	msg.ReplyMarkup = keyboards.NotificationKeyboard()
	if _, err := bot.Send(msg); err != nil {
		utils.ErrorLogger.Println("Failed to send message:", err)
		return
	}

	// Ждем ответа от пользователя
	response = <-responseChan

	// Получаем текст выбранной частоты уведомления
	frequency := response.Message.Text

	// Создаем структуру уведомления
	notification := models.Notification{
		ChatID:    update.Message.Chat.ID,
		Text:      text,
		TimeDate:  timeDate,
		Frequency: frequency,
	}

	// Сохраняем уведомление в базе данных
	if err := notification.Create(services.DB.Conn); err != nil {
		utils.ErrorLogger.Println("Failed to create notification:", err)
		return
	}

	// Отправляем сообщение пользователю, что уведомление создано
	msg = tgbotapi.NewMessage(update.Message.Chat.ID,
		"Уведомление создано:\n"+
			"Текст: "+notification.Text+"\n"+
			"Частота: "+notification.Frequency+"\n"+
			"Время: "+notification.Time.String())
	if _, err := bot.Send(msg); err != nil {
		utils.ErrorLogger.Println("Failed to send message:", err)
		return
	}
}
