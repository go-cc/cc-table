// Package ccst provides Conversion between Simplified and Traditional Chinese-Characters.
//
// This is a simple implementation of Chinese one-to-one conversion.
// If you need more advanced converter, please visit
// OpenCC https://github.com/BYVoid/OpenCC

package ccst

import (
	"unicode/utf8"
)

var t2sMapping = make(map[rune]rune)
var s2tMapping = make(map[rune]rune)

func init() {
	if len([]rune(ChT)) != len([]rune(ChS)) {
		panic("cht and chs data length not equal")
	}

	for index, runeValueT := range ChT {
		runeValueS, _ := utf8.DecodeRuneInString(ChS[index:])
		t2sMapping[runeValueT] = runeValueS
		// Simplified => Traditional is one to many, pick first
		if _, ok := s2tMapping[runeValueS]; !ok {
			s2tMapping[runeValueS] = runeValueT
		}
	}
}

// Convert Traditional Chinese to Simplified Chinese
func T2S(s string) string {
	var chs []rune
	for _, runeValue := range s {
		if v, ok := t2sMapping[runeValue]; ok {
			chs = append(chs, v)
		} else {
			chs = append(chs, runeValue)
		}
	}
	return string(chs)
}

// Convert Simplified Chinese to Traditional Chinese
func S2T(s string) string {
	var cht []rune
	for _, runeValue := range s {
		if v, ok := s2tMapping[runeValue]; ok {
			cht = append(cht, v)
		} else {
			cht = append(cht, runeValue)
		}
	}
	return string(cht)
}
