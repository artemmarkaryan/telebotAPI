package telebotAPI

import "strconv"

func (bot *Bot) SendMessage(chatId int, text string) (err error){
	err = bot.sendRequest(
		"sendMessage",
		map[string]string{
			"chat_id": strconv.Itoa(chatId),
			"text": text,
		})
	return
}
