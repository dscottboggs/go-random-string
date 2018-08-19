package random

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"path"
	"strings"
)

const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const wordListURL = "http://svnweb.freebsd.org/csrg/share/dict/words" +
	"?view=co&content-type=text/plain"

// where the word list will be cached to disk
var wordListLocation string

// where the current executable is located
var ThisFile string

func init() {
	var err error
	ThisFile, err = os.Executable()
	if err != nil {
		log.Fatal("Couldn't get executable file of random.go.")
	}
	wordListLocation = path.Join(path.Dir(ThisFile), "wordlist.txt")
}

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
		password += wordlist[index.Int64()] + sep
	}
	return password, nil
}

func getWordList() ([]string, error) {
	var (
		out     []string
		str     string
		bytestr []byte
		f       *os.File
		err     error
	)
	f, err = os.Open(wordListLocation)
	defer f.Close()
	if os.IsNotExist(err) {
		str, err = downloadWordList()
	} else {
		if err != nil && !os.IsExist(err) {
			return out, err
		}
		bytestr, err = ioutil.ReadAll(f)
		str = string(bytestr)
	}
	if err != nil {
		return out, err
	}
	return strings.Split(string(str), "\n"), nil
}

func downloadWordList() (string, error) {
	var (
		wordList string
		response *http.Response
		body     []byte
		err      error
	)
	response, err = http.Get(wordListURL)
	if err != nil {
		return wordList, err
	}
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return wordList, err
	}
	wordList = string(body)
	if response.StatusCode != 200 {
		return wordList, fmt.Errorf(
			"Got response %s(%d) when trying to download word list. Body "+
				"was %s",
			response.Status,
			response.StatusCode,
			wordList)
	}
	f, err := os.Create(wordListLocation)
	defer f.Close()
	if err != nil {
		return wordList, err
	}
	_, err = f.WriteString(wordList)
	return wordList, err
}
