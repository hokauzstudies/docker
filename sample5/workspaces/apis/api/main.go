package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func api() {
	gin.Default()
}

func main() {
	api()
	fmt.Println("Start api")
}
