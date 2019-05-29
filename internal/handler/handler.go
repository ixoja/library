package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ixoja/library/internal/models"
	"github.com/ixoja/library/internal/restapi/operations"
)

type BookController interface{
	Create(book *models.Book) (*models.Book, error)
	Delete(id string) error
	Get(id string) (*models.Book, error)
	GetAll() ([]*models.Book, error)
	Rate(id, rate string) error
	Update(id, status string) error
}

type Handler struct {
	Controller *BookController
}

func New(c BookController) *Handler {
	return &Handler{Controller: &c}
}

func (h *Handler) CreateBookHandler(p operations.CreateBookParams) middleware.Responder {

	return &operations.CreateBookOK{Payload: &models.Book{}}
}

func (h *Handler) DeleteBookHandler(p operations.DeleteBookParams) middleware.Responder {
	return &operations.DeleteBookOK{}
}

func (h *Handler) GetAllBooksHandler(p operations.GetAllBooksParams) middleware.Responder {
	return &operations.GetAllBooksOK{Payload: &operations.GetAllBooksOKBody{Books: []*models.Book{}}}
}

func (h *Handler) GetBookHandler(p operations.GetBookParams) middleware.Responder {
	return &operations.GetBookOK{Payload: &models.Book{}}
}

func (h *Handler) UpdateBookHandler(p operations.UpdateBookParams) middleware.Responder {
	return &operations.UpdateBookOK{}
}
