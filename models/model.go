package models

// Message is the structure of message sent in telegram
type Message struct {
	ChatID      string      `json:"chat_id"`
	Text        string      `json:"text"`
	ParseMode   string      `json:"parse_mode"`
	// ReplyMarkup ReplyMarkup `json:"reply_markup"`
}

type ReplyMarkup struct {
	InlineKeyBoard [][]InlineButton `json:"inline_keyboard"`
}

type InlineButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
	Url          string `json:"url"`
}
