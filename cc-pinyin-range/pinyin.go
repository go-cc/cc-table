// package pybr stands for pinyin by range
package pybr

import (
	"bytes"

	"github.com/axgle/mahonia"
)

var (
	pyTableLen = len(pyValue)
)

type Pinyin struct {
	//分隔符
	Split string
	//是否首字母大写
	Upper bool
}

func New(split string) Pinyin {
	return Pinyin{
		Split: split,
		Upper: true,
	}
}

//get chinese pinyin number code
//param s must be chinese character with utf8 encoding
func Code(s string) int {
	//print(s)
	gbkString := UTF8ToGBK(s)
	//print(len(gbkString))
	var i1, i2 int
	i1 = int(gbkString[0])
	i2 = int(gbkString[1])
	return i1*256 + i2 - 65536
}

// converts utf8 chinese to pinyin
func (c *Pinyin) Convert(s string) string {
	pyString := bytes.NewBufferString("")
	var str string
	var code int

	for _, rune := range s {
		str = string(rune)
		//println(str + ": " + pyString.String())

		if rune <= '~' {
			pyString.WriteString(str)
		} else {
			code = Code(str)
			if code > 0 && code < 160 {
				pyString.WriteString(str)
			} else {
				if v, ok := tableMap[code]; ok { //map by table
					pyString.WriteString(v + c.Split)
				} else {
					found := false
					for i := pyTableLen - 1; i >= 0; i-- {
						if pyValue[i] <= code {
							pyString.WriteString(pyName[i] + c.Split)
							found = true
							break
						}
					}
					if !found {
						pyString.WriteString(str)
					}
				}
			}
		}
	}
	return pyString.String()
}

//define func UTF8ToGBK by using mahonia package
var enc = mahonia.NewEncoder("gbk")

func UTF8ToGBK(s string) string {
	return enc.ConvertString(s)
}
