package proxy

import (
	//"github.com/micro/go-micro"
	"github.com/micro/go-micro/selector"
)

type NextHandler (selector.Next)

func SelectorProxy(r selector.Selector) {
	//r.Options()
}
