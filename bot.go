package telebotAPI

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Bot struct {
	Token string
}

func (bot *Bot) getBaseUrl() (string, error) {
	if bot.Token == "" {
		return "", errors.New("bot Token not provided")
	}
	return fmt.Sprintf("https://api.telegram.org/bot%v", bot.Token), nil
}

func (bot *Bot) makeRequestUrl(
	method string,
	params map[string]string,
) (requestUrl string, err error) {

	baseUrl, err := bot.getBaseUrl()

	if err != nil {
		return "", err
	}

	requestUrl = fmt.Sprintf("%v/%v?", baseUrl, method)

	var paramStrings []string

	for paramName, paramValue := range params {
		paramString := fmt.Sprintf("%v=%v", paramName, paramValue)
		paramStrings = append(paramStrings, paramString)
	}
	requestUrl += strings.Join(paramStrings, "&")
	return
}

func (bot *Bot) sendRequest(
	method string,
	params map[string]string,
) (err error) {
	requestUrl, err := bot.makeRequestUrl(method, params)

	if err != nil {
		return err
	}

	response, err := http.Get(requestUrl)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New(fmt.Sprintf("bad response at %v", method))
	}

	return
}

func (bot *Bot) UpdatesGoroutine(
	updatesChan chan Update,
	errorChan chan error,
	interval time.Duration,
) {
	offset := 1

	for {
		updates, err := bot.getUpdates(offset)

		for _, update := range updates {
			updatesChan <- update
			offset = update.UpdateID + 1
		}

		if err != nil {
			errorChan <- err
		}

		time.Sleep(interval)
	}
}
