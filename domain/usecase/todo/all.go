package todo

import (
	"togo-web/domain/data"
	"togo-web/domain/data/repository"
)

type GetUseCase interface {
	All() []data.Todo
	FindById(id int) (*data.Todo, error)
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
