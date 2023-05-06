package todo

import (
	"errors"
	"fmt"
	"time"
	"togo-web/domain/data"
	"togo-web/domain/data/repository"
)

type GetUseCase interface {
	All() []data.Todo
	FindById(id int) (*data.Todo, error)
	Save(req data.Todo) error
}

type DefaultGetUseCase struct {
	Repository repository.TodoRepository
}

func (useCase DefaultGetUseCase) All() (*[]data.Todo, error) {
	return useCase.Repository.All()
}

func (useCase DefaultGetUseCase) FindById(id string) (*data.Todo, error) {
	return useCase.Repository.FindById(id)
}

func (useCase DefaultGetUseCase) Save(req Request) error {
	createdAt, err := time.Parse("2006-01-02", req.CreatedAt)
	updatedAt, err := time.Parse("2006-01-02", req.UpdatedAt)
	if err != nil {
		return errors.New(fmt.Sprintf("updateAtがパースできません: %s", err))
	}
	return useCase.Repository.Save(data.Todo{
		Id:          req.Id,
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
		State:       data.GetState(req.State),
		CreatedAt:   createdAt,
		CreatedBy:   req.CreatedBy,
		UpdatedAt:   updatedAt,
		UpdatedBy:   req.UpdatedBy,
	})
}
