package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("Hangman.txt")
	var list_of_words []string
	var word string
	var result string
	//enregistrement des mots du fichier txt dans une liste
	for index, char := range data {
		if char == 13 {
			list_of_words = append(list_of_words, word)
			word = ""
		} else if char != 10 {
			word += string(char)
		}
		if index == len(data)-1 {
			list_of_words = append(list_of_words, word)
		}
	}
	//premier fragement
	result = list_of_words[0]
	//deuxieme fragment
	result += " " + list_of_words[len(list_of_words)-1]
	//troisieme fragment
	for index, word := range list_of_words {
		if word == "before" {
			direction, _ := strconv.Atoi(list_of_words[index+1])
			result += " " + list_of_words[direction-1]
		}
	}
	//quatrieme fragment
	for index, word := range list_of_words {
		if strings.Contains(word, "now") {
			word2 := list_of_words[index-1]
			result += " " + list_of_words[int(word2[1])/len(list_of_words)-1]
		}
	}
	//génération d'un nombre aléatoire avec une seed différente à chaque fois
	rand.Seed(int64(os.Getpid()))
	println(result, " : ", rand.Intn(1000000000000000000))
}
