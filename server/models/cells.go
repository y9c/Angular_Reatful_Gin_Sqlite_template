package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// data in sqlite table
// index      tsne1      tsne2  label                 anno
// test Brain cell only

type Cell struct {
	ID    int32   `json:"id" gorm:"column:id;not null"`
	Tsne1 float32 `json:"tsne1" gorm:"column:tsne1;not null"`
	Tsne2 float32 `json:"tsne2" gorm:"column:tsne2;not null"`
	Label int32   `json:"label" gorm:"column:label;not null"`
	Anno  string  `json:"anno" gorm:"column:anno"`
}

func (Cell) TableName() string {
	return "tsne_Brain"
}

// custom and exported Init function, this will not be called automatically
// by the go runtime like the special `init` function and therefore must be called
// manually by the package that imports this one.
func InitCell(gormdb *gorm.DB, ginrouter *gin.Engine) {
	db = gormdb // set package global
	db.AutoMigrate(&Cell{})
	router = ginrouter
	api := router.Group("/api/v1/cell")
	{
		api.GET("/", GetCells)
		api.GET("/:id", GetCell)
	}
}

// READ
func GetCells(c *gin.Context) {
	var cell []Cell
	if err := db.Find(&cell).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cell)
	}
}

func GetCell(c *gin.Context) {
	var cell Cell
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&cell).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cell)
	}
}

// return a series for echarts
func GetCellSeries(c *gin.Context) {
}
