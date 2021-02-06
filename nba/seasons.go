package nba

import (
	"context"
	"net/http"
)

type Seasons struct {
    API struct {
        Seasons []string `json:"seasons"`
    } `json:"api"`
}

func (c *Client) GetSeasons(ctx context.Context) (*interface{}, error) {
	relativePath := "seasons/"
	seasons := new(Seasons)
	req, err := c.GetRequestResult(ctx,http.MethodGet, relativePath, "", seasons)

	if err != nil {
		return nil, err
	}

	return &req, err
	
}
