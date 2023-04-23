package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    // 初始化gin引擎
    router := gin.Default()

    // 定义API的路由和处理函数
    router.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Hello, world111!"})
    })

    // 启动HTTP服务器，监听8080端口
    router.Run(":8080")
}