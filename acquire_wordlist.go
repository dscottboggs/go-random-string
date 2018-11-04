package random

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const (
	wordListURL = "http://app.aspell.net/create?max_size=95&spelling=US&spelling=GBs&spelling=GBz&spelling=CA&spelling=AU&max_variant=3&diacritic=strip&special=hacker&special=roman-numerals&download=wordlist&encoding=utf-8&format=inline"
)

func getExecutable() (exe string) {
	var err error
	exe, err = os.Executable()
	if err != nil {
		log.Fatal("Couldn't get executable file of random.go.")
	}
	return
}

var WordListLocation = path.Join(path.Dir(getExecutable()), "wordlist.txt")

func pull(callback func(string), filterOut func(string) bool) error {
	res, err := http.Get(wordListURL)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("got status %d trying to get wordlist", res.StatusCode)
	}
	defer res.Body.Close()
	return mapResults(callback, filterOut, res.Body)
}

func mapResults(callback func(string), filterOut func(string) bool, body io.Reader) error {
	var lastWordFromPreviousRound string
	buffer := make([]byte, 128) // definitly longer than any single word
	for {
		rcvd, err := body.Read(buffer)
		if err == io.EOF {
			if rcvd == 0 {
				break
			}
			words := strings.Split(
				lastWordFromPreviousRound+string(buffer),
				"\n",
			)
			for _, word := range words {
				if !filterOut(word) {
					callback(word)
				}
			}
			break
		} else if err != nil {
			return err
		}
		words := strings.Split(
			lastWordFromPreviousRound+string(buffer),
			"\n",
		)
		for i, word := range words {
			if i == len(words)-1 {
				lastWordFromPreviousRound = word
				break
			}
			if !filterOut(word) {
				callback(word)
			}
		}
	}
	return nil
}

func getWordList() ([]string, error) {
	var (
		out []string
		f   *os.File
		err error
	)
	f, err = os.Open(WordListLocation)
	defer f.Close()
	if os.IsNotExist(err) {
		downloadWordList()
		return getWordList()
	} else {
		if err != nil && !os.IsExist(err) {
			return out, err
		}
		return out, mapResults(
			func(word string) {
				out = append(out, word)
			},
			func(word string) bool {
				// filter out words with apostrophes
				return strings.Index(word, "'") != -1
			},
			f,
		)

	}
}

func downloadWordList() error {
	var (
		response *http.Response
		err      error
	)
	response, err = http.Get(wordListURL)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf(
			"Got response %s(%d) when trying to download word list. Body "+
				"was %s",
			response.Status,
			response.StatusCode,
			string(body))
	}
	f, err := os.Create(WordListLocation)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = io.Copy(f, response.Body)
	return err
}
