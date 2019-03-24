// Package structure contenant la structure principale de Goshield.
package structure

// ENCRYPT représente la constante quand l'utilisateur veut chiffrer.
const ENCRYPT int = 1

// DECRYPT représente la constante quand l'utilisateur veut déchiffrer. 
const DECRYPT int = 2

// Documents représente l'ensemble des fichiers/dossiers à (dé)chiffrer avec un mot de passe. 
type Documents struct {
	// Mode ENCRYPT (chiffrement) ou DECRYPT (dechiffrement).
	Mode int

	// Password représente le mot de passe choisi par l'utilisateur. 
	Password string

	// Doc représente l'ensemble des chemins des fichiers/dossiers à (dé)chiffrer.
	Doc []string
}