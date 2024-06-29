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

	resp := &response.CollectionDelete{}
	err := client.CollectionDelete(
		context.Background(),
		&request.CollectionDelete{
			CollectionName: "test",
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
