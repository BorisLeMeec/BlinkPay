package grifts

import (
	"github.com/BorisLeMeec/BlinkPay/api/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
