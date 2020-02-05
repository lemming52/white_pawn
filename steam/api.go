package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"steam/clients"
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
	ID   int16  `json:"id"`
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
	HeroID     int16 `json:"hero_id"`
}

func main() {
	key := os.Getenv("STEAM_API_KEY")

	dClient := clients.NewDota2Client(key)
	serverID := "90128209361158154"
	stats, err := dClient.GetRealtimeStats(serverID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stats.CurrentStatus().Print())
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
