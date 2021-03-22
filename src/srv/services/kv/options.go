package kv

import (
	"context"
)

type Options struct {
	Address   string
	CheckUser bool
	UserName  string
	Password  string

	Context context.Context
}

type Option func(*Options)

func Address(Address string) Option {
	return func(o *Options) {
		o.Address = Address
	}
}

func Auth(UserName string, Passwd string) Option {
	return func(o *Options) {
		o.CheckUser = true
		o.UserName = UserName
		o.Password = Passwd
	}
}
func SetJsonContext(key, val interface{}) Option {
	return func(o *Options) {
		o.Context = context.WithValue(o.Context, key, val)
	}
}
