package main

//importation nécéssaire
import (
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fichier, _ := os.ReadFile("Hangman.txt") //on lis le fichier txt grâce au package os
	var liste_de_mots []string               //list qui servira à stocker la liste de mots
	var mot string                           //variable qui sert pour identifier et enregistrer les mots dans la liste
	var resultat string                      //variable de résultat final

	//enregistrement des mots du fichier txt dans une liste
	for index, caractère := range fichier {
		if caractère == 13 { //lorsque l'on va à la ligne le mot est fini
			liste_de_mots = append(liste_de_mots, mot) //on l'append
			mot = ""                                   //on rénitialise mot
		} else if caractère != 10 { //sinon on ajoute le caractère a mot
			mot += string(caractère)
		}
		if index == len(fichier)-1 { //on vérifie la fin
			liste_de_mots = append(liste_de_mots, mot)
		}
	}

	//premier fragement
	resultat = liste_de_mots[0]

	//deuxieme fragment
	resultat += " " + liste_de_mots[len(liste_de_mots)-1]

	//troisieme fragment
	for index, mot := range liste_de_mots { //parcours la liste de mots
		if mot == "before" { //on chercher before
			direction, _ := strconv.Atoi(liste_de_mots[index+1]) //on attribue a direction le chiffre en int, du mot après before
			resultat += " " + liste_de_mots[direction-1]         //on ajoute au résultat
		}
	}

	//quatrieme fragment
	for index, mot := range liste_de_mots { //parcours la liste de mots
		if strings.Contains(mot, "now") { //si le mot actuel contients les caractère "now"
			mot2 := liste_de_mots[index-1]                                     //mot temporaire est le mot précédent now
			resultat += " " + liste_de_mots[int(mot2[1])/len(liste_de_mots)-1] //on ajoute au résultat le mot à l'indice - 1 de la division entre la valeur ascii du deuxième caractère de mot2 et le nombre total de mot
		}
	}

	//cinquième fragement
	rand.Seed(int64(os.Getpid()))                            //on regénère la seed en se basant sur l'identifiant du processus actuel go
	println(resultat, " : ", rand.Intn(1000000000000000000)) //on génère un nombre aléatoire
}
