package lark_rate_limiter_test

import (
	"github.com/chyroc/lark"
	"github.com/chyroc/lark_rate_limiter"
)

// Example_Wait: limit lark api request 5 qps
func Example_Wait() {
	lark.New(
		lark.WithAppCredential("", ""),
		lark.WithApiMiddleware(lark_rate_limiter.Wait(5, 5)),
	)
}
