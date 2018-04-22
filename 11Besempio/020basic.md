# Prototipo di Base

## Il singolo Blocco

Partiamo da una visione semplificata di un blocco, definita da:

```go
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}
```

Ove:

* `Timestamp` è l'istante di creazione del blocco, ricavato dal sistema operativo
* `Data` - sono i dati che vengono trasportati dal blocco, saranno le transazioni registrate
* `PrevBlockHash` - è lo hash del blocco precedente
* `Hash` - è lo hash del blocco corrente

Per calcolare lo hash di un blocco viene usato il metodo:

```go
func (b *Block) SetHash() {
  // Il timestamp deve venir convertito da intero ad array di bytes
  timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
  // Vengono uniti i campi della testata
  headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
  // Viene calcolato lo hash
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}
```

La funzione di hash selezionata è **SHA256**, supportata dalla libreria del Go.

Il costruttore del blocco è:

```go
func NewBlock(data string, prevBlockHash []byte) *Block {
  // Creazione dell'oggeto blocco
  block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
  // Calcol dello hash
	block.SetHash()
	return block
}
```

## Il Blockchain

La catena di blocchi si può implementare in Go in due modi: o come array oppure come mappa. Per il momento la implementiamo come array di riferimenti. Dato che un array è ordinato, risulta più facile recuperare il blocco precedente.

```go
type Blockchain struct {
	blocks []*Block
}
```

Il metodo per aggiungere un blocco è:

```go
func (bc *Blockchain) AddBlock(data string) {
  // Recupero del blocco precedente
  prevBlock := bc.blocks[len(bc.blocks)-1]
  // Generazione di un nuovo blocco
  newBlock := NewBlock(data, prevBlock.Hash)
  // Aggiunta al chain
	bc.blocks = append(bc.blocks, newBlock)
}
```

Occorre naturalmente il primo blocco, detto il **Blocco Genesi**. Viene qui creato molto semplicemente con un contenuto predeterminato:

```go
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
```

Ora per creare l'intero blockchain basta la funzione:

```go
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
```

## Programma di Test

La funzione _main_ testa il nostro progresso fin qui:

```go
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
```

## Organizzazione Pratica

Creare una directory di progetto, p.es. `~/go/src/bkchain-1`. Tracciarla con Git.

Creare il file `main.go` e porvi dentro tutti i segmenti di codice indicati sopra. Dichiarare all'inizio `package main`; completare con tutti gli import necessari.

Il risultato è il seguente:

[import lang:"golang"](../gitbook/code/bkchain-1/main.go)

L'esecuzione con:

```text
go run main.go
```

fornisce l'output:

```text
Prev. hash: 
Data: Genesis Block
Hash: 5bc80cfd521c323a0a9983c02e1448e14eb3dd9db984fcd72bfcdd519ad33203

Prev. hash: 5bc80cfd521c323a0a9983c02e1448e14eb3dd9db984fcd72bfcdd519ad33203
Data: Send 1 BTC to Ivan
Hash: 38f78a37e0406cd6c2c02bfc452cf78a90e5a34c9576aabcdbc254d1e23d351c

Prev. hash: 38f78a37e0406cd6c2c02bfc452cf78a90e5a34c9576aabcdbc254d1e23d351c
Data: Send 2 more BTC to Ivan
Hash: d81622053d2221350ed46d717c663f4aec2202bcef033235e02edede24d7badf

```
