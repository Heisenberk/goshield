# GoShield

## Sommaire

1. [Manuel de l'utilisateur](#utilisateur)
	1. [Installation](#installation)
	2. [Compilation et première exécution](#compilation)
	3. [Tests unitaires](#tests)
	4. [Documentation](#doc)
	5. [Exécution](#execution)
2. [Manuel technique](#technique)
3. [Solution explicative](#solution)

## Manuel de l'utilisateur <a id="utilisateur"></a> 

### Installation <a id="installation"></a> 

Pour télécharger l'application, deux méthodes sont disponibles :

- Taper la commande `go get github.com/Heisenberk/goshield` 

<img src="report/pictures/go_get.PNG" alt="Commande go get"/>

- Taper la commande `git clone https://www.github.com/Heisenberk/goshield` ou télécharger le code source directement sur Github 

<img src="report/pictures/github.PNG" alt="Github"/>

Dans le dossier `$(GOPATH)`, le package `github.com/Heisenberk/goshield` sera installé. 

<img src="report/pictures/content.PNG" alt="Contenu du package goshield"/>

Un Makefile y est intégré pour faciliter les tâches pour le développeur/utilisateur. 

- Exécuter la commande `make install` pour compiler et exécuter le programme en ligne de commande avec le mot clé `goshield`. 

### Compilation et première exécution <a id="compilation"></a> 

Pour compiler et générer ainsi un exécutable manuellement, taper la commande `make compil && ./goshield` ou tout simplement `make all`. La liste des commandes seront affichées : 

<img src="report/pictures/make_all.PNG" alt="make all"/>

### Tests unitaires <a id="tests"></a> 

Pour exécuter les tests unitaires, écrire `make test`, qui testeront les sous-packages `goshield/crypto` et `goshield/command`. 

<img src="report/pictures/make_test.PNG" alt="Commande go test"/>

### Documentation <a id="doc"></a> 

Taper dans le terminal la commande `godoc -http=:8080`, puis, dans le même temps, visiter l'adresse sur un navigateur `http://localhost:8080/pkg/github.com/Heisenberk/goshield`. 

<img src="report/pictures/godoc.PNG" alt="Commande godoc"/>

<img src="report/pictures/doc_firefox.PNG" alt="godoc Firefox"/>

### Exécution <a id="execution"></a> 

- Pour afficher les commandes : 

<img src="report/pictures/goshield.PNG" alt="goshield"/>

- Pour chiffrer avec le mot de passe "soutenance" la suite de fichiers/dossiers suivants : 

<img src="report/pictures/encrypt.PNG" alt="encrypt"/>

- Pour déchiffrer avec le mot de passe "soutenance" la suite de fichiers/dossiers suivants : 

<img src="report/pictures/decrypt.PNG" alt="decrypt"/>