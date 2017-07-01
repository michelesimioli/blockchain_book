## Blocco

Un blocco è composto di:
* testata
* lista transazioni
* lista delle testate degli _Uncle_

La testata ha la struttura:
* Parent hash - del blocco precedente
* Ommers hash - della lista delle testate di Uncles
* State root - hash della radice del Trie di stato
* Transaction root - hash della radice del Trie delle transazioni
* Receipts root - hash della radice del Trie delle ricevute
* Logs bloom - filtro che punta all'indirizzo del logger ed entries di log
* Difficulty - livello corrente di difficoltà
* Number - numero sequenzialedel blocco corrente; il Blocco Genesis è lo zero
* Gas limit - limite di consumo di gas dell'intero blocco
* Gas used - totale del consumo di gas di tutte le transazioni del blocco
* Timestamp - data ed ora di generazione del blocco
* Extra data - altri dati relativi al blocco
* Mixhash - coinvolto nel calcolo dello hash di difficoltà
* Nonce - da aggiustare come in Bitcoin per il calcolo dello hash di difficoltà

