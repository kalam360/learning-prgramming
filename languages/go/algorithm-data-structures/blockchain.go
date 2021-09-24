package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	timestamp time.Time
	transactions []string
	prevHash []byte
	Hash []byte
}

func NewHash(time time.Time, transactions []string, prevHash []byte) []byte {
	input := append(prevHash, time.String()...)

	for transaction := range transactions {
		input = append(input, string(rune(transaction)...))

	}

	hash := sha256.Sum256(input)

	return hash[:]

}

func NewBlock(transactions []string, prevHash []byte) *Block {
	currentTime := time.Now()

	return &Block {
		timestamp: currentTime,
		transactions: transactions,
		prevHash: prevHash,
		Hash: NewHash(currentTime, transactions, prevHash),
	}

}