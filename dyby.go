package main

import (
	"dtby/config"
	"dtby/models"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func testMW() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Set example variable
		c.Set("example", "12345")

		// before request
		status1 := c.Writer.Status()
		log.Println(status1)
		log.Print("bgg")
		c.Next()

		// after request
		log.Print("aff")

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	configFilePath := flag.String("config", "config/setting.yml", "config file path")
	err := config.LoadConfig(*configFilePath)
	if err!=nil{
		fmt.Printf("LoadError: config.LoadConfig--%v", err)
	}
	db, err := models.InitDB()
	if err!=nil{
		panic(err)
	}
	defer db.Close()
	r := gin.Default()
	r.Use(testMW())
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("cctest")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//r.GET("/test", func(c *gin.Context) {
	//	// 创建并获取
	//	testh:=models.H{}
	//	testh.Name="axe"
	//	testh.Armor=2
	//	testh.Damage=123
	//	testh.MagicResist=234.3
	//	testh.HID=998
	//})
	r.GET("/test", func(c *gin.Context) {
		var level models.HeroLevel
		req:=db.First(&level)
		if req.Error != nil{
			fmt.Println(req.Error)
		}
		fmt.Println(level)
		c.JSON(200, gin.H{
			"success": 1,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

