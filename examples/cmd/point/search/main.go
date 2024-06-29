package main

import (
	"context"
	"fmt"

	qdrantgo "github.com/ale14890591/qdrant-go-2"
	"github.com/ale14890591/qdrant-go-2/request"
	"github.com/ale14890591/qdrant-go-2/response"
)

func main() {

	client := qdrantgo.New("http://localhost:6333", "")

	withPayload := true
	resp := &response.PointsSearch{}
	err := client.PointsSearch(
		context.Background(),
		&request.PointsSearch{
			CollectionName: "test",
			Vector:         []float64{1.1, 2.2, 3.3, 4.4},
			Limit:          10,
			WithPayload:    &withPayload,
			WithVector:     &withPayload,
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
