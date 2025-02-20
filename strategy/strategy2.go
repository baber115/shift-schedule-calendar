package strategy

import (
	"fmt"
	"time"
)

var _ Strategy = (*Strategy2)(nil)

type Strategy2 struct {
	c     Common
	tmp   []string
	shift int
}

func NewStrategy2() Strategy1 {
	tmp := Tmp[1]
	return Strategy1{
		c:   NewCommon(tmp),
		tmp: tmp,
	}
}

func (s Strategy2) Cal() error {
	now, _ := time.Parse("20060102", s.c.getStartedAt())
	otherTime, _ := time.Parse("20060102", s.c.date)

	duration := now.Sub(otherTime).Abs()
	days := int(duration.Hours() / 24)
	for i := 0; i <= days; i++ {
		j := i % len(s.tmp)
		str := now.AddDate(0, 0, i).Format(time.DateOnly)
		fmt.Println(j, str, s.tmp[j])
	}

	return nil
}
