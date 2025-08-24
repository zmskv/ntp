package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zmskv/ntp/internal/fetcher"
)

func main() {
	log.SetOutput(os.Stderr)

	server := "pool.ntp.org"
	if len(os.Args) > 1 {
		server = os.Args[1]
	}

	if err := fetcher.ValidateServer(server); err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}

	currentTime, err := fetcher.Fetch(server)
	if err != nil {
		log.Printf("error get time: %v", err)
		os.Exit(1)
	}

	fmt.Println(fetcher.FormatRFC1123(currentTime))
}
