## Funzioni di Hash

Una funzione di **hash** o di **digest** opera su uno stream di dati di lunghezza qualsiasi e ritorna uno stream di bit di lunghezza nota, detto _hash_ o _segnatura_:

Due stream di input diversi producono _hash_ diversi, due stream uguali producono _hash_ uguali.

Quando due stream di input diversi producono lo stesso _hash_, l'effetto si chiama **collisione**. In una funzione di hash di buona qualità, la collisione è _computazionalmente impossibile_.

La funzione di hash non è invertibile: non è possibile partire dallo hash, con procedimenti deterministici, ottenere lo stream di input originario.

Per ottenere uno hash dato, l'unico procedimento è la **forza bruta**, ovvero provare tutti gli stream di input possibili.

Lo hash in blockchain viene usato in due campi:
* come identificativo univoco dello stream di input che può essere
    * un intero blocco
    * una struttura dati
* come prova di sforzo, **Proof of Work**, poichè è stato ottenuto uno hash predeterminato o con certe proprietà, e questo lo si è potuto far solo con _forza bruta_.

Vi sono numerose funzioni di hash, con due proprietà più o meno desiderabili:
* difficoltà di attacco, tipicamente in relazione alla lunghezza dello hash
* tempo richiesto per l'ottenimento dello hash

Alcuni esempi sono:
* **MD5**
* **SHA1**
* **SHA256**
* **RIPEMD160**

