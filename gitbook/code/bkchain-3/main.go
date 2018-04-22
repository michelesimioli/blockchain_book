package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

/* Blocco */

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

/* Blockchain */

// Struttura per il Blockchain
// Modificato per supporto a BoltDB
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

// Iteratore su Blockchain
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Creazione di Iteratore
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// Ritorna il prossimo blocco tramite l'iteratore
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}

// Aggiungere un blocco al Blockchain
// Occorre fornire i dati da mettervi dentro
// Modificata per supporto a BoltDB
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte // hash dell'ultimo blocco

	// Transazione di tipo READ per leggere un record
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket)) // TODO: controllo errori
		lastHash = b.Get([]byte("l"))

		return nil
	})

	// Per sicurezza, se la transazione è fallita
	if err != nil {
		log.Panic(err)
	}

	// Creazione di un nuovo blocco
	newBlock := NewBlock(data, lastHash)

	// Transazione WRITE di aggiunta del nuovo blocco
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket)) // TODO: se b è nil
		// Inserimento coppia chiavi
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		// Aggiorna il riferimento allo hash dell'ultimo blocco
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		// Anche al nostro blockchain in memoria
		bc.tip = newBlock.Hash

		return nil
	})
}

// Occorre il blocco iniziale
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// File per il database di BoltDB
const dbFile = "bkchain.dat"

// Il bucket dei blocchi
const blocksBucket = "blocks"

// Creazione di un nuovo blockchain col blocco iniziale
// Modificata per il supporto a BoltDB
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	// Questa è una transazione WRITE
	// Update di Bolt ha come argomento una funzione anonima
	err = db.Update(func(tx *bolt.Tx) error {
		// Apre il bucket 'blocks'
		b := tx.Bucket([]byte(blocksBucket))

		// Se non c'è
		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")
			genesis := NewGenesisBlock() // crea il blocco Genesi

			// Crea il bucket 'blocks'
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			// Vi pone la coppia di chiavi del blocco genesi
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			// Registra lo hash dell'ultimo blocco
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			// registra il 'tip' del blocco Genesi
			tip = genesis.Hash
		} else { // se il bucket 'blocks' esiste
			// ottiene il 'tip' dal bucket
			tip = b.Get([]byte("l"))
		}

		return nil // se è qui non vi sono stati errori
	}) // Qui finisce lo Update

	// Per sicurezza, ma abbiamo gestito ogni errore fin'ora
	if err != nil {
		log.Panic(err)
	}

	// Costruzione di un oggetto Blockchain
	bc := Blockchain{tip, db}

	return &bc
}

/* Main */

// Esempio d'uso
// Modificato per l'uso del CLI
func main() {
	bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}

/* Proof Of Work */

// Difficoltà del mining
// Il tempo impiegato aumenta considerevolmente con la difficoltà
const targetBits = 20

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

/* Supporto a BoltDB */

// Serializzazione
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		panic("Cannot encode")
	}

	return result.Bytes()
}

// Deserializzazione
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		panic("Cannot decode")
	}

	return &block
}

/* CLI */

// Command Line Interface
type CLI struct {
	bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// Parse della linea di comando
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
