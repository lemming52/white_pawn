package model

import "fmt"

type StatsResponse struct {
	Match      *MatchStats      `json:"match"`
	Teams      []*TeamStats     `json:"teams"`
	Buildings  []*BuildingStats `json:"buildings"`
	GraphData  *GraphData       `json:"graph_data"`
	DeltaFrame bool             `json:"delta_frame"`
}

type MatchStats struct {
	ServerID     int   `json:"server_steam_id"`
	MatchID      int   `json:"matchid"`
	Timestamp    int16 `json:"timestamp"`
	GameTime     int16 `json:"game_time"`
	GameMode     uint8 `json:"game_mode"`
	LeagueID     int   `json:"league_id"`
	LeagueNodeID int   `json:"league_node_id"`
	GameState    uint8 `json:"game_state"`
}

type TeamStats struct {
	TeamNumber  uint8          `json:"team_number"`
	TeamID      int            `json:"team_id"`
	TeamName    string         `json:"team_name"`
	TeamTag     string         `json:"team_tag"`
	TeamLogo    int            `json:"team_logo"`
	Score       int            `json:"score"`
	NetWorth    int            `json:"net_worth"`
	TeamLogoURL string         `json:"team_logo_url"`
	Players     []*PlayerStats `json:"players"`
}

type PlayerStats struct {
	AccountID int      `json:"accountid"`
	PlayerID  uint8    `json:"playerid"`
	Name      string   `json:"name"`
	Team      uint8    `json:"team"`
	HeroID    uint16   `json:"heroid"`
	Level     uint8    `json:"level"`
	Kills     uint8    `json:"kill_count"`
	Deaths    uint8    `json:"death_count"`
	Assists   uint8    `json:"assists_count"`
	Denies    uint8    `json:"denies_count"`
	LastHists uint16   `json:"lh_count"`
	Gold      uint16   `json:"gold"`
	XCoord    float64  `json:"x"`
	YCoord    float64  `json:"y"`
	NetWorth  uint32   `json:"net_worth"`
	Abilities []uint16 `json:"abilities"`
	Items     []uint16 `json:"items"`
}

type BuildingStats struct {
	Team         uint8   `json:"team"`
	Heading      float64 `json:"heading"`
	BuildingType uint8   `json:"type"`
	Lane         uint8   `json:"lane"`
	Tier         uint8   `json:"tier"`
	XCoord       float64 `json:"x"`
	YCoord       float64 `json:"y"`
	Destroyed    bool    `json:"destroyed"`
}

type GraphData struct {
	GraphGold []int `json:"graph_gold"`
}

type Status struct {
	GameTime           string `json:"game_time"`
	RadiantKills       int    `json:"radiant_kills"`
	DireKills          int    `json:"dire_kills"`
	RadiantNetworth    int    `json:"net_worth_radiant"`
	DireNetworth       int    `json:"net_worth_dire"`
	NetWorthDifference int    `json:"net_worth_diff"`
}

func (st *Status) Print() string {
	return fmt.Sprintf("Time: %s\nRadiant Kills: %d\nDire Kills: %d\nRadiant NW: %d\nDire NW: %d\nDelta NW: %d", st.GameTime, st.RadiantKills, st.DireKills, st.RadiantNetworth, st.DireNetworth, st.NetWorthDifference)
}

func (stats *StatsResponse) CurrentStatus() *Status {
	var radTeam, direTeam *TeamStats
	for _, team := range stats.Teams {
		if team.TeamNumber == 2 {
			radTeam = team
		} else if team.TeamNumber == 3 {
			direTeam = team
		}
	}
	return &Status{
		GameTime:           fmt.Sprintf("%d:%d", stats.Match.GameTime/60, stats.Match.GameTime%60),
		RadiantKills:       radTeam.Score,
		DireKills:          direTeam.Score,
		RadiantNetworth:    radTeam.NetWorth,
		DireNetworth:       direTeam.NetWorth,
		NetWorthDifference: radTeam.NetWorth - direTeam.NetWorth,
	}
}
