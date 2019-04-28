package nlp_test

import (
	"fmt"
	"github.com/olgashtivelman/nlp"

)

func ExampleTokenize(){
	text := "what's up doc"
	tokens := nlp.Tokenize(test)
	fmt.Println(tokens)
	
	// Output:
	// what s up doc
}
