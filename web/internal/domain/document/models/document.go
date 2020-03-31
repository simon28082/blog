package models

import (
	"crypto/sha1"
	"time"
)

type Document struct {
	Uuid string `gorm:"char(32);primary_key;"`
	Title string `gorm:"type:varchar(60)"`
	Content string `grom:"type:text;default null;"`
	Description string `grom:"type:varchar(150);default null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d Document) TableName() string  {
	return `documents`
}

func (d *Document) BeforeCreate()  {
	d.Uuid = string(sha1.New().Sum([]byte(d.Title)))
}