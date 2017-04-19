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
    s := "天气不错"

    py := pybr.New()
    p, _ := py.Convert(s)

    fmt.Println(p)
    //Tian Qi Bu Cuo


    //设置分隔符
    //首字母是否大写
    py.Split = "-"
    py.Upper = false
    p, _ = py.Convert(s)
    fmt.Println(p)
    //tian-qi-bu-cuo
}
```

# Credits

Ported from [scofieldpeng/golang-chinese-to-pinyin](https://github.com/scofieldpeng/golang-chinese-to-pinyin/), which is from [l2x/golang-chinese-to-pinyin](https://github.com/l2x/golang-chinese-to-pinyin).

The code->pinyin mapping table is originally coded By [github.com/axgle](https://github.com/axgle) and [Leon @ github.com/leonzhu1981](https://github.com/leonzhu1981).
