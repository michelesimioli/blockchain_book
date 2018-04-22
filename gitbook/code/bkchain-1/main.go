package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Struttura di base di un blocco
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// Metodo di calcolo dello hash
func (b *Block) SetHash() {
	// occorre convertire l'intero in array di bytes
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// notare l'ordine di join deciso
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// Creazione di un nuovo blocco
// Occorre fornire i dati da maeetervi dentro e lo hash del blocco precedente
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// Struttura per il Blockchain
type Blockchain struct {
	blocks []*Block
}

// Aggiungere un blocco al Blockchain
// Occorre fornire i dati da mettervi dentro
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// Occorre il blocco iniziale
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Creazione di un nuovo blockchain col blocco iniziale
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// Esempio d'uso
func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
