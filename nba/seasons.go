package nba

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path"
)

type Seasons struct {
    API struct {
        Seasons []string `json:"seasons"`
    } `json:"api"`
}

func (c *Client) GetSeasons(ctx context.Context) (*Seasons, error) {
	relativePath := path.Join("seasons")
	req, err := c.NewRequest(ctx, http.MethodGet, relativePath, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	// send request
	seasons := new(Seasons)
	code, err := c.DoRequest(req, &seasons)
	if (err != nil) {
		return nil, err
	}

	switch code {
	case http.StatusOK:
		return seasons, nil
	case http.StatusBadRequest:
		return nil, errors.New("bad request. some parameters may be invalid")
	case http.StatusNotFound:
		return nil, fmt.Errorf("not found. user with id '%s' may not exist")
	default:
		return nil, errors.New("unexpected error1")
	}
}
