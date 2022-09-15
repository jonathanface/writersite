package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("ok!!!")
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.0.1"})
	r.Use(static.Serve("/", static.LocalFile("./ui/public", true)))
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
