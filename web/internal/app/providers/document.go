package providers

import (
	"github.com/crcms/blog/web/internal/app/handlers"
	"github.com/firmeve/firmeve/kernel"
	"github.com/firmeve/firmeve/kernel/contract"
)

type DocumentProvider struct {
	kernel.BaseProvider
}

func (d DocumentProvider) Name() string {
	return `document`
}

func (d DocumentProvider) Register() {

}

func (d *DocumentProvider) Boot() {
	documentRoute(d.Firmeve.Get(`http.router`).(contract.HttpRouter))
}

func documentRoute(router contract.HttpRouter)  {
	router.GET("/list", handlers.List)
}