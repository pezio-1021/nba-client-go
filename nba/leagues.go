package nba

import (
	"context"
	"net/http"
)

type Leagues struct {
    API struct {
        Leagues []string `json:"leagues"`
    } `json:"api"`
}

func (c *Client) GetLeagues(ctx context.Context) (*interface{}, error) {
	relativePath := "leagues/"
	leagues := new(Leagues)
	req, err := c.GetRequestResult(ctx,http.MethodGet, relativePath, "", leagues)

	if err != nil {
		return nil, err
	}

	return &req, err
	
}
