package client

import (
	"context"
	"github.com/baoog/demo-goeth/config"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

const URL = "https://bsc.getblock.io"

func NewClient() *ethclient.Client {
	cl, err := rpc.DialContext(context.Background(), URL)
	if err != nil {
		log.Fatal(err)
	}

	apiKey := config.Get("GETBLOCK_API_KEY", "")

	cl.SetHeader("x-api-key", apiKey)
	cl.SetHeader("Content-Type", "application/json")
	return ethclient.NewClient(cl)
}
