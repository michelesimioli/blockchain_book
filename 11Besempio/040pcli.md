# Persistenza e CLI

Intendiamo in questa iterazione aggiungere due _features_:

1. Lo storaggio dei dati su file system in modo persistente
1. L'interazione tramite sottocomandi da interfaccia **CLI** (_Command Line Interface_)

## Database

Il Blockchain è indipendente da particolari databases. Non servono in realtà le funzionalità di un Database Relazionale, ma solo:

1. Inserire un record come coppia (`chiave`, `valore`)
1. Recuperare il `valore` di una certa `chiave`
1. Cancellare un record tramite `chiave`

Il progetto _Bitcoin Core_ usa **LevelDB**, scritto in C++.

Noi useremo **BoltDB**, scritto in Go, che presenta le stesse funzionalità.
Le sue proprietà principali sono:

1. Tutti i dati sono arrays di bytes, non vi sono tipi
1. Le coppie _chiave-valore_ sono storati in **buckets**, che raggruppano coppie simili (funzionalmente simili a _tabelle_): per accedere a un record si usano il `bucket` e la `chiave`

Il package Go per BoltDB deve venire importato nel programma con:

```go
import "github.com/boltdb/bolt"
```

e deve anche essere scaricato dalla rete con il comando da terminale:

```bash
go get github.com/boltdb/bolt
```

### Design del Database

Il design del database è lo stesso di quello di Bitcoin e prevede la gestione di transazioni, che non sono considerate in questa iterazione di esempio.

Creiamo due _buckets_:

* `blocks` per i blocchi di un blockchain
* `chainstate` per lo stato del blockchain, incluso alcuni metadati

In **`blocks`**, le coppie `key -> value` sono:

* `'b' + 32-byte block hash` -> block index record
* `'f' + 4-byte file number` -> file information record
* `'l'` -> 4-byte file number: l'ultimo numero di file di blocco usato
* `'R'` -> 1-byte boolean: se è in corso il processo di reindicizzazione
* `'F' + 1-byte flag name length + flag name string` -> 1 byte boolean: vari flag che possono essere on o off
* `'t' + 32-byte transaction hash` -> transaction index record

In **`chainstate`**, le coppie `key -> value` sono:

* `'c' + 32-byte transaction hash` -> _unspent transaction output record_ (**UTXO**) per quella transazione
* `'B'` -> 32-byte block hash: lo hash fino al quale il database rappresenta gli _unspent transaction outputs_

In questa iterazione useremo solo il bucket `blocks` e avremo solo due tipi di coppie:

* `32-byte block-hash` -> Block structure (serializzata)
* `'l'` -> lo hash dell'ultimo blocco del blockchain

### Serializzazione

Per la serializzazione e deserializzazione dei dati usiamo il protocollo **gob** nativo di Go.

```go
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
```

### Struttura del Blockchain

Un blockchain non è più un array di blocchi mantenuto in memoria, ma un database BoltDB su file. Quindi ha bisogno del riferimento a tale file, che chiameremo `db`.
Inoltre ha bisogno dello hash dell'ultimo blocco, da cui è possibile risalire a tutti i blocchi precedenti, che chiameremo `tip`.

```go
// Struttura per il Blockchain
// Modificato per supporto a BoltDB
type Blockchain struct {
  tip []byte
  db  *bolt.DB
}
```

### Creazione del Blockchain

Questo metodo ora deve implementare la persistenza sul DB. Compie i seguenti passi:

* Apre il file del DB
* Controlla se contiene un blockchain
* Se c'è un blockchain nel file:
  * crea un oggetto istanza di `Blockchain`
  * punta il campo `db` al bucket del database
  * carica il `tip` con lo hash dell'ultimo blocco, descritto dal database
* Se nel file non c'è un blockchain
  * crea il blocco Genesi
  * lo pone nel database
  * salva lo hash del blocco Genesi come hash dell'ultimo blocco
  * crea un'istanza di `Blockchain` con queste caratteristiche

```go
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
```

### Aggiunta di Blocco

Il metodo di aggiunta di un blocco al blockchain diventa ora più complesso, poichè non è più l'aggiunta di un elemento ad un array, ma di una coppia di chiavi al database.

```go
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
```
