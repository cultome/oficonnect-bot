package main

import (
	"log"
	"os"

	oficonnectbot "github.com/cultome/oficonnect-bot"
)

func main() {
	oficonnect_id := os.Args[1]
	log.Printf("[*] Getting events for %s...", oficonnect_id)

	config := oficonnectbot.ReadConfig()
	bot := oficonnectbot.BuildBot(oficonnect_id)

	info, err := bot.RetrivePersonalInformation()

	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Events for Marshal [%s] %s %s", info.ID, info.Name, info.LastName)

	events, err := bot.RetriveEvents()

	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, evt := range events {
		if evt.Open == "1" {
			if evt.Confimed == "0" {
				if !isExcluded(config.Excludes, evt) {
					tryToRegister(evt, bot)
				}
			}

			confirmations, _ := bot.RetriveConfirmationsByEvent(evt.EventID)
			log.Printf("[%s] %5t (%2s/%2d) - %s\n", evt.EventID, evt.Confimed == "1", evt.Quota, confirmations, evt.EventName)
		}
	}
}

func tryToRegister(evt *oficonnectbot.Event, bot *oficonnectbot.Bot) {
	log.Printf("[*] Intentando registrarte para [%s]...", evt.EventName)

	registrationResponse, err := bot.RegisterForEvent(evt)

	if err != nil {
		log.Fatalf(err.Error())
	}

	if registrationResponse.Status == "lleno" {
		log.Printf("[-] El registro para [%s] esta lleno!", evt.EventName)
	} else if registrationResponse.Status == "error" {
		log.Printf("[-] Ocurrio un error al registrate a [%s]!", evt.EventName)
	} else {
		log.Printf("[+] Te acabas de registrar para [%s]!", evt.EventName)
	}
}

func isExcluded(excludes []string, evt *oficonnectbot.Event) bool {
	for _, eventID := range excludes {
		if eventID == evt.EventID {
			log.Printf("[-] Excluding event [%s]!", evt.EventID)
			return true
		}
	}

	return false
}
