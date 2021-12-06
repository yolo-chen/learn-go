package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"time"
)

func tttt(ctx *gin.Context) {
	ctx.Writer.WriteString("hello")
}

func main() {
	g := gin.Default()
	g.GET("/hello", StatCost(), func(c *gin.Context) { // 使用中间件方式1
		get := c.MustGet("name")

		// 手动组装响应
		c.JSON(200, gin.H{
			"message": "hello",
			"get":     get,
		})
	})

	g.GET("/ttt", tttt) // 此处为回调函数  回调指的是将

	g.GET("/hei", func(context *gin.Context) { // 使用中间件方式 2
		start := time.Now()
		log.Println("start si ", start)
		context.Set("timeStart", start)
	}, func(c *gin.Context) {
		start := c.MustGet("timeStart")
		t := start.(time.Time)

		since := time.Since(t)
		log.Println("since is ", since)
		log.Println("end si ", time.Now())
		log.Printf("start type %v", reflect.TypeOf(start))

		//使用结构体响应  注意： 结构体内字段首字母必须大写，否则json格式化时会读取不到，导致响应为空
		var msg struct {
			Name string `json:"user"`
			Age  int
		}

		msg.Age = 13
		msg.Name = "lisi"

		c.JSON(http.StatusOK, msg)
	})

	g.Run(":8090")
}

func StatCost() gin.HandlerFunc { // 使用gin的中间件必须要实现 gin.HandlerFunc方法
	return func(context *gin.Context) {
		start := time.Now()
		context.Set("name", "adong")
		context.Next()

		cost := time.Since(start)
		log.Println("shijian =", cost)
	}
}
