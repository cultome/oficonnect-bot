package main

import (
	"log"
	"os"

	oficonnectbot "github.com/cultome/oficonnect-bot"
)

func main()  {
  oficonnect_id := os.Args[2]
  log.Printf("[*] Getting events for %s...", oficonnect_id)

  bot := oficonnectbot.Bot { OfiConnectID: oficonnect_id }

  events, err := bot.RetriveEvents()

  if err != nil {
    log.Fatalf(err.Error())
  }

  for _, evt := range events {
    if evt.Open == "1" {
      if evt.Confimed == "0" {
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

      log.Printf(" - %5t (%2s/xx) - %s\n", evt.Confimed == "1", evt.Quota, evt.EventName)
    }
  }
}
