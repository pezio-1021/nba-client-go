package nba

import (
	"context"
	"net/http"
)

type Players struct {
	API struct {
		Players []struct {
			FirstName         string `json:"firstName"`
			LastName          string `json:"lastName"`
			TeamID            string `json:"teamId"`
			YearsPro          string `json:"yearsPro"`
			CollegeName       string `json:"collegeName"`
			Country           string `json:"country"`
			PlayerID          string `json:"playerId"`
			DateOfBirth       string `json:"dateOfBirth"`
			Affiliation       string `json:"affiliation"`
			StartNba          string `json:"startNba"`
			HeightInMeters    string `json:"heightInMeters"`
			WeightInKilograms string `json:"weightInKilograms"`
			Leagues struct {
				Standard []struct {
					Jersey string `json:"jersey"`
					Active string `json:"active"`
					Pos    string `json:"pos"`
				} `json:"standard,omitempty"`
				Vegas []struct {
					Jersey string `json:"jersey"`
					Active string `json:"active"`
					Pos    string `json:"pos"`
				} `json:"standard,omitempty"`
			} `json:"leagues,omitempty"`
		} `json:"players"`
	} `json:"api"`
}

func (c *Client) GetPlayersFromTeamID(ctx context.Context, playerId string) (*interface{}, error) {
	relativePath := "players/teamId/"
	players := new(Players)
	req, err := c.GetRequestResult(ctx,http.MethodGet, relativePath, playerId, players)
	if err != nil {
		return nil, err
	}

	return &req, err
}

func (c *Client) GetPlayersFromID(ctx context.Context, playerId string) (*interface{}, error) {
	relativePath := "players/playerId/"
	players := new(Players)
	req, err := c.GetRequestResult(ctx,http.MethodGet, relativePath, playerId, players)
	if err != nil {
		return nil, err
	}

	return &req, err
}

