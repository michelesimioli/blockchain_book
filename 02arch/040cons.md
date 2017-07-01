## Consenso

![Consensus](../gitbook/images/consensus.png)

Accordo su un unico valore di dati.
Relativamente facile raggiungere il consenso in un sistema centralizzato, molto più complesso in un sistema distribuito.

Richiede normalmente un interscambio di numerosi messaggi tra i nodi partecipanti.

Requisiti del consenso:
* **Accordo** - tutti i nodi _onesti_ hanno lo stesso valore
* **Terminazione** - il procedimento di intercambio messaggi ha una durata finita (e più breve possibile)
* **Validità** - il valore finale è il valore iniziale proposto da almeno uno dei nodi onesti
* **Resilienza** - è tollerata la presenza di un certo numero di Nodi Bizantini
* **Integrità** - nessun nodo onesto 'cambia idea' durante il procedimento di consenso

