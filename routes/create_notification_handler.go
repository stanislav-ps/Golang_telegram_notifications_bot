package routes

import (
	"time"

	"obit_bot/keyboards"
	"obit_bot/models"
	"obit_bot/services"
	"obit_bot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var response tgbotapi.Update

// CreateNotificationHandler обработчик команды /create_notification
func CreateNotificationHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Запрашиваем у пользователя время и дату уведомления
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите время и дату уведомления (например, 13:30 01.05.2023):")
	if _, err := bot.Send(msg); err != nil {
		utils.ErrorLogger.Fatal(err)
	}

	// Ждем ответа от пользователя
	response := waitResponse(bot)

	// Получаем текст сообщения с временем и датой уведомления
	timeDateString := response.Message.Text

	// Парсим время и дату уведомления
	timeDate, err := time.Parse("15:04 02.01.2006", timeDateString)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректный формат даты и времени. Попробуйте еще раз.")
		if _, err := bot.Send(msg); err != nil {
			utils.ErrorLogger.Println("Failed to send message:", err)
			return
		}
		return
	}

	// Запрашиваем у пользователя текст уведомления
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Введите текст уведомления:")
	if _, err := bot.Send(msg); err != nil {
		utils.ErrorLogger.Println("Failed to send message:", err)
		return
	}

	// Ждем ответа от пользователя
	response = waitResponse(bot)

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
	response = waitResponse(bot)

	// Получаем текст выбранной частоты уведомления
	frequencyString := response.Message.Text

	// Парсим выбранную частоту уведомления
	frequency, err := parseFrequency(frequencyString)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректная частота уведомления. Попробуйте еще раз.")
		if _, err := bot.Send(msg); err != nil {
			utils.ErrorLogger.Println("Failed to send message:", err)
			return
		}
		return
	}
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
			"Время: "+notification.TimeDate.String())
	if _, err := bot.Send(msg); err != nil {
		utils.ErrorLogger.Println("Failed to send message:", err)
		return
	}
}
