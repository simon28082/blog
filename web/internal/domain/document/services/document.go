package services

import (
	"github.com/crcms/blog/web/internal/domain/document/models"
	"github.com/jinzhu/gorm"
	"github.com/firmeve/firmeve/converter/resource"
	"github.com/ulule/paging"
	"net/http"
)
var (
	pageOption = &paging.Options{
		DefaultLimit:  15,
		MaxLimit:      15 + 10,
		LimitKeyName:  "limit",
		OffsetKeyName: "offset",
	}
)

type Document struct {
}



func ( d *Document) List(db *gorm.DB,req *http.Request) *resource.Paginator {
	option := &resource.Option{
		Transformer: nil,
		Fields:      []string{"id", "content", "uuid",`title`},
	}
	var documents []models.Document
	store,_ := paging.NewGORMStore(db, &documents)
	return resource.NewPaginator(store,option,req,pageOption)
}
