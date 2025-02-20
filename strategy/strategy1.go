package strategy

import (
	"fmt"
	"time"
)

var _ Strategy = (*Strategy1)(nil)

type Strategy1 struct {
	c     Common
	tmp   []string
	shift int
}

func NewStrategy1() Strategy1 {
	tmp := Tmp[0]
	return Strategy1{
		c:   NewCommon(tmp),
		tmp: tmp,
	}
}

func (s Strategy1) Cal() error {
	now, _ := time.Parse("20060102", s.c.getStartedAt())
	otherTime, _ := time.Parse("20060102", s.c.date)

	duration := now.Sub(otherTime).Abs()
	// 计算天数差（考虑夏令时和闰秒的影响，使用Truncate方法）
	days := int(duration.Hours() / 24)
	for i := 0; i <= days; i++ {
		j := i % len(s.tmp)
		str := now.AddDate(0, 0, i).Format(time.DateOnly)
		fmt.Println(j, str, s.tmp[j])
	}

	return nil
}
