package gojieba

import (
	"github.com/dearchun/gojieba/deps/cppjieba"
	"github.com/dearchun/gojieba/deps/limonp"
	"github.com/dearchun/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
