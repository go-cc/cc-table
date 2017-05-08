////////////////////////////////////////////////////////////////////////////
// Purpose: Turn the duoyinzi-phrase.txt file to json as Go dictionary
// Authors: Tong Sun (c) 2017
// Sources: https://github.com/mozillazg/phrase-pinyin-data
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spakin/awk"
)

// json string to return
var bufRet = bytes.NewBufferString("")

func main() {
	// https://godoc.org/github.com/spakin/awk
	s := awk.NewScript()

	// == BEGIN
	sa := s.NewValueArray()
	// last word, its matching count and length
	sa.Set("wLast", "")
	sa.Set("cMatch", 0)
	sa.Set("cLength", 0)

	// == Match & Process
	s.AppendStmt(nil, func(s *awk.Script) {
		ww, py := s.F(1).String(), ""
		for ii := 2; ii <= s.NF; ii++ {
			py += s.F(ii).String() + " "
		}
		ww = ww[:len(ww)-1] // last char is ":"
		print(s.NR, ww, py)

		cLength := sa.Get("cLength").Int()
		cMatch := commPrefixLen(sa.Get("wLast").String(), ww)
		print(" ", cLength, " ", cMatch)
		sa.Set("wLast", ww)
		//print(" Saved:", sa.Get("wLast").String())

		if lDiff := cLength - cMatch; lDiff == 0 {
			// the new phrase is longer than last one, ignore it
			s.Next()
		} else {
			// Only partial match, close the json, with (lDiff-1) closes
			for ii := 1; ii < lDiff; ii++ {
				fmt.Fprintf(bufRet, "},")
			}
		}
		outputEntry(ww, py, sa.Get("cMatch").Int())
		sa.Set("cMatch", cMatch)
		sa.Set("cLength", len([]rune(ww)))
	})

	// == END
	s.End = func(s *awk.Script) {
		ret := bufRet.String()
		fmt.Printf("%s}\n", ret)
	}

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

func outputEntry(s, py string, start int) {
	print(" output:", s, " from ", start, "\n")
	rs := []rune(s)

	for ii := start; ii < len(rs)-1; ii++ {
		fmt.Fprintf(bufRet, `"%s":{`, string(rs[ii]))
	}
	fmt.Fprintf(bufRet, `"%s":"%s",`, string(rs[len(rs)-1]), py)
}

func commPrefixLen(s1, s2 string) int {
	print(" compare:", s1, "<=>", s2)
	rs1, rs2 := []rune(s1), []rune(s2)
	for ii := 0; ii < min(len(rs1), len(rs2)); ii++ {
		if rs1[ii] != rs2[ii] {
			return ii
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
 */
