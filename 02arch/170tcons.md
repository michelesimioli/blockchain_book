## Tipi di Consenso

### Nakamoto Consensus

Consenso ottenuto sull'elezione o autoelezione di un singolo, che decide quale sia il prossimo blocco valido da inserire nel chain.

La velocità di decisione determina la velocità di produzione blocchi, e quindi il flusso di transazioni accettabili per unità di tempo.

#### Proof of Work (PoW)

Consenso basato sulla dimostrazione di aver compiuto molto lavoro computazionale per poter costruire un blocco valido.
Usato attualmente da molte implementazioni, incluso Bitcoin e la versione corrente di Ethereum.
E' un metodo di provato successo e resistente agli attacchi.

Il requisito di capacità computazionale corrisponde purtroppo a vasti consumi di energia elettrica. Il consumo annuale per il solo Bitcoin corrisponde al consumo della Danimarca.

Al momento questo è il singolo metodo che dà più fiducia, ma si stanno alacremente valutando alternative più ecologiche.

#### Proof of Stake (PoS)

Consenso basato sull'investimento in criptovaluta.
Vengono scelti i blocchi da chi possiede più criptovaluta e/o da più tempo, oppure i nodi più _ricchi_ hanno maggior peso di voto nella scelta del blocco accettabile (**Delegated PoS**).

Il metodo è più efficiente del _PoW_ e consuma meno, ma vi sono tutta una serie di situazioni che possono causare un _double spending_ o una **fork** ostile. Le implementazioni che adottano questo metodo introducono anche una serie di controlli per evitare questi effetti.

#### Deposit-based

Il peso di voto nella scelta di un blocco è proporzionale al valore di un deposito in criptovaluta, che può essere perso in caso di comportamenti provatamente scorretti.

#### Proof of Importance

Estensione del _PoS_ con algoritmi che tengono conto non solo degli _asset_ di criptovaluta posseduti, ma anche del _volume di scambio_ effettuato in criptovaluta.

#### Reputation-based

Il peso di voto è influenzato dalla _reputazione_ del votante, ottenuta con meccanismi simili ai _likes_.
Si sospetta che possa essere un metodo _polarizzante_, a feedback positivo, fino potenzialmente a superare il 50% del controllo consenso.

### Byzantine Consensus

Metodi che richiedono molti interscambi di messaggi di voto prima di arrivare ad un accordo.
Teoricamente sopportano un numero più elevato di _Nodi Bizantini_ e sono sistemi più equi, ma soffrono di relativa lentezza di produzione dei blocchi.

#### Practical Byzantine Fault Tolerance (PBFT)

L'algoritmo classico basato su interscambio di messaggi potenzialmente tra nodi qualsiasi.

#### Federated Byzantine Consensus

Anche chiamato **Stellar Consensus Protocol**. Essenzialmente le transazioni sono validate da un sottogruppo di nodi più _fidati_, che determinano il blocco prescelto.

### Nuovi Metodi

Idee nuove basate su modi diversi di concepire lo stesso _consenso_.

#### Proof of Elapsed Time

Basato su hardware particolare, al momento fornito solo da **Intel SGX**, il processore con **Software Guard Extensions**, che crea l'ambiente **Trusted Execution Environment** (TEE).
Al momento **Hyperledger** (soprattutto in variante IBM) sta considerando questa soluzione.

#### Quantum Byzantine Agreement

Protocollo molto teorico ed esso stesso _bizantino_ basato su quantum computer a 4 qubits e che usa proprietà dello _entanglement_ quantistico (uso del Principio di Indeterminazione di Heisenberg). Futuribile, forse.
