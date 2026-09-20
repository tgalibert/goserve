# Feuille de route : Serveur HTTP/1.1 From Scratch en Go

L'objectif de ce projet est de concevoir et implémenter un serveur HTTP/1.1 complet et fonctionnel sans dépendre du package `net/http` de Go, en partant directement des sockets TCP bruts fournis par le package `net`.

---

## 📌 Vue d'ensemble de l'architecture

```
                      +-----------------------------+
                      |   Client (Navigateur/cURL)  |
                      +-----------------------------+
                                     │
                                     ▼  Connexion TCP
                      +-----------------------------+
                      |      net.Listener           |
                      |  (Accept & Goroutines)      |
                      +-----------------------------+
                                     │
                                     ▼  Flux d'octets brut
                      +-----------------------------+
                      |       Request Parser        |
                      |  - Request Line             |
                      |  - Headers                  |
                      |  - Body (Content-Length)    |
                      +-----------------------------+
                                     │
                                     ▼  Objet Request
                      +-----------------------------+
                      |      Router & Middleware    |
                      |  - Path & Method matching   |
                      |  - CORS & Preflight         |
                      +-----------------------------+
                                     │
                                     ▼  Objet Response
                      +-----------------------------+
                      |      Response Writer        |
                      |  - Status Line              |
                      |  - Headers                  |
                      |  - Body                     |
                      +-----------------------------+
                                     │
                                     ▼  Octets formatés (\r\n)
                      +-----------------------------+
                      |        Socket net.Conn      |
                      +-----------------------------+
```

---

## 🚀 Phases de développement

### Phase 1 : Socle Réseau TCP & Concurrence

_L'objectif est d'établir l'écoute réseau et d'isoler chaque connexion._

- [x] **1.1 Écoute réseau** : Ouvrir un socket TCP sur un port dédié (ex. `:8080`) avec `net.Listen`.
- [x] **1.2 Boucle d'acceptation** : Implémenter une boucle infinie appelant `listener.Accept()` pour recevoir les clients.
- [x] **1.3 Concurrence** : Déléguer le traitement de chaque `net.Conn` à une goroutine dédiée.
- [x] **1.4 Gestion du cycle de vie** : S'assurer de fermer la connexion (`conn.Close()`) en fin de traitement.

---

### Phase 2 : Parser de Requêtes HTTP (RFC 9112)

_L'objectif est de convertir le flux d'octets brut reçu en une structure Go `Request` exploitable._

- [ ] **2.1 Découpage des lignes** : Utiliser un lecteur tamponné (`bufio.Reader`) pour lire ligne par ligne en respectant le délimiteur `\r\n` (CRLF).
- [ ] **2.2 Parsing de la Request Line** :
  - Extraire la **Méthode** (`GET`, `POST`, `OPTIONS`, etc.).
  - Extraire le **Chemin** (`/`, `/api/items`) et séparer les éventuels _Query Parameters_ (`?key=value`).
  - Extraire et valider la **Version du protocole** (`HTTP/1.1`).
- [ ] **2.3 Parsing des Headers** :
  - Lire chaque ligne jusqu'à atteindre la ligne vide (`\r\n`).
  - Séparer chaque header au format `Clé: Valeur`.
  - Normaliser les clés de headers pour gérer l'insensibilité à la casse.
- [ ] **2.4 Lecture du Body** :
  - Vérifier la présence du header `Content-Length`.
  - Si présent, lire **exactement** le nombre d'octets spécifié depuis le socket.
- [ ] **2.5 Modélisation** : Définir la structure `Request` regroupant toutes ces informations.

---

### Phase 3 : Générateur de Réponses HTTP

_L'objectif est de formater et envoyer une réponse conforme aux spécifications._

- [ ] **3.1 Modélisation de la Réponse** : Définir une structure `Response` (statut, headers, body).
- [ ] **3.2 Formatage du protocole** :
  - **Status Line** : `HTTP/1.1 <Code> <Raison>\r\n` (ex: `HTTP/1.1 200 OK\r\n`).
  - **Headers** : Écrire chaque en-tête suivi de `\r\n`.
  - **Séparateur obligatoire** : Écrire une ligne vide supplémentaire (`\r\n`).
  - **Body** : Écrire les données brutes.
- [ ] **3.3 Calcul du `Content-Length`** : Injecter automatiquement la taille exacte du body en octets.

---

### Phase 4 : Gestion des Headers avancés & CORS

_L'objectif est de garantir l'interopérabilité avec les navigateurs et applications clientes._

- [ ] **4.1 Headers standards** :
  - `Content-Type` (ex: `text/plain`, `application/json`, `text/html`).
  - `Date` (format HTTP standard RFC 1123).
  - `Connection` (`close` vs `keep-alive`).
- [ ] **4.2 Gestion de CORS (Cross-Origin Resource Sharing)** :
  - Gérer les requêtes de pré-vérification (**Preflight Requests**) via la méthode `OPTIONS`.
  - Headers à injecter :
    - `Access-Control-Allow-Origin` (ex: `*` ou origines spécifiques).
    - `Access-Control-Allow-Methods` (`GET, POST, PUT, DELETE, OPTIONS`).
    - `Access-Control-Allow-Headers` (`Content-Type, Authorization`).
    - `Access-Control-Max-Age`.
  - Répondre avec un statut `204 No Content` (ou `200 OK`) lors d'un preflight réussi.

---

### Phase 5 : Système de Routage

_L'objectif est d'associer des routes à des fonctions de traitement (handlers)._

- [ ] **5.1 Table de routage** : Créer une structure associant un couple `(Méthode, Chemin)` à un handler.
- [ ] **5.2 Handlers** : Définir une signature type `func(req *Request, res *Response)`.
- [ ] **5.3 Gestion des erreurs HTTP de base** :
  - Renvoyer `404 Not Found` si le chemin n'existe pas.
  - Renvoyer `405 Method Not Allowed` si le chemin existe mais pas avec cette méthode.
  - Renvoyer `400 Bad Request` en cas d'erreur de parsing de la requête.

---

### Phase 6 : Robustesse & Perfectionnement

_L'objectif est de rendre le serveur stable face aux pannes et connexions lentes._

- [ ] **6.1 Timeouts & Deadlines** : Configurer `conn.SetDeadline` ou `conn.SetReadDeadline` pour éviter les connexions fantômes (attaques Slowloris).
- [ ] **6.2 HTTP Keep-Alive** : Permettre la réutilisation d'une connexion TCP pour plusieurs requêtes consécutives tant que `Connection: close` n'est pas demandé.
- [ ] **6.3 Panic Recovery** : Intercepter les panics dans les goroutines de traitement pour éviter le crash complet du serveur.
