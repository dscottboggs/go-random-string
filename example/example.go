package main

import (
	"fmt"
	"log"

	"github.com/dscottboggs/go-random-string"
)

func main() {
	fmt.Println(random.AlphanumericString(20)) //=> [twenty random alphanumeric characters]
	words, err := random.Words(5, "--")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(words) // [three random words separated by --]
}
