package models

import (
	"book/config"
	"gorm.io/gorm"
)

type Borrow struct {
	config.BaseModel
	BookID     uint   ` form:"book_id" json:"book_id"`
	Books      *Book  `gorm:"foreignKey:BookID"`
	UserID     uint   `form:"user_id" json:"user_id"`
	Users      *User  `gorm:"foreignKey:UserID"`
	Days       int    `form:"days" json:"days"`
	ReturnDate string `json:"return_date"  gorm:"type:date;comment:归还时间"`
	Note       string ` json:"note"` //正常 即将到期 已过期 已还
}

func GetBorrowList(user_name, book_name, book_id string) *gorm.DB {
	tx := DB.Model(new(Borrow)).Preload("Books").Preload("Users").
		Joins(" JOIN books ON books.id = borrows.book_id").
		Joins(" JOIN users ON users.id = borrows.user_id")
	if user_name != "" {
		tx.Where("users.name like ?", "%"+user_name+"%")
	}
	if book_name != "" {
		tx.Where("books.name like ?", "%"+book_name+"%")
	}
	if book_id != "" {
		tx.Where("books.id like ?", "%"+book_id+"%")
	}
	return tx
}
