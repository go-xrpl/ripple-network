package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	ripplenetwork "github.com/goxrp/ripple-network"
)

func main() {
	gateways := ripplenetwork.Gateways()
	for i, gw := range gateways {
		err := gw.LoadAccountRoot()
		if err != nil {
			log.Fatal(err)
		}
		gateways[i] = gw
	}
	bytes, err := json.MarshalIndent(gateways, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("gateways.json", bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DONE")
}