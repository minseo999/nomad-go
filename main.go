package main

import (
	"fmt"

	"github.com/minseo999/nomad-go/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "First word"}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	dictionary.Delete(baseWord)
	_, searchError := dictionary.Search(baseWord)
	fmt.Println(searchError)
}
