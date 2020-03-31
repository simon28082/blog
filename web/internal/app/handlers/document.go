package handlers

import (
	"fmt"
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

	fmt.Printf("%#v",result)

	ctx.RenderWith(200,render.JSON,result.CollectionData())

	ctx.Next()
}