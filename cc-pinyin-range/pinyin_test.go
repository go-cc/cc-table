package pybr_test

import (
	"fmt"

	pybr "github.com/go-cc/cc-table/cc-pinyin-range"
)

// for standalone test, change package to main and the next func def to,
// func main() {
func ExamplePinyin_output() {
	s := `名著：《红楼梦》〖清〗曹雪芹 著、高鹗 续／『人民文学』出版社／1996—9月30日／59.70【元】，《三国演义》〖明〗罗贯中。`

	py := pybr.New("")
	p := py.Convert(s)

	fmt.Println(p)
	// Output:
	//   not here!

	//设置分隔符
	//首字母是否大写
	py.Split = "-"
	py.Upper = false
	p = py.Convert(s)
	fmt.Println(p)

	// Output:
	// MingZhu：《HongLouMeng》〖Qing〗CaoXueQin Zhu、GaoZuo Xu／『RenMinWenXue』ChuBanShe／1996—9Yue30Ri／59.70【Yuan】，《SanGuoYanYi》〖Ming〗LuoGuanZhong。
	// ming-zhu-：《hong-lou-meng-》〖qing-〗cao-xue-qin- zhu-、gao-zuo- xu-／『ren-min-wen-xue-』chu-ban-she-／1996—9yue-30ri-／59.70【yuan-】，《san-guo-yan-yi-》〖ming-〗luo-guan-zhong-。

}
