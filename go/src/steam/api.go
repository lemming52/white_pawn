package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	key := os.Getenv("STEAM_API_KEY")
	response, err := getDotaHeroes(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, hero := range response.Result.Heroes {
		printHero(hero)
	}

}

type heroResponse struct {
	Result *heroResult `json:"result"`
}

type heroResult struct {
	Heroes []*HeroID `json:"heroes"`
	Status int       `json:"status"`
	Count  int       `json:"count"`
}

type HeroID struct {
	Name string `json:"name"`
	ID   int8   `json:"id"`
}

func getDotaHeroes(key string) (*heroResponse, error) {
	const root = "http://api.steampowered.com/IEconDOTA2_570/GetHeroes/v0001"
	uri := root + "?key=" + key
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Request Failed: %s", resp.Status)
	}

	var response heroResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &response, nil
}

func printHero(hero *HeroID) {
	fmt.Printf("Hero: %s, %d\n", hero.Name, hero.ID)
}
