package nlp

import (
	//"reflect"
	"testing"
	"github.com/stretchr/testify/require"

)

func TestTokemize(t *testing.T) {
	
	testCases := []struct {
		text     string
		expected []string
	}{
		{"who's on first", []string{"who", "s", "on", "first"}},
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
	//
	//	text := "Who's on first"
	//	expected := []string{"who", "s", "on", "first"}
	//
	//	out := Tokenize(text)
	//	//if expected != out {
	//	if !reflect.DeepEqual(expected, out) {
	//		t.Fatalf("%#v != %#v ", expected, out)
	//	}
	//
	//
	//	text = ""
	//	expected = []string(nil) //if this is not nil the test will fail
	//	out = Tokenize(text)
	//	if !reflect.DeepEqual(expected, out) {
	//		t.Fatalf("%#v != %#v ", expected, out)
	//	}

}
