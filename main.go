package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/mdns"
)

func main() {
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			json, err := json.MarshalIndent(entry, "", "	")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(json))
			os.Exit(1)
		}
	}()
	mdns.Lookup("_googlecast._tcp", entriesCh)
}
