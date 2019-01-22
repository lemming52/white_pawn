package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

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

type matchesResponse struct {
	Result *matchesResult `json:"result"`
}

type matchesResult struct {
	Status           int      `json:"status"`
	NumResults       int      `json:"num_results"`
	TotalResults     int      `json:"total_results"`
	ResultsRemaining int      `json:"results_remaining"`
	Matches          []*Match `json:"matches"`
}

type Match struct {
	MatchID       int       `json:"match_id"`
	MatchSeqNum   int       `json:"match_seq_num"`
	StartTime     int       `json:"start_time"`
	LobbyType     int       `json:"lobby_type"`
	RadiantTeamID int       `json:"radiant_team_id"`
	DireTeamID    int       `json:"dire_team_id"`
	Players       []*Player `json:"players"`
}

type Player struct {
	AccountID  int   `json:"account_id"`
	PlayerSlot uint8 `json:"player_slot"`
	HeroID     int8  `json:"hero_id"`
}

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

	matchResponse, err := getMatchHistory(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	lobbyTypeMap := make(map[int]int)
	for _, match := range matchResponse.Result.Matches {
		_, ok := lobbyTypeMap[match.LobbyType]
		if ok {
			lobbyTypeMap[match.LobbyType]++
		} else {
			lobbyTypeMap[match.LobbyType] = 1
		}
	}

	for lobbyType, count := range lobbyTypeMap {
		fmt.Printf("Lobby Type: %d\tCount: %d\n", lobbyType, count)
	}

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

func getMatchHistory(key string) (*matchesResponse, error) {
	const root = "http://api.steampowered.com/IDOTA2Match_570/GetMatchHistory/v1"
	uri := root + "?key=" + key
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Request Failed: %s", resp.Status)
	}

	var response matchesResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &response, nil
}
