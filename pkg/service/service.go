package service

import "github.com/m0n7h0ff/course-todo-app/pkg/repository"

//интерфейсы сущностей, название - участок бизнес логики , за который они отвечают
type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
