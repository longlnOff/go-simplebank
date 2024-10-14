package routers

import accounts "github.com/longln/go-simplebank/internal/routers/account"

type RouterGroup struct {
	Account accounts.AccountRouterGroup
}

var RouterGroupAccount = new(RouterGroup)
