package main

import (
	"MixkoPay/utils/config"
	"log"
	"strconv"
	"strings"
	"time"

	pp "github.com/Frontware/promptpay"
	telegram "gopkg.in/telebot.v3"
)

func main() {
	pref := telegram.Settings{
		Token:  config.C.TelegramToken,
		Poller: &telegram.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telegram.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(telegram.OnText, func(c telegram.Context) error {
		text := strings.Split(c.Text(), " ")
		amount, err := strconv.Atoi(text[0])
		var bank string

		if len(text) > 1 {
			bank = text[1]
		}

		if err != nil {
			return c.Send("Invalid number!")
		}

		var promptPayBank string
		if bank == "kbank" || bank == "k" {
			promptPayBank = config.C.PromptPayKbankId
			bank = "KBank"
		} else {
			// * Default to SCB
			promptPayBank = config.C.PromptPayScbId
			bank = "SCB"
		}

		// * Generate QR code
		payment := pp.PromptPay{
			PromptPayID: promptPayBank,   // Tax-ID/ID Card/E-Wallet
			Amount:      float64(amount), // Positive amount
		}

		// * Generate string to be use in QRCode
		qrcode, _ := payment.Gen()

		// * Send QR code to user
		v := &telegram.Photo{File: telegram.FromURL("https://chart.googleapis.com/chart?cht=qr&chs=500x500&chl=" + qrcode)}

		// * Send messages
		if sendErr := c.Send(bank); sendErr != nil {
			log.Fatal(sendErr)
		}
		if sendErr := c.Send(v); sendErr != nil {
			log.Fatal(sendErr)
		}
		return nil
	})

	b.Start()
}
