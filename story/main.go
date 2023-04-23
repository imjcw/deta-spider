package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化gin引擎
	router := gin.Default()

	// 定义API的路由和处理函数
	router.GET("/hello", func(c *gin.Context) {
		resp, err := http.Get("/fetcher/hello")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
			return
		}
		defer resp.Body.Close()
		b, _ := ioutil.ReadAll(resp.Body)
		c.JSON(http.StatusOK, gin.H{"message": string(b)})
	})

	// 启动HTTP服务器，监听8080端口
	router.Run(":8080")
}
