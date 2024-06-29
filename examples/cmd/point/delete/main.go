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
	resp := &response.PointsDelete{}
	err := client.PointsDelete(
		context.Background(),
		&request.PointsDelete{
			CollectionName: "test",
			Wait:           &wait,
			Filter: request.Filter{
				Must: []request.M{
					{
						"key": "key",
						"match": request.M{
							"value": "value",
						},
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
