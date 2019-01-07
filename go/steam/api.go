package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	key := os.Getenv("STEAM_API_KEY")
	get_dota_heroes(key)
}

func get_dota_heroes(key string) {
	const root = "http://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001"
	uri := root + "?key=" + key
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", uri, err)
		os.Exit(1)
	}
	err = ioutil.WriteFile("dota_heroes.json", b, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", uri, err)
		os.Exit(1)
	}
}
