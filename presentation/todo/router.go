package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	error2 "togo-web/application/data/error"
	"togo-web/domain/data/repository"
	"togo-web/domain/usecase/todo"
	"togo-web/infrastructure/firestore"
)

func SetRoute(r *gin.Engine) *gin.Engine {
	r.GET("/todo", func(c *gin.Context) {
		client, err := firestore.InitializeFirestore()
		if err != nil {
			c.JSON(http.StatusInternalServerError, error2.ApplicationError{
				Code:    "",
				Message: "予期しないエラーが発生しました。少し経ってからお試しください。",
			})
		}
		data, err := todo.DefaultGetUseCase{Repository: repository.DefaultTodoRepository{Client: client}}.All()
		if err != nil {
			c.JSON(http.StatusNotFound, error2.ApplicationError{
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
			c.JSON(http.StatusBadRequest, error2.ApplicationError{
				Code:    "",
				Message: "idは整数を指定してください",
			})
			return
		}
		client, err := firestore.InitializeFirestore()
		if err != nil {
			c.JSON(http.StatusInternalServerError, error2.ApplicationError{
				Code:    "",
				Message: "予期しないエラーが発生しました。少し経ってからお試しください。",
			})
		}
		var data, fetchErr = todo.DefaultGetUseCase{Repository: repository.DefaultTodoRepository{Client: client}}.FindById(id)
		if fetchErr != nil {
			c.JSON(http.StatusNotFound, error2.ApplicationError{
				Code:    "",
				Message: fetchErr.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, data)
	})
	return r
}
