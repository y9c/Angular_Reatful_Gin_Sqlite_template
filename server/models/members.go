package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Name  string `json:name`  // Proper name; may be empty.
	Email string `json:email` // user@domain
}

// custom and exported Init function, this will not be called automatically
// by the go runtime like the special `init` function and therefore must be called
// manually by the package that imports this one.
func InitMember(gormdb *gorm.DB, ginrouter *gin.Engine) {
	db = gormdb // set package global
	db.AutoMigrate(&Member{})
	router = ginrouter
	api := router.Group("/api/v1/member")
	{
		api.POST("/", CreateMember)
		api.GET("/", GetMembers)
		api.GET("/:id", GetMember)
		api.PUT("/:id", UpdateMember)
		api.DELETE("/:id", DeleteMember)
	}
}

// CREATE
func CreateMember(c *gin.Context) {
	var member Member
	c.BindJSON(&member)
	db.Create(&member)
	c.JSON(200, member)
}

// READ
func GetMembers(c *gin.Context) {
	var member []Member
	if err := db.Find(&member).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, member)
	}
}

func GetMember(c *gin.Context) {
	var member Member
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&member).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, member)
	}
}

// UPDATE
func UpdateMember(c *gin.Context) {
	var member Member
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&member).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&member)
	db.Save(&member)
	c.JSON(200, member)
}

// DELETE
func DeleteMember(c *gin.Context) {
	var member Member
	id := c.Params.ByName("id")
	d := db.Where("id = ?", id).Delete(&member)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
