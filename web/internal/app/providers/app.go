package providers

import (
	"github.com/crcms/blog/web/internal/interfaces/web/handlers"
	"github.com/firmeve/firmeve/kernel"
	"github.com/firmeve/firmeve/kernel/contract"
	"github.com/firmeve/firmeve/support/path"
)

type AppProvider struct {
	kernel.BaseProvider
}

func (a AppProvider) Name() string {
	return `app`
}

func (a AppProvider) Register() {

}

func (a *AppProvider) Boot() {
	appRoute(a.Firmeve.Get(`http.router`).(contract.HttpRouter))
}

func appRoute(router contract.HttpRouter) {
	router.Static("/static", path.RunRelative("../../../static/public"))
	//router.Static("/static", path.RunRelative("../../../public"))
	r := router.Group("")
	{
		r.GET("/", handlers.Index)
	}
}
