package models

import (
	"book/config"
	"gorm.io/gorm"
)

type Book struct {
	config.BaseModel
	Name        string ` form:"name" json:"name" gorm:"type:varchar(128);comment:书名"`
	Description string `form:"description" json:"description" gorm:"type:text;comment:描述"`
	PublishDate string ` form:"publish_date" json:"publish_date" gorm:"type:date;comment:出版时间"`
	Author      string ` form:"author" json:"author" gorm:"comment:出版时间"`
	Publisher   string `form:"publisher" json:"publisher" gorm:"type:varchar(64);comment:出版社"`
	Score       int    `form:"score" json:"score" gorm:"type:int(4);comment:积分"` //积分
	Nums        int    `form:"nums" json:"nums" gorm:"type:int(4);comment:数量" `  //数量
	CategoryID  uint   `form:"category_id" json:"category_id" gorm:"category_id"`
	//Categorys      *Category  `gorm:"foreignKey:CategoryID"`
	Cover string  `form:"cover"  json:"cover" gorm:"type:VARCHAR(255);comment:文件路径"`
	Users []*User `gorm:"many2many:Borrow;foreignKey:ID;joinForeignKey:BookID;References:ID;joinReferences:UserID"`
}

func GetBookList(name, bookNo string) *gorm.DB {
	tx := DB.Model(new(Book))
	if name != "" {
		tx.Where("name like ?", "%"+name+"%")
	}
	if bookNo != "" {
		tx.Where("id like ?", "%"+bookNo+"%")
	}
	return tx
}
