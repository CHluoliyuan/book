package services

import (
	"book/models"
	"book/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// EchartsData
// @Tags echarts数据
// @Summary echarts数据
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /echart_data [get]
func EchartsData(c *gin.Context) {
	utils.Log.Infoln("enter")
	//timeRange := c.Param("timeRange")
	type BorrowResult struct {
		CategoryID uint   `json:"category_id"`
		Count      string `json:"count"`
	}

	utils.Log.Infoln("getting data")
	tx := models.DB.Begin()
	var borrow_result []BorrowResult
	err := tx.Table("borrows").Select("books.category_id, count(*) as count").
		Joins("join books on books.id = borrows.book_id").Group("books.category_id").
		Find(&borrow_result).Error

	type ReturResult struct {
		CategoryID uint   `json:"category_id"`
		Count      string `json:"count"`
	}

	var retur_result []ReturResult
	err = tx.Table("returs").Select("books.category_id, count(*) as count").
		Joins("join books on books.id = returs.book_id").Group("books.category_id").
		Find(&retur_result).Error

	var categories []models.Category
	err = tx.Model(new(models.Category)).Select("id", "name").Find(&categories).Error

	if err != nil {
		utils.Log.Warning("get data error")
		tx.Rollback()
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "find category error:" + err.Error(),
		})
		return
	}

	data := make(map[uint][]string, len(categories))
	for _, v := range categories {
		data[v.ID] = append(data[v.ID], v.Name)
	}
	for _, v := range borrow_result {
		data[v.CategoryID] = append(data[v.CategoryID], v.Count)
	}
	for k, v := range data {
		if len(v) != 2 {
			data[k] = append(data[k], "0")
		}
	}

	for _, v := range retur_result {
		data[v.CategoryID] = append(data[v.CategoryID], v.Count)
	}
	for k, v := range data {
		if len(v) != 3 {
			data[k] = append(data[k], "0")
		}
	}

	jsonData := make(map[uint]interface{})
	for key, value := range data {
		jsonData[key] = value
	}

	utils.Log.Infoln("return")
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": jsonData,
	})
}
