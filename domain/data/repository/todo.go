package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"time"
	"togo-web/domain/data"
)

var ctx = context.Background()

type TodoRepository interface {
	All() (*[]data.Todo, error)
	FindById(id string) (*data.Todo, error)
	Save(req data.Todo) error
}

type DefaultTodoRepository struct {
	Client *firestore.Client
}

func (repository DefaultTodoRepository) All() (*[]data.Todo, error) {
	iter := repository.Client.Collection("todos").Documents(ctx)
	var result []data.Todo
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		v, err := docToTodo(doc.Data())
		if err != nil {
			return nil, err
		}
		result = append(result, *v)
	}
	return &result, nil
}

// FindById は、idで特定できるTodoの参照を返す(errorの有無によりnullableになる)
func (repository DefaultTodoRepository) FindById(id string) (*data.Todo, error) {
	iter := repository.Client.Collection("todos").Where("id", "==", id).Documents(ctx)
	var result []data.Todo
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		v, err := docToTodo(doc.Data())
		if err != nil {
			return nil, err
		}
		result = append(result, *v)
	}
	return &result[0], nil
}

func (repository DefaultTodoRepository) Save(req data.Todo) error {
	_, _, err := repository.Client.Collection("todos").Add(ctx, todoToMap(req))
	if err != nil {
		return err
	}
	return nil
}

func todoToMap(todo data.Todo) map[string]interface{} {
	var v = make(map[string]interface{})
	v["id"] = todo.Id
	v["userId"] = todo.UserId
	v["title"] = todo.Title
	v["state"] = todo.State.String()
	v["description"] = todo.Description
	v["createdAt"] = todo.CreatedAt.Format("2006-01-02")
	v["createdBy"] = todo.CreatedBy
	v["updatedAt"] = todo.UpdatedAt.Format("2006-01-02")
	v["updatedBy"] = todo.UpdatedBy
	return v
}

func docToTodo(doc map[string]interface{}) (*data.Todo, error) {
	createdAt, err := time.Parse("2006-01-02", doc["createdAt"].(string))
	updatedAt, err := time.Parse("2006-01-02", doc["updatedAt"].(string))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("updateAtがパースできません: %s", err))
	}
	return &data.Todo{
		Id:          doc["id"].(string),
		UserId:      doc["userId"].(string),
		Title:       doc["title"].(string),
		Description: doc["description"].(string),
		State:       data.GetState(doc["state"].(string)),
		CreatedAt:   createdAt,
		CreatedBy:   doc["createdBy"].(string),
		UpdatedAt:   updatedAt,
		UpdatedBy:   doc["updatedBy"].(string),
	}, nil
}
