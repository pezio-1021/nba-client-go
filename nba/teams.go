package nba

import (
	"context"
	"net/http"
)

type Teams struct {
    API struct {
			Teams   []struct {
				City         string `json:"city"`
				FullName     string `json:"fullName"`
				TeamID       string `json:"teamId"`
				Nickname     string `json:"nickname"`
				Logo         string `json:"logo"`
				ShortName    string `json:"shortName"`
				AllStar      string `json:"allStar"`
				NbaFranchise string `json:"nbaFranchise"`
				Leagues      struct {
					Standard struct {
						ConfName string `json:"confName"`
						DivName  string `json:"divName"`
					} `json:"standard"`
				} `json:"leagues"`
			} `json:"teams"`
    } `json:"api"`
}

func (c *Client) GetTeamsList(ctx context.Context, league string) (*interface{}, error) {
	relativePath := "teams/league/"
	teams := new(Teams)
	req, err := c.GetRequestResult(ctx,http.MethodGet, relativePath, league, teams)
	if err != nil {
		return nil, err
	}

	return &req, err
}

func (c *Client) GetTeam(ctx context.Context, teamID string) (*interface{}, error) {
	relativePath := "teams/teamId/"
	teams := new(Teams)
	req, err := c.GetRequestResult(ctx,http.MethodGet, relativePath, teamID, teams)
	if err != nil {
		return nil, err
	}

	return &req, err
}
