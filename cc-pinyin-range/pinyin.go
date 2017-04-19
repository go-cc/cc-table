// package pybr stands for pinyin by range
package pybr

import (
	"strings"
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

func New() Pinyin {
	return Pinyin{
		Split: " ",
		Upper: true,
	}
}

func (c *Pinyin) Convert(s string) (string, error) {
	var pyStr string

	pyArr := c.gbkToPinyin(utf8ToGbk(s))
	pyStr = strings.Join(pyArr, c.Split)

	return pyStr, nil
}

func (c *Pinyin) gbkToPinyin(gbk string) []string {
	var pyStr []string

	for i := 0; i < len(gbk); i++ {
		p := int(gbk[i])
		if p > 0 && p < 160 {
			pyStr = append(pyStr, string(gbk[i]))
		} else {
			i++
			q := int(gbk[i])
			p = p*256 + q - 65536

			py := pinyinSearch(p)
			if py != "" {
				if c.Upper == false {
					py = strings.ToLower(py)
				}
				pyStr = append(pyStr, py)
			}
		}
	}

	return pyStr
}

func utf8ToGbk(s string) string {
	enc := mahonia.NewEncoder("gbk")
	return enc.ConvertString(s)
}

func pinyinSearch(p int) string {
	if v, ok := tableMap[p]; ok {
		return v
	} else {
		for i := pyTableLen - 1; i >= 0; i-- {
			if pyValue[i] <= p {
				return pyName[i]
			}
		}
	}
	return ""
}
