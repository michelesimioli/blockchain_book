package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

// Difficoltà del mining
// Il tempo impiegato aumenta considerevolmente con la difficoltà
const targetBits = 20

// Struttura di base di un blocco
// Aggiunta di nonce per Proof Of Work
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
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
// Modificato con l'aggiunta del Proof Of Work
func NewBlock(data string, prevBlockHash []byte) *Block {
	// 0 è il valore del nonce
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

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
// Modificato per validare il Proof Of Work
func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

// Struttura per il POW
type ProofOfWork struct {
	block  *Block
	target *big.Int // per evitare overflow
}

// Generazione del POW
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// targetBits è la difficoltà corrente
	// Lsh è la funzione di Left Shift
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// preparazione dei dati per lo hash
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Algoritmo per il POW
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	// Massimo nonce per impedire un overflow matematico
	maxNonce := math.MaxInt64

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// Utility per convertire un Int in Hex
func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}

// Validazione del Proof Of Work
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
