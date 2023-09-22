package request

import (
	"fmt"
	"io"

	"github.com/henomis/restclientgo"
)

type PointsGet struct {
	CollectionName string  `json:"-"`
	Consistency    *string `json:"-"`
	IDs            []string
	WithPayload    *bool `json:"with_payload,omitempty"`
	WithVector     *bool `json:"with_vector,omitempty"`
}

func (p *PointsGet) Path() (string, error) {
	var urlValues restclientgo.URLValues
	urlValues.Add("consistency", p.Consistency)

	parameters := ""
	if len(urlValues) > 0 {
		parameters = "?" + urlValues.Encode()
	}

	return fmt.Sprintf("/collections/%s/points%s", p.CollectionName, parameters), nil
}

func (p *PointsGet) Encode() (io.Reader, error) {
	return nil, nil
}

func (p *PointsGet) ContentType() string {
	return ""
}
