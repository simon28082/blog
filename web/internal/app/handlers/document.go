package handlers

import (
	"github.com/crcms/blog/web/internal/domain/document/services"
	"github.com/firmeve/firmeve/kernel/contract"
	"github.com/firmeve/firmeve/render"
	"github.com/jinzhu/gorm"
)

func List (ctx contract.Context)  {

	service := ctx.Resolve(new(services.Document)).(*services.Document)

	result := service.List(
		ctx.Resolve(`db.connection`).(*gorm.DB),
		ctx.Protocol().(contract.HttpProtocol).Request(),
	)

	ctx.RenderWith(200,render.JSON,result)

	ctx.Next()
}