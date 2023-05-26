package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"net/http"
)

var siteDown bool = false
var siteUrl string
var botString string
var chatId string

func main() {
	// get env variables
	siteUrl = os.Getenv("SITE_URL")
	botString = os.Getenv("TELEGRAM_BOT_STRING")
	chatId = os.Getenv("TELEGRAM_CHAT_ID")
	intervalDuration, err := strconv.Atoi(os.Getenv("INTERVAL_DURATION"))
	if err != nil {
		panic("Interval duration invalid/not present!")
	}

	// start loop
	ticker := time.NewTicker(time.Duration(intervalDuration) * time.Minute)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				checkSiteStatus()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	log("Starting website downdetector on: " + siteUrl)
	log("Interval set to " + strconv.Itoa(intervalDuration) + " minutes.")
	checkSiteStatus()
	select {}
}

// Check if site is online and send out message if it's not
func checkSiteStatus() {
	response, err := http.Get(siteUrl)
	if err != nil {
		if !siteDown {
			log(err.Error())
		}
		handleSiteDown(0)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		log("status 200")
		if siteDown {
			log(colorGreen("Site back up"))
			siteDown = false
		}
	} else {
		handleSiteDown(response.StatusCode)
	}
}

// Send error message to telegram and print to log
func handleSiteDown(status int) {
	if siteDown {
		return
	}
	siteDown = true
	log(colorRed("--- WEBSITE DOWN! ---"))

	text := `
		<b>Website down!</b>
		
		HTTP Status: ` + strconv.Itoa(status) + `
		
		<a href='` + siteUrl + `'>` + siteUrl + `</a>
	`
	replacedText := strings.ReplaceAll(text, "\t\t", "")

	url := "https://api.telegram.org/" + botString + "/sendMessage"
	var jsonData = []byte(`{
		"chat_id": "` + chatId + `",
		"parse_mode": "html",
		"text": "` + replacedText + `",
	}`)
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if error != nil {
		log(colorRed("ERROR!") + " Could not make a request")
		log(error.Error())
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		log(colorRed("ERROR!") + " Could not make request to Telegram Bot API.")
		log(error.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		log(colorRed("ERROR!") + " Telegram message could not be sent.")
	}
}

// Color string red
func colorRed(msg string) string {
	return "\033[31m" + msg + "\033[0m"
}

// Color string green
func colorGreen(msg string) string {
	return "\033[32m" + msg + "\033[0m"
}

// Log message with a timestamp
func log(msg string) {
	fmt.Println(
		"["+time.Now().Format("2006-01-02 15:04:05")+"]",
		msg,
	)
}
