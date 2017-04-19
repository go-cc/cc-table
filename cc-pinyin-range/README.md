# cc-pinyin-range
Chinese-Character to Pinyin lookup table by range. 
中文转拼音golang实现

## example

```go
package main

import (
	"fmt"
	pybr "github.com/go-cc/cc-table/cc-pinyin-range"
)

func main() {
	s := `名著：《红楼梦》〖清〗曹雪芹 著、高鹗 续／『人民文学』出版社／1996—9月30日／59.70【元】，《三国演义》〖明〗罗贯中。`

	py := pybr.New("")
	p := py.Convert(s)
	fmt.Println(p)
    // MingZhu：《HongLouMeng》〖Qing〗CaoXueQin Zhu、GaoZuo Xu／『RenMinWenXue』ChuBanShe／1996—9Yue30Ri／59.70【Yuan】，《SanGuoYanYi》〖Ming〗LuoGuanZhong。

	//设置分隔符
	//首字母是否大写
	py.Split = "-"
	py.Upper = false
	p = py.Convert(s)
	fmt.Println(p)
    // Ming-Zhu-：《Hong-Lou-Meng-》〖Qing-〗Cao-Xue-Qin- Zhu-、Gao-Zuo- Xu-／『Ren-Min-Wen-Xue-』Chu-Ban-She-／1996—9Yue-30Ri-／59.70【Yuan-】，《San-Guo-Yan-Yi-》〖Ming-〗Luo-Guan-Zhong-。
}
```

# Credits

Ported from [scofieldpeng/golang-chinese-to-pinyin](https://github.com/scofieldpeng/golang-chinese-to-pinyin/), which is from [l2x/golang-chinese-to-pinyin](https://github.com/l2x/golang-chinese-to-pinyin).

The code->pinyin mapping table is originally coded By [github.com/axgle](https://github.com/axgle) and [Leon @ github.com/leonzhu1981](https://github.com/leonzhu1981).
