package main

import (
	"fmt"
	"goblockchain/block"
	"goblockchain/wallet"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	walletM := Wallet.NewWallet()
	walletA := Wallet.NewWallet()
	walletB := Wallet.NewWallet()

	t := wallet.NewTransaction(
		walletA.PrivateKey(),
		walletB.PublicKey(),
		walletA.BlockchainAddress(),
		walletB.BlockchainAddress(),
		1.0)

	blockchain := block.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), walletA.PublicKey, t.GenerateSignature())

	fmt.Printf("Added? ", isAdded)

	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))
}
