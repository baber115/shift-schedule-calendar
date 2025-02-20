package strategy

import (
	"fmt"
	strings2 "strings"
)

var Tmp = [][]string{
	{"早班", "晚班", "晚班休息", "休息"},
	{"早班", "晚班", "休息"},
}

func ChooseTmp() Strategy {
	for i, strings := range Tmp {
		fmt.Printf("模版%d：%s\n", i+1, strings2.Join(strings, " "))
	}

	var tmp int
	fmt.Print("请选择模板编号（默认1）：")
	_, _ = fmt.Scanln(&tmp)
	switch tmp {
	case 1:
		return NewStrategy1()
	case 2:
		return NewStrategy2()
	default:
		return NewStrategy1()
	}
}
