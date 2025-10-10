package main

import (
	"log"
	"os"
	"slices"

	oficonnectbot "github.com/cultome/oficonnect-bot"
	"github.com/fatih/color"
)

func main() {
	oficonnect_id := os.Args[1]
	checkOnly := os.Args[2] == "true"
	// log.Printf("[*] Getting events for %s...", oficonnect_id)

	if checkOnly {
		log.Printf("[*] !!!! CHECK ONLY MODE !!!!!!")
	}

	config := oficonnectbot.ReadConfig()
	bot := oficonnectbot.BuildBot(oficonnect_id)

	info, err := bot.RetrivePersonalInformation()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Events for Marshal [%d] %s %s", info.ID, info.Name, info.LastName)

	events, err := bot.RetriveEvents()

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, evt := range events {
		if evt.Open == 1 {
			if !checkOnly {
				if evt.Confimed == 0 {
					if isExcluded(config.Excludes, evt) {
						continue
					}

					tryToRegister(evt, bot)
				}
			}

			confirmations, _ := bot.RetriveConfirmationsByEvent(evt.EventID)

			confirm := "Sin confirmar"
			if evt.Confimed == 1 {
				confirm = "Confirmado"
			}

			if isExcluded(config.Excludes, evt) {
				continue
			}

			color.Blue("[%d] {%s} (%2d/%2s) - %s\n", evt.EventID, confirm, confirmations, evt.Quota, evt.EventName)
		}
	}
}

func tryToRegister(evt *oficonnectbot.Event, bot *oficonnectbot.Bot) {
	log.Printf("[*] Intentando registrarte para [%s]...", evt.EventName)

	registrationResponse, err := bot.RegisterForEvent(evt)

	if err != nil {
		log.Fatal(err.Error())
	}

	if registrationResponse.Message == "Cupo lleno" {
		color.Red("[-] El registro para [%s] esta lleno!", evt.EventName)
	} else if registrationResponse.Status == "error" {
		color.Red("[-] Ocurrio un error al registrate a [%s]! %s", evt.EventName, registrationResponse.Message)
	} else {
		color.Red("[+] Te acabas de registrar para [%s]!", evt.EventName)
	}
}

func isExcluded(excludes []int, evt *oficonnectbot.Event) bool {
	return slices.Contains(excludes, evt.EventID)
}
