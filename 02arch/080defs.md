## Definizioni

### Blockchain

1. Un Blockchain è un meccanismo di consenso decentralizzato.
2. Un Blockchain è un registro distribuito di transazioni, raggruppate in blocchi.
3. Un Blockchai è una lista collegata che usa hash anzichè puntatori.

### Indirizzi

Identificativi univoci che denotano gli originanti e i riceventi in una transazione.
Sono solitamente delle chiavi pubbliche o derivati.
Non corrispondono ad un utente fisico: ogni utente può avere più indirizzi, ma un indirizzo non può corrispondere a più utenti.
E' suggerito al limite che un utente generi un nuovo indirizzo per ogni transazione, se vuole evitare correlazioni e mantenere l'anonimato.
Non è affatto detto che un utente fisico debba esistere; gli indirizzi possono riferirsi ad entità software.

### Transazioni

Una transazione è l'unità fondamentale di un Blockchain e rappresenta il trasferimento di valore tra due indirizzi.
Anche quando non vi è una transazione esplicitamente di _bonifico_, esistono quasi sempre dei costi di transazione, commissioni, ecc.

### Blocchi

Un blocco è composto da una o solitamente più transazioni, più campi di gestione.

### Peer-to-Peer

E' una rete di pari, senza differenze gerarchiche a livello software di rete. Tutti i nodi possono, in linea di principio, inviare e ricevere messaggi.
Se vi sono limitazioni d'accesso, sono configurate a livelli software superiori.

### Linguaggio di Scripting

Una transazione può essere corredata da una procedura eseguibile, scritta in un opportuno linguaggio che dipende dall'implementazione. Questo aggiunge un livello di _intelligenza_ alla transazione,descrivendo condizioni e vincoli alla sua esecuzione.
Vi sono molti linguaggi di scripting in varie implementazioni, ma non esistono ancora standard.

### Turing Complete

Un linguaggio di programmazione è detto _Turing Complete_ se può emulare in tutti i suoi aspetti una _Macchina di Turing_, un nastro infinito con testina di lettura e scrittura e movimento bidirezionale. Si tace di solito che qualsiasi computer fisico ha memoria finita.
In pratica un linguaggio imperativo con costrutti decisionali e di ciclo viene considerato Turing Complete.
Non esiste correlazione tra la _sicurezza_ di un linguaggio e la sua _completezza di Turing_. La sicurezza, formale o meno, è ad un livello più alto.

### Macchina Virtuale

E' una estensione alla struttura di un Blockchain che permette l'esecuzione di procedure _Turing Complete_ presenti nelle transazioni.
Evidentemente sono i nodi che gestiscono le transazioni, non il Blockchain stesso, che eseguono le procedure.
Le macchine virtuali sono fortemente dipendenti dall'implementazione e non esistono ancora standard affermati o tentativi.
Nella giusta ottica, si può adottare l'astrazione che il Blockchain è una singola Macchina Virtuale distribuita.

### Macchina a Stati Finiti

Una transazione compie il passaggio da uno stato iniziale ad uno finale, ed è rappresentabile nell'ambito della teoria informatica degli _Automi a Stati Finiti_.

### Nonce

E' un campo di un blocco che contiene un valore simil-casuale.
L'intero hash del blocco, incluso il _nonce_ ha certè proprietà definite a priori, per esempio deve essere inferiore ad un certo valore. Con certi schemi di consenso, chi genera il blocco deve inventare un nonce che produca queste proprietà. Dato che lo hash non è invertibile, il _nonce_ può essere solo indovinato. Questo implica numerosi tentativi e capacità computazionale notevole.

### Smart Contract

Una particolare procedura inglobata in una transazione ed eseguita dalla Macchina Virtuale, con _Logica Business_ che risolve determinati problemi.
Visto che un Blockchain non è modificabile, anche uno _Smart Contract_ non lo è.
Dipende fortemente dall'implementazione e linguaggio di scripting e non è presente in tutti i tipi di Blockchain.
