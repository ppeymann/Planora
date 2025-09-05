package service

import (
	"context"
	"errors"

	todopb "github.com/ppeymann/Planora.git/proto/todo"
	"github.com/ppeymann/Planora/todo/models"
)

type TodoServiceServer struct {
	todopb.UnimplementedTodoServiceServer
	repo models.TodoRepository
}

func NewTodoServiceServer(r models.TodoRepository) *TodoServiceServer {
	return &TodoServiceServer{
		repo: r,
	}
}

func (s *TodoServiceServer) AddTodo(_ context.Context, in *todopb.AddTodoRequest) (*todopb.Todo, error) {
	todo, err := s.repo.Create(in)
	if err != nil {
		return nil, err
	}

	return &todopb.Todo{
		Model:       models.ToBaseModel(todo),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      string(todo.Status),
		UserId:      uint64(todo.UserID),
	}, nil
}

func (s *TodoServiceServer) UpdateTodo(_ context.Context, in *todopb.UpdateTodoRequest) (*todopb.Todo, error) {
	todo, err := s.repo.FindByID(uint(in.GetId()))
	if err != nil {
		return nil, err
	}

	if todo.UserID != uint(in.Todo.GetUserId()) {
		return nil, errors.New("permission denied")
	}

	todo.Title = in.Todo.GetTitle()
	todo.Description = in.Todo.GetDescription()

	err = s.repo.Update(todo)
	if err != nil {
		return nil, err
	}

	return &todopb.Todo{
		Model:       models.ToBaseModel(todo),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      string(todo.Status),
		UserId:      uint64(todo.UserID),
	}, nil
}

func (s *TodoServiceServer) GetAllTodo(ctx context.Context, in *todopb.GetAllTodoRequest) (*todopb.GetAllTodoResponse, error) {
	todos, err := s.repo.FindAllTodo(uint(in.GetUserId()))
	if err != nil {
		return nil, err
	}

	var todoResponse []*todopb.Todo
	for _, todo := range todos {
		t := todopb.Todo{
			Model:       models.ToBaseModel(&todo),
			Title:       todo.Title,
			Description: todo.Description,
			Status:      string(todo.Status),
			UserId:      uint64(todo.UserID),
		}

		todoResponse = append(todoResponse, &t)
	}

	return &todopb.GetAllTodoResponse{
		Todos: todoResponse,
	}, nil
}
