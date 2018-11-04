package random

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func AsciiPrintable() rune {
	char, err := rand.Int(rand.Reader, big.NewInt(126-33))
	if err != nil {
		return rune(0)
	}
	return rune(char.Int64() + 33)
}
func Alphanumeric() rune {
	ch := AsciiPrintable()
	if strings.Index(alphanumeric, string(ch)) == -1 {
		return Alphanumeric()
	}
	return ch
}

func String(length uint16) string {
	// this actually happens more slowly concurrently.
	var i uint16 = 0
	out := ""
	for ; i < length; i++ {
		out += string(AsciiPrintable())
	}
	return out
}

func AlphanumericString(length uint16) string {
	var out string
	var i uint16 = 0
	for ; i < length; i++ {
		out += string(Alphanumeric())
	}
	return out
}

//  from getWordList() of pwlen words separated by pwsep.
func Words(words uint8, sep string) (string, error) {
	wordlist, err := getWordList()
	if err != nil {
		return "", err
	}
	numberOfWords := big.NewInt(int64(len(wordlist)))
	var password string
	var i uint8
	for ; i < words; i++ {
		index, _ := rand.Int(rand.Reader, numberOfWords)
		if i == words-1 {
			password += wordlist[index.Int64()]
		} else {
			password += wordlist[index.Int64()] + sep
		}
	}
	return password, nil
}
