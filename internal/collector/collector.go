package collector

import (
	"errors"

	"github.com/gocolly/colly/v2"
)

type Collector struct {
	*colly.Collector
}

func NewCollector(opts ...colly.CollectorOption) (*Collector, error) {
	c := colly.NewCollector(opts...)
	if c == nil {
		return nil, errors.New("failed to create collector")
	}
	return &Collector{c}, nil
}
