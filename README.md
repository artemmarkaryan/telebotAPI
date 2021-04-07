# tg-bot-basic-api
Basic telegram bot api for Go. Polls updates and provides wrapper to make requests.

## Usage

``` go
var bot = tgMessageApi.Bot{Token: "bot_token"}
var updatesChan = make(chan tgMessageApi.Update)
var errChan = make(chan error)

go bot.UpdatesGoroutine(updatesChan, errChan, time.Second/20)

for {
  select {
  case update := <-updatesChan:
    log.Printf("update: %v", update)

  case err := <-errChan:
    log.Printf("error: %v", err)
  }
}
```
