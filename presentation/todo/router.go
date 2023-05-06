package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	e "togo-web/application/data/error"
	"togo-web/domain/data/repository"
	"togo-web/domain/usecase/todo"
	"togo-web/infrastructure/firestore"
)

func SetRoute(r *gin.Engine) *gin.Engine {
	r.GET("/todo", func(c *gin.Context) {
		client, err := firestore.InitializeFirestore()
		if err != nil {
			c.JSON(http.StatusInternalServerError, e.ApplicationError{
				Code:    "",
				Message: "予期しないエラーが発生しました。少し経ってからお試しください。",
			})
			return
		}
		data, err := todo.DefaultGetUseCase{Repository: repository.DefaultTodoRepository{Client: client}}.All()
		if err != nil {
			c.JSON(http.StatusNotFound, e.ApplicationError{
				Code:    "",
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, data)
	})
	r.GET("/todo/:id", func(c *gin.Context) {
		id, paramErr := c.Params.Get("id")
		if !paramErr {
			c.JSON(http.StatusBadRequest, e.ApplicationError{
				Code:    "",
				Message: "idは整数を指定してください",
			})
			return
		}
		client, err := firestore.InitializeFirestore()
		if err != nil {
			c.JSON(http.StatusInternalServerError, e.ApplicationError{
				Code:    "",
				Message: "予期しないエラーが発生しました。少し経ってからお試しください。",
			})
			return
		}
		var data, fetchErr = todo.DefaultGetUseCase{Repository: repository.DefaultTodoRepository{Client: client}}.FindById(id)
		if fetchErr != nil {
			c.JSON(http.StatusNotFound, e.ApplicationError{
				Code:    "",
				Message: fetchErr.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, data)
	})
	r.POST("/todo", func(c *gin.Context) {
		var req todo.Request
		if c.Bind(&req) != nil {
			panic(c.Bind(&req))
		}
		client, err := firestore.InitializeFirestore()
		if err != nil {
			c.JSON(http.StatusInternalServerError, e.ApplicationError{
				Code:    "",
				Message: "予期しないエラーが発生しました。少し経ってからお試しください。",
			})
			return
		}
		err = todo.DefaultGetUseCase{Repository: repository.DefaultTodoRepository{Client: client}}.Save(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, e.ApplicationError{
				Code:    "",
				Message: "データの登録に失敗しました。",
			})
			return
		}
		c.JSON(http.StatusOK, req)
	})
	return r
}
