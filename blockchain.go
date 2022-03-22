package main

import (
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func newBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp       %d\n", b.timestamp)
	fmt.Printf("timestamp       %d\n", b.nonce)
	fmt.Printf("previous_hash   %s\n", b.previousHash)
	fmt.Printf("transactions    %s\n", b.transactions)
}

func main() {
	b := newBlock(0, "init hash")
	fmt.Println(b)
	b.Print()
}
