package handler

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/ixoja/library/internal/controller"
	"github.com/ixoja/library/internal/handler/mocks"
	"github.com/ixoja/library/internal/models"
	"github.com/ixoja/library/internal/restapi/operations"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateBookHandler(t *testing.T) {
	t.Run("bad argument", func(t *testing.T) {
		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())

		for name, tc := range map[string]struct {
			book *models.Book
		}{
			"nil book": { book: nil },
			"nil title": { book: &models.Book{
				Title: nil,
				Author: &author,
				Publisher: &publisher,
				PublicationDate: &date,
			} },
			"nil author": { book: &models.Book{
				Title: &title,
				Author: nil,
				Publisher: &publisher,
				PublicationDate: &date,
			} },
			"nil publisher": { book: &models.Book{
				Title: &title,
				Author: &author,
				Publisher: nil,
				PublicationDate: &date,
			} },
			"nil publication date": { book: &models.Book{
				Title: &title,
				Author: &author,
				Publisher: &publisher,
				PublicationDate: nil,
			} },
		}{
			t.Run(name, func(t *testing.T) {
				h := Handler{}
				res := h.CreateBookHandler(operations.CreateBookParams{Book: tc.book})
				assert.IsType(t, &operations.CreateBookBadRequest{}, res)
			})
		}
	})

	t.Run("internal error", func(t *testing.T) {
		c := mocks.BookController{}
		h := Handler{&c}
		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())
		book := models.Book{
			Title: &title,
			Author: &author,
			Publisher: &publisher,
			PublicationDate: &date,
		}

		c.On("Create", &book).Return(nil, controller.ErrInternal)
		res := h.CreateBookHandler(operations.CreateBookParams{Book: &book})
		assert.IsType(t, &operations.CreateBookInternalServerError{}, res)
		c.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		c := mocks.BookController{}
		h := Handler{&c}
		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())
		book := models.Book{
			Title: &title,
			Author: &author,
			Publisher: &publisher,
			PublicationDate: &date,
		}
		retBook := book
		retBook.ID = id()

		c.On("Create", &book).Return(&retBook, nil)
		res := h.CreateBookHandler(operations.CreateBookParams{Book: &book})
		assert.Equal(t, &operations.CreateBookOK{Payload: &retBook}, res.(*operations.CreateBookOK))
		c.AssertExpectations(t)
	})
}

func TestDeleteBookHandler(t *testing.T) {
	t.Run("not found error", func(t *testing.T) {
		c := mocks.BookController{}
		h := Handler{&c}

		id := id()
		err := controller.ErrNotFound
		c.On("Delete", id.String()).Return(err)
		res := h.DeleteBookHandler(operations.DeleteBookParams{ID: id})
		assert.Equal(t, &operations.DeleteBookNotFound{Payload: &models.Error{
			Message: err.Error()}}, res.(*operations.DeleteBookNotFound))
		c.AssertExpectations(t)
	})

	t.Run("internal error", func(t *testing.T) {
		c := mocks.BookController{}
		h := Handler{&c}

		id := id()
		err := controller.ErrInternal
		c.On("Delete", id.String()).Return(err)
		res := h.DeleteBookHandler(operations.DeleteBookParams{ID: id})
		assert.Equal(t, &operations.DeleteBookInternalServerError{
			Payload: &models.Error{Message: err.Error()}}, res.(*operations.DeleteBookInternalServerError))
		c.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		c := mocks.BookController{}
		h := Handler{&c}

		id := id()
		c.On("Delete", id.String()).Return(nil)
		res := h.DeleteBookHandler(operations.DeleteBookParams{ID: id})
		assert.Equal(t, &operations.DeleteBookOK{}, res.(*operations.DeleteBookOK))
		c.AssertExpectations(t)
	})
}

func TestGetAllBooksHandler(t *testing.T) {
	t.Run("internal error", func(t *testing.T) {
		c := mocks.BookController{}
		h := Handler{&c}

		err := controller.ErrInternal
		c.On("GetAll").Return(nil, err)
		res := h.GetAllBooksHandler(operations.GetAllBooksParams{})
		assert.Equal(t, &operations.GetAllBooksInternalServerError{Payload: &models.Error{Message: err.Error()}},
			res.(*operations.GetAllBooksInternalServerError))
		c.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		c := mocks.BookController{}
		h := Handler{&c}
		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())
		book := models.Book{
			ID: id(),
			Title: &title,
			Author: &author,
			Publisher: &publisher,
			PublicationDate: &date,
		}
		books := []*models.Book{&book}

		c.On("GetAll").Return(books, nil)
		res := h.GetAllBooksHandler(operations.GetAllBooksParams{})
		assert.Equal(t, &operations.GetAllBooksOK{Payload: &operations.GetAllBooksOKBody{Books: books}},
			res.(*operations.GetAllBooksOK))
		c.AssertExpectations(t)
	})
}

func TestGetBookHandler(t *testing.T) {
	t.Run("not found error", func(t *testing.T) {
		c := mocks.BookController{}
		h := Handler{&c}
		id := id()

		err := controller.ErrNotFound
		c.On("Get", id.String()).Return(nil, err)
		res := h.GetBookHandler(operations.GetBookParams{ID: id})
		assert.Equal(t, &operations.GetBookNotFound{Payload: &models.Error{Message: err.Error()}},
			res.(*operations.GetBookNotFound))
		c.AssertExpectations(t)
	})

	t.Run("internal error", func(t *testing.T) {
		c := mocks.BookController{}
		h := Handler{&c}
		id := id()

		err := controller.ErrInternal
		c.On("Get", id.String()).Return(nil, err)
		res := h.GetBookHandler(operations.GetBookParams{ID: id})
		assert.Equal(t, &operations.GetBookInternalServerError{Payload: &models.Error{Message: err.Error()}},
			res.(*operations.GetBookInternalServerError))
		c.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		c := mocks.BookController{}
		h := Handler{&c}
		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())
		id := id()
		book := models.Book{
			ID: id,
			Title: &title,
			Author: &author,
			Publisher: &publisher,
			PublicationDate: &date,
		}

		c.On("Get", id.String()).Return(&book, nil)
		res := h.GetBookHandler(operations.GetBookParams{ID: id})
		assert.Equal(t, &operations.GetBookOK{Payload: &book}, res.(*operations.GetBookOK))
		c.AssertExpectations(t)
	})
}

func TestUpdateBookHandler(t *testing.T) {
	t.Run("no update no error", func(t *testing.T) {
		h := Handler{}
		res := h.UpdateBookHandler(operations.UpdateBookParams{})
		assert.Equal(t, &operations.UpdateBookOK{}, res.(*operations.UpdateBookOK))
	})

	t.Run("rate error", func(t *testing.T) {
		t.Run("not found", func(t *testing.T) {
			c := mocks.BookController{}
			h := Handler{&c}
			id := id()
			rate := 1
			err := controller.ErrNotFound

			c.On("Rate", id.String(), rate).Return(err)
			res := h.UpdateBookHandler(
				operations.UpdateBookParams{ID: id, BookUpdate: operations.UpdateBookBody{Rating: int64(rate)}})
			assert.Equal(t, &operations.UpdateBookNotFound{Payload: &models.Error{Message: err.Error()}},
			res.(*operations.UpdateBookNotFound))
			c.AssertExpectations(t)
		})

		t.Run("internal", func(t *testing.T) {
			c := mocks.BookController{}
			h := Handler{&c}
			id := id()
			rate := 1
			err := controller.ErrInternal

			c.On("Rate", id.String(), rate).Return(err)
			res := h.UpdateBookHandler(
				operations.UpdateBookParams{ID: id, BookUpdate: operations.UpdateBookBody{Rating: int64(rate)}})
			assert.Equal(t, &operations.UpdateBookInternalServerError{Payload: &models.Error{Message: err.Error()}},
				res.(*operations.UpdateBookInternalServerError))
			c.AssertExpectations(t)
		})
	})

	t.Run("update status error", func(t *testing.T) {
		t.Run("not found", func(t *testing.T) {
			c := mocks.BookController{}
			h := Handler{&c}
			id := id()
			err := controller.ErrNotFound
			status := models.BookStatusCheckedIn

			c.On("UpdateStatus", id.String(), status).Return(err)
			res := h.UpdateBookHandler(
				operations.UpdateBookParams{ID: id, BookUpdate: operations.UpdateBookBody{Status: status}})
			assert.Equal(t, &operations.UpdateBookNotFound{Payload: &models.Error{Message: err.Error()}},
				res.(*operations.UpdateBookNotFound))
			c.AssertExpectations(t)
		})

		t.Run("internal", func(t *testing.T) {
			c := mocks.BookController{}
			h := Handler{&c}
			id := id()
			err := controller.ErrInternal
			status := models.BookStatusCheckedIn

			c.On("UpdateStatus", id.String(), status).Return(err)
			res := h.UpdateBookHandler(
				operations.UpdateBookParams{ID: id, BookUpdate: operations.UpdateBookBody{Status: status}})
			assert.Equal(t, &operations.UpdateBookInternalServerError{Payload: &models.Error{Message: err.Error()}},
				res.(*operations.UpdateBookInternalServerError))
			c.AssertExpectations(t)
		})

		t.Run("conflict", func(t *testing.T) {
			c := mocks.BookController{}
			h := Handler{&c}
			id := id()
			err := controller.ErrConflict
			status := models.BookStatusCheckedIn

			c.On("UpdateStatus", id.String(), status).Return(err)
			res := h.UpdateBookHandler(
				operations.UpdateBookParams{ID: id, BookUpdate: operations.UpdateBookBody{Status: status}})
			assert.Equal(t, &operations.UpdateBookConflict{Payload: &models.Error{Message: err.Error()}},
				res.(*operations.UpdateBookConflict))
			c.AssertExpectations(t)
		})
	})
	
	t.Run("rate success", func(t *testing.T) {
		c := mocks.BookController{}
		h := Handler{&c}
		id := id()
		rate := 1

		c.On("Rate", id.String(), rate).Return(nil)
		res := h.UpdateBookHandler(
			operations.UpdateBookParams{ID: id, BookUpdate: operations.UpdateBookBody{Rating: int64(rate)}})
		assert.Equal(t, &operations.UpdateBookOK{}, res.(*operations.UpdateBookOK))
		c.AssertExpectations(t)
	})

	t.Run("update status success", func(t *testing.T) {
		c := mocks.BookController{}
		h := Handler{&c}
		id := id()
		status := models.BookStatusCheckedIn

		c.On("UpdateStatus", id.String(), status).Return(nil)
		res := h.UpdateBookHandler(
			operations.UpdateBookParams{ID: id, BookUpdate: operations.UpdateBookBody{Status: status}})
		assert.Equal(t, &operations.UpdateBookOK{}, res.(*operations.UpdateBookOK))
		c.AssertExpectations(t)
	})
}

func id() strfmt.UUID {
	return strfmt.UUID(uuid.New().String())
}