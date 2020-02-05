package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"steam/model"
)

type DotaClient struct {
	authKey string
}

func NewDota2Client(key string) *DotaClient {
	return &DotaClient{
		authKey: key,
	}
}

func (dc *DotaClient) GetRealtimeStats(serverID string) (*model.StatsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", BaseURI, RealTimeStatsURI), nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("key", dc.authKey)
	q.Add("server_steam_id", serverID)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	res, err := client.Do(req)
	fmt.Println(fmt.Sprintf("%s/%s", BaseURI, RealTimeStatsURI))
	if err != nil {
		return nil, err
	}

	var response model.StatsResponse
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
