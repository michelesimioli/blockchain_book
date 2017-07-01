## Transazioni

Una transazione è un messaggio firmato incluso in un blocco, e può essere di due tipi:
* **Message Call** - semplice interscambio di informazioni
* **Contract Creation** - per un nuovo contratto, anche transazione puramente monetaria

Una transazione ha i campi:
* **Nonce** - identificativo incrementato ad ogni transazione
* **Gas Price** - costo della transazione in Wei.
* **Gas Limit** - massimo numero di Wei che il mittente è disposto a pagare
* **To** - destinatario
* **Value** - valore trasferito
* **Signature** - firma crittografica della transazione

Tutte le transazioni sono raccolte in un _Patricia Trie_, e lo hash della sua radice è un campo della testata del blocco corrente.

