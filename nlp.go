package nlp

import (
	"regexp"
	"strings"
	
	"github.com/olgashtivelman/nlp/stemmer"
)

var (
	wordRe = regexp.MustCompile("[[:alpha:]]+")
)

func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		// TODO: stem
		//token := strings.ToLower(w)
		token := stemmer.Stem(strings.ToLower(w))
		if token != ""{
			tokens = append(tokens, token)
		}
	}
	return tokens
}
