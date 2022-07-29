package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	infuraURL := "https://mainnet.infura.io/v3/31eeea37900149ad9434747761bed71f"

	// Connection to the blockchain via infura node

	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal("Whoops something went wrong!", err)
	}
	defer client.Close()

	// Get and print transaction details

	ctx := context.Background()
	tx, pending, _ := client.TransactionByHash(ctx, common.HexToHash("0x30999361906753dbf60f39b32d3c8fadeb07d2c0f1188a32ba1849daac0385a8"))
	if !pending {
		fmt.Println(tx)
	}

	// get and print last block number
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("error to get a block:%v", err)
	}

	fmt.Println(block.Number())
}
