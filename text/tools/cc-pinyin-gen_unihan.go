////////////////////////////////////////////////////////////////////////////
// Purpose: Create pinyin dictionary from Unihan/Unihan_Readings.txt
// Authors: Tong Sun (c) 2017
// Sources: http://www.unicode.org/Public/UNIDATA/Unihan.zip
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/spakin/awk"
)

/*

Input format:

U+3401 kMandarin tiàn
U+3406 kCantonese zaan2
U+3406 kMandarin yǐn
U+340C kCantonese zyu4
U+340C kMandarin yí
U+3437 kMandarin mǎ
U+3437 kCantonese maa6
U+884C kMandarin xíng
U+884C kMandarin háng
U+884C kCantonese haang4
U+884C kCantonese hang4
U+884C kCantonese hang6
U+884C kCantonese hong4
U+3437 kMandarin mǎ
U+3437 kCantonese maa6

Test:

cat cc-pinyin-gen_unihan.txt | go run cc-pinyin-gen_unihan.go kMandarin
cat cc-pinyin-gen_unihan.txt | go run cc-pinyin-gen_unihan.go kCantonese

*/

var focus = ""

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:\n   cc-pinyin-gen <kMandarin|kCantonese>")
		os.Exit(0)
	}

	focus = os.Args[1]

	// https://godoc.org/github.com/spakin/awk
	s := awk.NewScript()

	// == BEGIN
	s.State = ""

	// == Match & Process
	// /^U\+(\w+)\s+(kCantonese|kMandarin)\s+(.+)$/
	s.AppendStmt(func(s *awk.Script) bool {
		return s.F(0).Match(`^U\+(\w+)\s+(kCantonese|kMandarin)\s+(.+)$`)
	}, func(s *awk.Script) {
		// print(s.NR)
		// if $2 eq 'kCantonese'|'kMandarin'
		if s.F(2).Match(focus) {
			// $code = $1; $pinyin = $3;
			code := s.F(1).String()
			pinyin := s.F(3).String()
			//print(code, pinyin)
			// if code != last
			if code != s.State {
				if s.State != "" {
					fmt.Println()
				}
				//code = s.State.(string)
				fmt.Printf("%s\t", code)
				s.State = code
			} else {
				fmt.Printf(",")
			}
			fmt.Printf("%s", pinyin)
		}
		// print(" ")
	})

	// == END
	s.End = func(s *awk.Script) {
		fmt.Println()
	}

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "Succeeded")
}

/*
 */
