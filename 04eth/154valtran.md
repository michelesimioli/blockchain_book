## Validazione

E' compiuta dai nodi _miner_ durante la costruzione di un nuovo blocco.

### Validazione delle Transazioni

Una transazione è valida se:
* ben formata in codifica RLP
* la firma della transazione è valida
* il nonce della transazione è uguale a quello dell'emittente
* non ha ecceduto i limiti di gas
* il mittente ha fondi per coprire la commissione

#### Sottostato di Transazione

Creato durante la transazione e processato al suo termine.
E' composto da:
* Suicide Set - lista degli account da eliminare al termine della transazione
* Log Series - lista di eventi da comunicare ad applicativi esterni
* Refund Balance - saldo del gas che deve essere reso

### Validazione del Blocco

La validazione di un blocco è composta da:
* Validazione del Blocco Precedente
* Validazione degli _Uncles_ - devono essere blocchi validi a loro volta
    * a profondità non superiore a 6 dal blocco corrente
    * non più di 2
* Validazione delle Transazioni
* Validazione dello Stato
* Validazione del Timestamp - successivo a quello del blocco precedente
    * Viene tollerato fino a 15 minuti nel futuro
* Validazione del Nonce - compatibile con la difficoltà

### Ricompensa del miner

Ogni miner riceve la ricompensa corrente per la generazione di un nuovo blocco, al momento di 5 Eth.

Inoltre riceve il gas di commissione per le transazioni incluse nel blocco.

I miners dei blocchi _Uncle_ ricevono i 7/8 della ricompensa.

Anche i miners dei blocchi _Stale_ ricevono una ricompensa, di 1/32 del valore di quella del blocco accettato.

