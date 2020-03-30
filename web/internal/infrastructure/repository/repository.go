package repository

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Repository struct {
}

func (u *Repository) Transaction(db *gorm.DB, fn func(db *gorm.DB) interface{}) (result interface{}, tError error) {
    var (
        tx     *gorm.DB
    )

    tx = db.Begin()

    defer func() {
        if err := recover(); err != nil {
            if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
                tError = fmt.Errorf("transaction rollabck execute error %w", rollbackErr)
            } else {
                tError = fmt.Errorf("%w", err)
            }
            return
        }

        if commitErr := tx.Commit().Error; commitErr != nil {
           tError = fmt.Errorf("transaction commit execute error %w", commitErr)
        }
    }()

    result = fn(tx)

    // 这里不能直接返回 defer的特性
    // 函数的返回过程是这样  先给返回赋值->调用defer->返回到调用函数中
    // 如果 return tError 则会先赋值 tError 此时 defer 还没有执行 也就是 nil
    // 参见 https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.4.html
    return
}
