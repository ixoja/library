package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ixoja/library/internal/controller"
	"github.com/ixoja/library/internal/models"
	"github.com/ixoja/library/internal/restapi/operations"
)

type BookController interface {
	Create(book *models.Book) (*models.Book, error)
	Delete(id string) error
	Get(id string) (*models.Book, error)
	GetAll() ([]*models.Book, error)
	Rate(id string, rate int) error
	UpdateStatus(id, status string) error
}

type Handler struct {
	Controller BookController
}

func New(c BookController) *Handler {
	return &Handler{Controller: c}
}

func (h *Handler) CreateBookHandler(p operations.CreateBookParams) middleware.Responder {
	if p.Body == nil || p.Body.Author == nil ||
		p.Body.PublicationDate == nil || p.Body.Publisher == nil || p.Body.Title == nil {
		return &operations.CreateBookBadRequest{Payload: &models.Error{Message: "book and its fields cannot be null"}}
	}

	newBook, err := h.Controller.Create(&models.Book{
		Title:           p.Body.Title,
		Author:          p.Body.Author,
		Publisher:       p.Body.Publisher,
		PublicationDate: p.Body.PublicationDate,
	})

	if err != nil {
		return &operations.CreateBookInternalServerError{Payload: &models.Error{Message: err.Error()}}
	}
	return &operations.CreateBookOK{Payload: newBook}
}

func (h *Handler) GetAllBooksHandler(p operations.GetAllBooksParams) middleware.Responder {
	books, err := h.Controller.GetAll()
	if err != nil {
		return &operations.GetAllBooksInternalServerError{Payload: &models.Error{Message: err.Error()}}
	}
	return &operations.GetAllBooksOK{Payload: &operations.GetAllBooksOKBody{Books: books}}
}

func (h *Handler) GetBookHandler(p operations.GetBookParams) middleware.Responder {
	book, err := h.Controller.Get(p.ID)
	switch err {
	case nil:
		return &operations.GetBookOK{Payload: book}
	case controller.ErrNotFound:
		return &operations.GetBookNotFound{Payload: &models.Error{Message: err.Error()}}
	default:
		return &operations.GetBookInternalServerError{Payload: &models.Error{Message: err.Error()}}
	}
}

func (h *Handler) UpdateBookHandler(p operations.UpdateBookParams) middleware.Responder {
	if p.BookUpdate.Rating != 0 {
		switch err := h.Controller.Rate(p.ID, int(p.BookUpdate.Rating)); err {
		case nil:
		case controller.ErrNotFound:
			return &operations.UpdateBookNotFound{Payload: &models.Error{Message: err.Error()}}
		default:
			return &operations.UpdateBookInternalServerError{Payload: &models.Error{Message: err.Error()}}
		}
	}

	if p.BookUpdate.Status != "" {
		switch err := h.Controller.UpdateStatus(p.ID, p.BookUpdate.Status); err {
		case nil:
		case controller.ErrNotFound:
			return &operations.UpdateBookNotFound{Payload: &models.Error{Message: err.Error()}}
		default:
			return &operations.UpdateBookInternalServerError{Payload: &models.Error{Message: err.Error()}}
		}
	}

	return &operations.UpdateBookOK{}
}

func (h *Handler) DeleteBookHandler(p operations.DeleteBookParams) middleware.Responder {
	switch err := h.Controller.Delete(p.ID); err {
	case nil:
		return &operations.DeleteBookOK{}
	case controller.ErrNotFound:
		return &operations.DeleteBookNotFound{Payload: &models.Error{Message: err.Error()}}
	default:
		return &operations.DeleteBookInternalServerError{Payload: &models.Error{Message: err.Error()}}
	}
}
