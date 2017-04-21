package pybr_test

import (
	"fmt"

	pybr "github.com/go-cc/cc-table/cc-pinyin-range"
)

// _output
func ExamplePinyin() {
	s := `名著：《红楼梦》〖清〗曹雪芹 著、高鹗 续／『人民文学』出版社／1996—9月30日／59.70【元】，《三国演义》〖明〗罗贯中。`

	py := pybr.New("")
	p := py.Convert(s)

	fmt.Println(p)

	//设置分隔符
	//首字母是否大写
	py.Split = "-"
	py.Upper = false
	p = py.Convert(s)
	fmt.Println(p)
}
