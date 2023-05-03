package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// NotificationKeyboard возвращает клавиатуру для обработки уведомлений
func NotificationKeyboard() tgbotapi.ReplyKeyboardMarkup {
	// Создаем клавиатуру с кнопками
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Однократное"),
			tgbotapi.NewKeyboardButton("Каждый день"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Каждую неделю"),
			tgbotapi.NewKeyboardButton("Каждый месяц"),
		),
	)

	// Включаем режим автоматического скрытия клавиатуры
	keyboard.OneTimeKeyboard = true

	return keyboard
}
