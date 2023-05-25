package main

import (
	"book/models"
	"book/router"
)

func main() {
	models.InitMysqlDB()
	models.InitRedisDB()
	r := router.Router()
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
