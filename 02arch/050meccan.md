## Meccanismi di Consenso

Due tipologie:

##### Byzantine Fault Tolerance

* Nessun nodo è prominente
* Interscambio di un certo numero di messaggi firmati, con regole prestabilite di selezione del _migliore_, fino ad un _quorum_ minimo.
* E' predefinita una _qualità_ che determina il migliore.

##### Basati su Leaders

* Un nodo decide. Il nodo può essere eletto.
* Ruoli diversi dei nodi:
    * _Proposer_, _Acceptor_, _Learner_ (Paxos)
    * _Follower_, _Candidate_, _Leader_ (RAFT)

Possono esserci sistemi ibridi: rari, perchè diventano presto complessi.
