package models

import (
	"book/config"
	"gorm.io/gorm"
)

type Category struct {
	config.BaseModel
	Name   string  `gorm:"type:varchar(32);comment:分类名" json:"name" form:"name"`
	Remark string  `gorm:"type:text;comment:备注" json:"remark" form:"remark"`
	Books  []*Book `gorm:"foreignKey:CategoryID;references:ID"`
}

func GetCategoryList(name string) *gorm.DB {
	tx := DB.Model(new(Category))
	if name != "" {
		tx.Where("name like ?", "%"+name+"%")
	}
	return tx
}
