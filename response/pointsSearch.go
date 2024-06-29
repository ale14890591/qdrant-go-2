package response

import (
	"encoding/json"
	"io"

	"golang.org/x/exp/constraints"
)

type PointsSearch[T string | constraints.Integer] struct {
	Response
	Result []PointsSearchResult[T] `json:"result"`
}

type PointsSearchResult[T string | constraints.Integer] struct {
	ID      T                      `json:"id"`
	Version uint64                 `json:"version"`
	Score   float64                `json:"score"`
	Payload map[string]interface{} `json:"payload,omitempty"`
	Vector  []float32              `json:"vector,omitempty"`
}

func (p *PointsSearch[T]) AcceptContentType() string {
	return "application/json"
}

func (p *PointsSearch[T]) Decode(body io.Reader) error {
	err := json.NewDecoder(body).Decode(p)
	if err != nil {
		return err
	}

	p.Response.SetStatusMessage()
	return nil
}
