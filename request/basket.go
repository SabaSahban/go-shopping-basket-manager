package request

import "encoding/json"

type (
	CreateBasketRequest struct {
		Data   json.RawMessage `json:"data"`
		State  string          `json:"state"`
		UserID uint            `json:"user_id"`
	}

	UpdateBasketRequest struct {
		Data  json.RawMessage `json:"data"`
		State string          `json:"state"`
	}
)
