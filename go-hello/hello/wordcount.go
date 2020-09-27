package main

import (
	"fmt"
	"github.com/goinaction/code/chapter3/words"
	"io/ioutil"
	"os"
)

func main() {

	filename := os.Args[0]

	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)

	fmt.Printf("There are  %d words in your text. \n", count)

}
