package model

import (
	"github.com/silly-zero/uuSchool/utils"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID         string    `json:"id" gorm:"primary_key;size:36;not null;unique"`
	CreateTime time.Time `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"updateTime" gorm:"autoUpdateTime"`
	DeleteTime time.Time `json:"deleteTime" gorm:"default:null"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = util.GenID()
	}
	return
}
