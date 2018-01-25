package main

type Counter interface {
	Add(v int) int
}

type counterService struct {
	v int
}

func (c *counterService) Add(v int) int {
	c.v += v
	return c.v
}
