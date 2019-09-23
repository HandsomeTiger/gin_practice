package main

import "github.com/gin-gonic/gin"

// exampleFastBuild 快速开始
func exampleFastBuild() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 在 0.0.0.0:8080 上监听并服务
}

func exampleRouter(){
	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件创建一个 gin 路由：
	// 日志与恢复中间件（无崩溃）。
	router := gin.Default()

	router.GET("/someGet", exampleApi)
	router.POST("/somePost", exampleApi)
	router.PUT("/somePut", exampleApi)
	router.DELETE("/someDelete", exampleApi)
	router.PATCH("/somePatch", exampleApi)
	router.HEAD("/someHead", exampleApi)
	router.OPTIONS("/someOptions", exampleApi)

	// 默认情况下，它使用：8080，除非定义了 PORT 环境变量。
	router.Run()
	// router.Run(":3000") 硬编码端口
}

// exampleApi 测试接口用例
func exampleApi(ctx *gin.Context){

}
