package main

import (
	"fmt"

	"github.com/woong97/learngo/mydict"
)

// func main() {
// 	dictionary := mydict.Dictionary{"first": "First word"}
// 	word := "hello"
// 	definition := "Greeting"
// 	err := dictionary.Add(word, definition)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	hello, _ := dictionary.Search(word)
// 	fmt.Println("found", word, "definition:", hello)
// 	err2 := dictionary.Add(word, definition)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 	}
// }

// func main() {
// 	dictionary := mydict.Dictionary{}
// 	baseword := "hello"
// 	dictionary.Add(baseword, "First")
// 	err := dictionary.Update(":sdsds", "Second")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	word, _ := dictionary.Search(baseword)
// 	fmt.Println(word)
// }

func main() {
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	dictionary.Add(baseword, "First")
	dictionary.Search(baseword)
	dictionary.Delete(baseword)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
