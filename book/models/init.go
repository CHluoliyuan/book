package models

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
)

func migrage(db *gorm.DB) {
	err := db.SetupJoinTable(&User{}, "Books", &Borrow{})
	err = db.AutoMigrate(&Admin{})
	err = db.AutoMigrate(&Category{}, &Book{})
	err = db.AutoMigrate(&User{})

	err = db.AutoMigrate(&Retur{})
	if err != nil {
		return
	}
}

func InitMysqlDB() {
	username := "123"   //账号
	password := "123"   //密码
	host := "127.0.0.1" //数据库地址，可以是Ip或者域名
	port := 3306        //数据库端口
	Dbname := "book"    //数据库名
	timeout := "10s"    //连接超时，10秒

	// username:password@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s",
		username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 迁移表
	migrage(db)
	// 连接成功
	DB = db
}

func InitRedisDB() {
	db := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	RDB = db
}
