package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Paper struct {
	gorm.Model
	Name    string `json:name`    // Proper name; may be empty.
	Date    string `json:date`    // Publish Date
	Journal string `json:journal` // Journal
	Doi     string `json:doi`     // Doi
}

// custom and exported Init function, this will not be called automatically
// by the go runtime like the special `init` function and therefore must be called
// manually by the package that imports this one.
func InitPaper(gormdb *gorm.DB, ginrouter *gin.Engine) {
	db = gormdb // set package global
	db.AutoMigrate(&Paper{})
	router = ginrouter
	api := router.Group("/api/v1/paper")
	{
		api.POST("/", CreatePaper)
		api.GET("/", GetPapers)
		api.GET("/:id", GetPaper)
		api.PUT("/:id", UpdatePaper)
		api.DELETE("/:id", DeletePaper)
	}
}

// CREATE
func CreatePaper(c *gin.Context) {
	var paper Paper
	c.BindJSON(&paper)
	db.Create(&paper)
	c.JSON(200, paper)
}

// READ
func GetPapers(c *gin.Context) {
	var paper []Paper
	if err := db.Find(&paper).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, paper)
	}
}

func GetPaper(c *gin.Context) {
	var paper Paper
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&paper).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, paper)
	}
}

// UPDATE
func UpdatePaper(c *gin.Context) {
	var paper Paper
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&paper).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&paper)
	db.Save(&paper)
	c.JSON(200, paper)
}

// DELETE
func DeletePaper(c *gin.Context) {
	var paper Paper
	id := c.Params.ByName("id")
	d := db.Where("id = ?", id).Delete(&paper)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
