package main

import (
	"fmt"

	"./models"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// read from sqlite database
	db, err := gorm.Open("sqlite3", "./server/data/db.sqlite3")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// build router
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./client/dist", true)))
	// manually initialize imported packages
	models.InitMember(db, router)
	models.InitPaper(db, router)
	// run router
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
