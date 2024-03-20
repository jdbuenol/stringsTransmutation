package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Concat(s1 string, s2 string) string {
	return s1 + " " + s2
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	//initialises a router with the default functions.
	router := gin.Default()

	router.GET("/Concat/:str1/:str2", func(context *gin.Context) {
		str1 := context.Param("str1")
		str2 := context.Param("str2")
		context.String(http.StatusOK, Concat(str1, str2))
	})

	router.GET("/Reverse/:str", func(context *gin.Context) {
		str := context.Param("str")
		context.String(http.StatusOK, Reverse(str))
	})

	// starts the server at port 8080
	router.Run(":8080")
}
