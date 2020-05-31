package main

import "fmt"

func main()  {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 Coin to Salman")
	bc.AddBlock("Send 4 Coins to Farooq")

	for i, block := range bc.blocks {
		fmt.Printf("Block # %d: \n", i)
		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}

}
