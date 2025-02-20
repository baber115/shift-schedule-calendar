package strategy

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Strategy interface {
	Cal() error
}

type Common struct {
	date  string
	shift int
}

func NewCommon(tmp []string) Common {
	return Common{
		shift: getShift(tmp),
		date:  getDate(),
	}
}

func getShift(tmp []string) int {
	var shift int

	for i, v := range tmp {
		fmt.Printf("班次%d：%s\n", i+1, v)
	}
	fmt.Print("输入今天的班次编号（默认1）：")
	_, err := fmt.Scanln(&shift)
	if err != nil || shift < 0 || shift > len(tmp)-1 {
		return 0
	}
	return shift - 1
}

func getDate() string {
	var date string

	fmt.Print("输入今天的日期（默认今天）：")
	_, err := fmt.Scanln(&date)
	if err != nil || date == "" {
		date = time.Now().Format("20060102")
	}
	reg := regexp.MustCompile(`\d+`)
	ia := reg.FindAllString(date, -1)
	date = strings.Join(ia, "")

	return date
}

func (c *Common) getStartedAt() string {
	return time.Now().AddDate(0, 0, -c.shift).Format("20060102")
}
