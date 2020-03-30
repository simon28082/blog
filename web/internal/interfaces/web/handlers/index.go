package handlers

import (
	"github.com/firmeve/firmeve/kernel/contract"
	//"github.com/firmeve/firmeve/render"
 	"html/template"
)

func Index(ctx contract.Context)  {
	//ctx.RenderWith(200,render.Plain,"Index")


	//@todo 需要一个setHeader方法 仅作用于Response
	ctx.Protocol().(contract.HttpProtocol).ResponseWriter().Header().Set(`Content-Type`,`text/html`)
	t, _ := template.ParseFiles("views/index.html")

	t.Execute(ctx.Protocol().(contract.HttpProtocol).ResponseWriter(),nil)
}
