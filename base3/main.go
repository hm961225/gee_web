package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
	})
	r.GET("assets/*filePath", func(c *gee.Context){
		c.JSON(http.StatusOK, gee.H{"file_path": c.Param("filePath")})
	})
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	err := r.Run(":9999")
	if err != nil {
		errText := fmt.Sprintf("启动失败%s", err)
		fmt.Println(errText)
	}
}
