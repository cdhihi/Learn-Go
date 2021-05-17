package main

import (
	"errors"
)

type Pay interface {
	Pay(string) string
}

var pays = make(map[string]func() Pay)

func Register(flag string, factory func() Pay) {
	pays[flag] = factory
}

func Create(flag string) (Pay, error) {

	if f, ok := pays[flag]; ok {
		return f(), nil
	} else {
		return nil, errors.New("找不到方法")
	}
}
