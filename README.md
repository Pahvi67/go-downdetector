# Go Downdetector

Lightweight and simple downdetector that'll notify user through a Telegram bot if the given URL does not respond with HTTP 200.

## Installation

```bash
$ git clone https://github.com/Pahvi67/go-downdetector.git
$ cd go-downdetector
$ go mod download # if running locally
$ cp .env.example .env # edit after copying
```
### .env values
| Env variable        | Example             | Description                                      |
| ------------------- | ------------------- | ------------------------------------------------ |
| ENV                 | `development`       | runs either `go run .` or `go build` accordingly |
| SITE_URL            | `www.google.com`    | the site url that'll be pinged                   |
| TELEGRAM_BOT_STRING | `bot58...:AAGxV...` | Telegram bot API string                          |
| TELEGRAM_CHAT_ID    | `837402353`         | Chat ID to which the bot should send the message |
| INTERVAL_DURATION   | `15`                | Interval in minutes                              |


## Usage
```bash
$ ./start.sh
# or if using docker:
$ docker compose up
```