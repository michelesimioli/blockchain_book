## Organizzazioni Decentralizzate

### Tipi Fondamentali

Una **Decentralized Organization** (DO) è un programma software che:
* esegue in un Blockchain
* è composto di Smart Contracts
* modella l'organizzazione di un ente umano, con ruoli e protocolli

A questo livello richiedono ancora interazione umana - input, decisioni.

Una **Decentralized Autonomous Organization** (DAO) è una DO organizzata per procedere autonomamente nell'espletamento delle attività programmate, senza alcun intervento umano.

Vi sono comunque persone coinvolte nel funzionamento di un DAO:
* Curator - per la manutenzione del codice
* Contractor - personale d'aiuto tipicamente esterno e temporaneo

### Esempio

Il progetto più famoso in questo senso è stato _il_ **DAO**, che è durato daal 30/4 al 17/6/2016.

Era un centro di gestione fondi VC (_Venture Capital_) e _Crowdsourcing_, basato su Ethereum e sulla criptovaluta Ether.

Il DAO era basato su una serie di Smart Contracts scritti nel linguaggio **Solidity** di Ethereum, forniti come templati e Open Source, come _framework_ per la gestione di specifiche _DAOs_ profit e no-profit.

Costituiva veramente uno shift paradigmatico per il VC:
* gli investitori non avevano il diritto di recupero dei fondi, ma una quota di voto nel finanziamento di progetti innovativi
* i profitti dei progetti finanziati sarebbero ritornati agli investitori
* gli investitori potevano essere ovunque nel mondo e non dovevano identificarsi umanamente
* la governance tecnica era completamente decentralizzata
* l'interfaccia umana era una ditta registrata come SRL in Svizzera, collegata ad un _Exchange_ per l'aquisto di ETH.


In pochi mesi il DAO aveva raccolto capitali per un equivalente di milioni di dollari.

Sfortunatamente la presenza di bachi nel software ha permesso agli hacker di sottrarre una vasta quantità di fondi cripto accumulati e smistarli in altri account.

A questo punto il progetto DAO è stato effettivamente sospeso.

Come conseguenza tecnica Ethereum ha subito uno **hard fork**, dovuto alle patch estensive al codice installato, con conseguente ulteriori perdite di criptovaluta, stimate sui $50 milioni.

Parte della comunità Ethereum si è rifiutata di transire al nuovo software e ha mantenuto il vecchio blockchain con il nome **Ethereum Classic** e nuova criptovaluta `ETC`.

`ETC` non è `ETH` ed è disconosciuto dall'organizzazione Ethereum ufficiale:
* non riceve le innovazioni architettoniche di Ethereum
* introduce proprie innovazioni indipendenti
* ha successo di mercato in costante calo

#### Lessons Learnt

* Gli Smart Contracts, come ogni software, sono soggetti a **Vulnerabilità** ed **Esposizioni**, che possono permettere attacchi hacker
* In particolare: **Solidity** ha perso fiducia. Vi sono commenti che un linguaggio per Smart Contract dovrebbe essere **Demonstrably Secure**, non solo **Turing Complete**
* Il rischio relativo deve essere accuratamente valutato e mitigato prima di dar vita ad un **DAO**
* Uno **hard fork** obbliga
    * o a perdita di criptovaluta
    * o a generazione di una nuova valuta indipendente, con perdita di CAP aggregato per entrambi i _branch_; alla fine probabilmente uno dei branch soccombe
* Probabilmente in generale le innovazioni tecnologiche radicali non saranno senza spargimento di sangue

#### Necessità per un altro DAO globale

Un DAO non ha status legale in nessuna nazione al mondo, infatti viene a volte contrastato.

Non è definibile la giurisdizione legale di un DAO o il _forum_ di competenza per le dispute.

Si propone di introdurre un **Autonomous Agent**, un particolare applicativo distribuito, a sua volta basato sugli Smart Contracts, col compito specifico di monitorare le _compliance_ degli altri DAO.

### Sottoinsiemi di DAO

**Decentralized Autonomous Corporations** (DAC) - che si mappano su corporazioni (_profit_) veramente esistenti e con identità giuridica.
Un DAC sarebbe finanziato con normale emissione di partecipazioni e avrebbe uno schema di pagamento dividendi.

**Decentralized Autonomous Societies** (DAS) - che si mappano su organizzazioni (_no-profit_) esistenti: Stato, NGO, ecc.
Specie a livello di Pubblica Amministrazione un DAS può teoricamente aumentare l'efficienza, ridurre gli sprechi, evitare la corruzione.
Tecnicamente è fattibile, ma politicamente?
