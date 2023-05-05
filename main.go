package main

import (
	"github.com/gin-gonic/gin"
	"togo-web/presentation"
	"togo-web/presentation/todo"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r = presentation.SetRoute(r)
	r = todo.SetRoute(r)
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
