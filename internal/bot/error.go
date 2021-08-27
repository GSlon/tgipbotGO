package bot

func (b *Bot) handleError(chatID int64, err error) {
	b.SendMessage(chatID, err.Error())
}