package nlp

import (
	//"reflect"
	"testing"
//	"testing/quick" 
	
	"github.com/stretchr/testify/require"

)
//fuzzy testing?, hypothesis for python!

func TestTokemize(t *testing.T) {
	
	testCases := []struct {
		text     string
		expected []string
	}{
		{"who's on first", []string{"who", "on", "first"}},
		{"", []string(nil)},
	}


	for _, tc := range testCases {
		
		name := tc.text
		if name == ""{
			name = "<empty>"
		}
		
		t.Run(tc.text, func(t *testing.T) {
				require := require.New(t)
			out := Tokenize(tc.text)
			//if expected != out {
			
			require.Equal(tc.expected, out)
//			if !reflect.DeepEqual(tc.expected, out) {
//				t.Fatalf("%#v != %#v ", tc.expected, out)
//			}
		})
	}
}



//
//func TestQuick(t *testing.T){
//	require := require.New(t)
//	fn := func(text string) bool{
//		tokens := Tokenize(text)
//		if len(wordRe.FindAllString(text, -1) != len(tokens){
//				t.Log(text)
//				return false
//			}
//			return true
//		}
//	require.NoError(quick.Check(fn, nil))
//}
