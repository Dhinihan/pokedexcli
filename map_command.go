package main

import (
	"fmt"
)

type state struct {
	started bool
	page    int
	limit   int
}

var s state

func (s *state) start() {
	if s.started {
		return
	}
	s.started = true
	s.page = 0
	s.limit = 20
}

func (s *state) next() {
	s.page += 1
}

func (s *state) previous() {
	s.page -= 1
	if s.page < 1 {
		s.page = 1
	}
}

func (s *state) offset() int {
	return (s.page - 1) * s.limit
}

func commandMap(c config) error {
	s.start()
	s.next()
	locations, err := c.client.GetLocation(s.limit, s.offset())
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, l := range locations {
		fmt.Println(l.Name)
	}
	return nil
}

func commandMapBack(c config) error {
	s.start()
	s.previous()
	locations, err := c.client.GetLocation(s.limit, s.offset())
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, l := range locations {
		fmt.Println(l.Name)
	}
	return nil
}
