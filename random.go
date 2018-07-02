package random

import (
	"math/rand"
	"strings"
	"time"
)

const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func AsciiPrintable() rune {
	rand.Seed(time.Now().UnixNano())
	return rune(rand.Intn(126-33) + 33)
}
func Alphanumeric() rune {
	ch := AsciiPrintable()
	if strings.Index(alphanumeric, string(ch)) == -1 {
		return Alphanumeric()
	}
	return ch
}

func String(length int) string {
	// this actually happens more slowly concurrently.
	out := ""
	for i := 0; i < length; i++ {
		out += string(AsciiPrintable())
	}
	return out
}

func AlphanumericString(length int) string {
	var out string
	for i := 0; i < length; i++ {
		out += string(Alphanumeric())
	}
	return out
}
