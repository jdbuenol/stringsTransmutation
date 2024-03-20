package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//initialises a router with the default functions.
	router := gin.Default()

	router.GET("/concat/:str1/:str2", func(context *gin.Context) {
		str1 := context.Param("str1")
		str2 := context.Param("str2")
		context.String(http.StatusOK, str1+str2)
	})

	router.GET("/revert/:str", func(context *gin.Context) {
		str := context.Param("str")
		runes := []rune(str)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		context.String(http.StatusOK, string(runes))
	})

	// starts the server at port 8080
	router.Run(":8080")
}
