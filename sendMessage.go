package tgbotbasicapi

func (bot *Bot) SendMessage(chatId int, text string) (err error){
	err = bot.sendRequest(
		"sendMessage",
		map[string]string{
			"text": text,
		})
	return
}
