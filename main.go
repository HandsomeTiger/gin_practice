package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	s:=gin.Default()
	s.GET("/")
}



func example(){
	serve:=gin.Default()
	serve.Use(func(ctx *gin.Context) {
		fmt.Println("middleer")})

	serve.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"l":"lksadjf",
		})},)

	gin.Logger()
	serve.Run()
}
