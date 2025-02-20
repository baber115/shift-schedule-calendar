package main

import (
	"github.com/baber115/shift-schedule-calendar/strategy"
)

var s strategy.Strategy

func init() {
	s = strategy.ChooseTmp()
}

type Client struct {
	s strategy.Strategy
}

func NewClient(s strategy.Strategy) *Client {
	return &Client{
		s: s,
	}
}

func (c Client) do() {
	err := c.s.Cal()
	if err != nil {
		panic(err)
	}
}

func main() {
	client := NewClient(s)
	client.do()
}
