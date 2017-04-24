////////////////////////////////////////////////////////////////////////////
// Purpose: Turn the OpenCC STCharacters.txt file to dictionary
// Authors: Tong Sun (c) 2017
// Sources: https://github.com/go-cc/opencc-dict/blob/master/data/dictionary/STCharacters.txt
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/spakin/awk"
)

func main() {
	// https://godoc.org/github.com/spakin/awk
	s := awk.NewScript()

	// == BEGIN
	sa := s.NewValueArray()
	sa.Set("cS", "")
	sa.Set("cT", "")

	// == Match & Process
	s.AppendStmt(nil, func(s *awk.Script) {
		// print(s.NR)
		c1 := s.F(1).String()
		for ii := 2; ii <= s.NF; ii++ {
			c2 := s.F(ii).String()
			if c1 == c2 {
				continue
			}
			c2 = s.F(ii).String()
			//fmt.Fprintf(os.Stderr, "%v:%v\n", c1, c2)
			//fmt.Fprintf(os.Stderr, "%v:%v (%+q:%+q)\n", c1, c2, c1, c2)
			sa.Set("cS", fmt.Sprintf("%s%+q", sa.Get("cS"), c1))
			sa.Set("cT", fmt.Sprintf("%s%+q", sa.Get("cT"), c2))
		}
		// print(" ")
	})

	// == END
	s.End = func(s *awk.Script) {
		dqRegex := regexp.MustCompile(`"`)
		// s.Println(sa.Get("cS"), "\n", sa.Get("cT"))
		cS := dqRegex.ReplaceAllString(sa.Get("cS").String(), "")
		cT := dqRegex.ReplaceAllString(sa.Get("cT").String(), "")
		fmt.Printf(`// Package ccst - Chinese-Characters Simplified <=> Traditional
package ccst

// Automatically generated by cc-dict-gen.go
// Do not edit. 

const chS = "%s"`+`

const chT = "%s"
`, cS, cT)
	}

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
	//fmt.Fprintf(os.Stderr, "Succeeded")
}

/*
 */
