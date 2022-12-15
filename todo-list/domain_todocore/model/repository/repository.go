package repository

import (
	"context"
	"todo-list/domain_todocore/model/entity"
	"todo-list/domain_todocore/model/vo"
)

type SaveTodoRepo interface {
	SaveTodo(ctx context.Context, obj *entity.Todo) error
}

type FindAllTodoRepo interface {
	FindAllTodo(ctx context.Context, page, size int64) ([]*entity.Todo, int64, error)
}

type FindOneTodoByIDRepo interface {
	FindOneTodoByID(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error)
}
