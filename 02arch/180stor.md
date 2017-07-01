## Tipi di Storaggio

I dati possono essere _storati_ (conservati) direttamente all'interno dei blocchi del Blockchain, e questa è la soluzione tradizionale di Bitcoin ed Ethereum.

Vi sono due principali problemi di _scalabilità_:
* la quantità di dati cresce col tempo e la diffusione delle tecnologie Blockchain
* ogni nodo _miner_ deve scaricare dalla rete tutto il Blockchain corrente e mantenerlo costantemente aggiornato. Si parla tipicamente al momento di 100 GB per Bitcoin ed Ethereum

La soluzione deve necessariamente andare verso lo storaggio dati esterno al Blockchain, e l'inserimento nel chain di soli riferimenti ai dati veri.

Vi è già esperienza in questo campo acquisita in reti Peer-to-Peer come BitTorrent, Napster, ecc.

### Distributed Hash Tables (DHT)

Quasi tutti i Torrents lo usano. I dati sono distribuiti sui sistemi dei partecipanti.
Problemi:
* I partecipanti non tengono i dati all'infinito, che invece è un requisito del Blockchain. Anche in caso di numerose repliche non è garantita la conservazione di tutti i dati.
* In un determinato istante alcuni dati possono essere inaccessibili perchè i nodi che li detengono sono offline.

### Inter Planetary file System (IPFS)

Risolve i problemi del _DHT_ e si propone in futuro di sostituire interamente il protocollo HTTP.
E' basato su Kademilia DHT per lo storaggio e Mercle DAG (Directed Acyclic Graphs) per le ricerche dati.
Include un modello di versionamento dei dati simile a quello di **Git**.
L'incentivo alla ritenzione dati è stato dal pagamento ai nodi in criptovaluta, come nel sistema **Filecoin**.

### Ethereum Swarm e Whisper

Sono prodotti dell'ecosistema Ethereum e in continua evoluzione. **Swarm** è un sistema cluster per lo storaggio, **Whisper** è un protocollo di comunicazione.

### Maidsafe

Altro progetto che intende sostituire HTTP e formare una rete più efficiente e sicura.
Basato su _crowdsourcing_ e il pagamento servizi di storaggio nella criptovaluta **Safecoin**.

### BigChainDB

Iniziativa Open Source con estensioni Enterprise.
Combina caratteristiche di database distribuiti **NoSQL** con cartteristiche del Blockchain.
Si dice compatibile con **IPFS** e i prodotti **Ethereum**.
Può essere un metodo promettente nel non lontano futuro.
