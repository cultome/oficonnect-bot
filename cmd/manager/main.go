package main

import (
	"flag"
	"log"

	oficonnectbot "github.com/cultome/oficonnect-bot"
)

func main() {
	var exclusion = flag.String("exclude", "", "Exclude an event ID from registering")
	flag.Parse()

	log.Printf("[*] Add eventID [%s] to register exclusion...", *exclusion)

	config := oficonnectbot.ReadConfig()
	config.Excludes = append(config.Excludes, *exclusion)
	config.Persist()

	log.Printf("[+] Done!")
}
