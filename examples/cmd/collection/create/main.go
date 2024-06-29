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

	onDisk := true
	resp := &response.CollectionCreate{}
	err := client.CollectionCreate(
		context.Background(),
		&request.CollectionCreate{
			CollectionName: "test",
			Vectors: request.VectorsParams{
				Size:     4,
				Distance: request.DistanceCosine,
				OnDisk:   &onDisk,
			},
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
