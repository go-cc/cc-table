////////////////////////////////////////////////////////////////////////////
// Purpose: Turn the duoyinzi-phrase.txt file to json as Go dictionary
// Authors: Tong Sun (c) 2017
// Sources: https://github.com/mozillazg/phrase-pinyin-data
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"fmt"
)

// duoyinzi phrase dictionary from json
var dypDict map[string]interface{}

func main() {
	err := json.Unmarshal(dyPhrase, &dypDict)
	check(err)
	//fmt.Printf("%+v\n", dypDict)

	rs := []rune(strIn)
	// loop through the input string
	for ii := 0; ii < len(rs); ii++ {
		// fmt.Printf("%c", rs[ii])
		jj := ii
		dict, ok := dypDict[string(rs[ii])]
		if !ok {
			// not in duoyinzi phrase dictionary, output as-is
			fmt.Printf("%c", rs[ii])
			continue
		}
		// decent into the dictionary until found or mismatch
		for ok {
			//fmt.Printf(">%c", rs[ii])
			// is it the last node or still in the middle?
			if py, isStr := dict.(string); isStr {
				dict, ok = py, false
			} else {
				ii++
				dict, ok = dict.(map[string]interface{})[string(rs[ii])]
			}
		}
		switch dict.(type) {
		case string:
			// found it in the dictionary
			fmt.Printf("%s", dict.(string))
		// case map[string]interface{}:
		// 	// matching breaks at ii, output from jj to ii
		// 	fmt.Printf("=%s", rs[jj:ii])
		default:
			// matching breaks at ii, output from jj to ii
			fmt.Printf("%s", string(rs[jj:ii+1]))
		}
	}
	fmt.Println()
}

// Output:
// XYZEFGKE F G H I J XA L M N O P Q XA B C G H I J K XA B C D K UVXA B C D G X

//==========================================================================
// support functions

// check will quit on unanticipated errors with stack trace info for debug
func check(e error) {
	if e != nil {
		panic(e)
	}
}
