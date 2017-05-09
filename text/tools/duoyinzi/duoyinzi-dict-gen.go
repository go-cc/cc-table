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
	"strings"

	"github.com/spakin/awk"
)

// json string to return
var bufRet = bytes.NewBufferString("")

// var p = print X: use of builtin print not in function call
// Remove "//D>" to debug

func main() {
	// https://godoc.org/github.com/spakin/awk
	s := awk.NewScript()

	// == BEGIN
	sa := s.NewValueArray()
	// last word, and its length
	sa.Set("wLast", "")
	sa.Set("cLength", 0)
	bufRet.WriteByte('{')

	// == Match & Process
	s.AppendStmt(nil, func(s *awk.Script) {
		ww, py := s.F(1).String(), ""
		for ii := 2; ii <= s.NF; ii++ {
			py += s.F(ii).String() + " "
		}
		ww = ww[:len(ww)-1] // last char is ":"
		//D>print(s.NR, ww, py)

		// count of current word length and match length with last word
		cLength := sa.Get("cLength").Int()
		cMatch := commPrefixLen(sa.Get("wLast").String(), ww)
		//D>print(" ", cLength, " ", cMatch)

		lDiff := cLength - cMatch
		if lDiff == 0 && sa.Get("wLast").String() != "" {
			// the new phrase is longer than last one, ignore it
			//D>println()
			s.Next()
		}

		sa.Set("wLast", ww)
		//print(" Saved:", sa.Get("wLast").String())
		// Only partial match, close the json, with (lDiff-1) closes
		for ii := 1; ii < lDiff; ii++ {
			fmt.Fprintf(bufRet, "},")
		}
		outputEntry(ww, py, cMatch)
		sa.Set("cLength", len([]rune(ww)))
	})

	// == END
	s.End = func(s *awk.Script) {
		for ii := 0; ii < sa.Get("cLength").Int(); ii++ {
			bufRet.WriteByte('}')
		}
		//ret := bufRet.String()
		ret := strings.Replace(bufRet.String(), ",}", "}", -1)
		fmt.Printf("%s\n", ret)
	}

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}

func outputEntry(s, py string, start int) {
	//D>println(" output ", s, " from ", start)
	rs := []rune(s)

	for ii := start; ii < len(rs)-1; ii++ {
		fmt.Fprintf(bufRet, `"%s":{`, string(rs[ii]))
	}
	fmt.Fprintf(bufRet, `"%s":"%s",`, string(rs[len(rs)-1]), py)
}

func commPrefixLen(s1, s2 string) int {
	//D>print(" compare ", s1, ":", s2)
	rs1, rs2 := []rune(s1), []rune(s2)
	ii := 0
	for ; ii < min(len(rs1), len(rs2)); ii++ {
		if rs1[ii] != rs2[ii] {
			return ii
		}
	}
	return ii
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func debug(args ...interface{}) {
	// print(args...) X: invalid use of ... with builtin print
}

/*
 */
