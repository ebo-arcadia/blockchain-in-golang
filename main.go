package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Sega")
	bc.AddBlock("Send 2 dodge coins to Jeffrey")
	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}