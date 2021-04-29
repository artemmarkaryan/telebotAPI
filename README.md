Basic telegram bot api for Go. Polls updates and provides wrapper to make requests.

## Usage

``` go
var bot = telebotapi.Bot{Token: "bot_token"}
var updatesChan = make(chan telebotapi.Update)
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
