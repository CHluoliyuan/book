package models

import (
	"book/config"
	"gorm.io/gorm"
)

type User struct {
	config.BaseModel
	Name    string  `gorm:"type:varchar(32);comment:姓名;" form:"name" json:"name"`
	Age     int     `gorm:"type:tinyint;comment:年龄;" json:"age" form:"age"`
	Address string  `gorm:"type:varchar(128);comment:地址;" json:"address" form:"address"`
	Phone   string  `gorm:"type:varchar(16);comment: 手机号" json:"phone"  form:"phone"`
	Sex     string  `gorm:"type:varchar(8);comment:性别" json:"sex" form:"sex"`
	Account int     `gorm:"type:int;comment:账户积分;default:0" json:"account" form:"account"`
	Books   []*Book `gorm:"many2many:Borrow;foreignKey:ID;joinForeignKey:UserID;References:ID;joinReferences:BookID"`
}

func GetUserList(phone, name string) *gorm.DB {
	tx := DB.Model(new(User))
	if phone != "" {
		tx.Where("phone like ?", "%"+phone+"%")
	}
	if name != "" {
		tx.Where("name like ?", "%"+name+"%")
	}
	return tx
}
