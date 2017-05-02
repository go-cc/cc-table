package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "中国人民银行"
  
  	// 默认
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong] [guo] [ren] [min] [yin] [xing]]

	// 包含声调
	a.Style = pinyin.Tone
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhōng] [guó] [rén] [mín] [yín] [xíng]]

	// 声调用数字表示
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng] [guo2] [re2n] [mi2n] [yi2n] [xi2ng]]

	// 声母风格
	a.Style = pinyin.Initials
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zh] [g] [r] [m] [] [x]]

	// 韵母风格
	a.Style = pinyin.FinalsTone
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[ōng] [uó] [én] [ín] [iín] [íng]]
}
