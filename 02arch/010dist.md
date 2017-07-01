## Sistemi Distribuiti

Quello dei sistemi distribuiti è un paradigma computazionale ormai molto affermato.
Un sistema distribuito è costituito da _nodi_ di rete che si interscambiano _messaggi.

Sono possibili tre topologie di rete:

![Centralizzata](../gitbook/images/centralized1.png)

1. **Centralizzata**
Un **Server** centrale interagisce con _n_ **Clients**.
Isolata e non resiliente: punto singolo di fallimento.

![Decentralizzata](../gitbook/images/decentr.png)

2. **Decentralizzata**
Più reti centralizzate sono interconnesse tra loro. Esempio: _Internet_ moderna.
Alquanto resiliente: possibilità di _sezionamento_ in sottoreti irraggiungibili.

![Distribuita](../gitbook/images/distribuite.png)

3. **Distribuita**
Il numero di interconnessioni è molto elevato, almeno 3 per ogni nodo. Internet futura.
Molto resilente.
