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

	isTrue := true
	resp := &response.PointsGet{}
	err := client.PointsGet(
		context.Background(),
		&request.PointsGet{
			CollectionName: "test",
			IDs:            []string{"45b07125-f592-414f-a9d0-160c8ecc283a"},
			WithVector:     &isTrue,
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)
}
