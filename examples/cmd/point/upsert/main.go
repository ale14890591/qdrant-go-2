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

	wait := true
	resp := &response.PointsUpsert{}
	err := client.PointsUpsert(
		context.Background(),
		&request.PointsUpsert{
			CollectionName: "test",
			Wait:           &wait,
			Points: []request.Point{
				{
					ID:     "45b07125-f592-414f-a9d0-160c8ecc283a",
					Vector: []float32{1.1, 2.2, 3.3, 4.4},
					Payload: map[string]interface{}{
						"key": "value",
					},
				},
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
