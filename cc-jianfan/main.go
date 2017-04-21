// Package ccst provides Conversion between Simplified and Traditional Chinese-Characters.
// ported from the lastest version from gojianfan 
// https://raw.githubusercontent.com/siongui/gojianfan/32795bb4bbc71a0a3a6c62907b1322355c2d49ae/jianfan.go
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
	if len(ChT) != len(ChS) {
		panic("cht and chs data length not equal")
	}

	for index, runeValueT := range ChT {
		runeValueS, _ := utf8.DecodeRuneInString(ChS[index:])
		t2sMapping[runeValueT] = runeValueS
		s2tMapping[runeValueS] = runeValueT
	}
}

// Convert Traditional Chinese to Simplified Chinese
func T2S(s string) string {
	var chs []rune
	for _, runeValue := range s {
		v, ok := t2sMapping[runeValue]
		if ok {
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
		v, ok := s2tMapping[runeValue]
		if ok {
			cht = append(cht, v)
		} else {
			cht = append(cht, runeValue)
		}
	}
	return string(cht)
}
