## Difficoltà

Sono considerati blocchi validi quelli che superano un certo livello di difficoltà. Questa è data da un numero intero e calcolata con l'algoritmo **Ethash**. (`https://github.com/ethereum/wiki/wiki/Ethash`)

L'algoritmo è complesso e l'unica soluzione è data da un approccio di **forza bruta**. _Ethash_ è concepito per favorire grandi quantità di RAM, anzichè potenza di CPU. Questo consente l'uso di GPU anzichè di ASIC.
Le piattaforme di riferimento per i miners sono:
* Ubuntu Linux ultima versione
* EthOS - un linux modificato per il mining di Ethereum (`http://ethosdistro.com/`)

La difficoltà varia ad ogni blocco nuovo, nel tentativo di mantenere costante il rateo di generazione blocchi a circa 15 secondi.
La difficoltà è calcolata dall'espressione:

```
block_diff = parent_diff + parent_diff / 2048 *  
max(1 - (block_timestamp - parent_timestamp) / 10, -99) +  
int(2**((block.number / 100000) - 2))
```

La difficoltà aumenta se il rateo è inferiore a 10 secondi, diminuisce se superiore a 20 secondi.

La difficoltà inoltre aumenta esponenzialmente ogni 100.000 blocchi.

### Time Bomb

Anche detta **Ice Age**. L'incremento esponenziale di difficoltà è voluto, per costringere i miners col tempo a passare dal consenso **Proof of Work** al futuro consenso **Proof of Stake**, che sarà implementato del protocollo **Casper** e finalizzato nella release **Serenity**.
