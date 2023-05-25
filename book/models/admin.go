package models

import (
	"book/config"
	"gorm.io/gorm"
)

type Admin struct {
	config.BaseModel
	Uuid     string `gorm:"type:varchar(36);comment:uuid唯一标识;index;unique" json:"uuid"`
	Name     string `gorm:"type: varchar(64);comment: 用户名" json:"name" form:"name"`
	Password string `gorm:"type: varchar(32);comment: 密码" json:"password" form:"password"`
	Phone    string `gorm:"type: varchar(16);comment: 手机号" json:"phone"  form:"phone"`
	Email    string `gorm:"type: varchar(64);comment: 邮箱" json:"email" form:"email"`
}

func GetAdminList(name, phone, email string) *gorm.DB {
	tx := DB.Model(new(Admin))
	if name != "" {
		tx.Where("name LIKE ?", "%"+name+"%")
	}
	if phone != "" {
		tx.Where("phone LIKE ?", "%"+phone+"%")
	}
	if email != "" {
		tx.Where("email LIKE ?", "%"+email+"%")
	}
	return tx
}
