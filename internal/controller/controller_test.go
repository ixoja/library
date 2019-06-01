package controller

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/ixoja/library/internal/controller/mocks"
	"github.com/ixoja/library/internal/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func TestController_Create(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}
		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())
		book := models.Book{
			Title:           &title,
			Author:          &author,
			Publisher:       &publisher,
			PublicationDate: &date,
		}

		s.On("Save", &book).Return(nil, errors.New("message"))
		_, err := c.Create(&book)
		assert.Equal(t, ErrInternal, errors.Cause(err))
		s.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}
		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())
		book := models.Book{
			Title:           &title,
			Author:          &author,
			Publisher:       &publisher,
			PublicationDate: &date,
		}
		retBook := book
		retBook.ID = id()

		s.On("Save", &book).Return(&retBook, nil)
		res, err := c.Create(&book)
		require.NoError(t, err)
		assert.Equal(t, &retBook, res)
		s.AssertExpectations(t)
	})
}

func TestController_Delete(t *testing.T) {
	t.Run("internal error", func(t *testing.T) {
		t.Run("get", func(t *testing.T) {
			s := mocks.Storage{}
			c := Controller{&s}
			id := id().String()

			s.On("Get", id).Return(nil, false, errors.New("storage error"))
			err := c.Delete(id)
			assert.Equal(t, ErrInternal, errors.Cause(err))
			s.AssertExpectations(t)
		})

		t.Run("delete", func(t *testing.T) {
			s := mocks.Storage{}
			c := Controller{&s}
			id := id().String()

			s.On("Get", id).Return(nil, true, nil)
			s.On("Delete", id).Return(errors.New("storage error"))
			err := c.Delete(id)
			assert.Equal(t, ErrInternal, errors.Cause(err))
			s.AssertExpectations(t)
		})
	})

	t.Run("not found", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}
		id := id().String()

		s.On("Get", id).Return(nil, false, nil)
		err := c.Delete(id)
		assert.Equal(t, ErrNotFound, errors.Cause(err))
		s.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}
		id := id().String()

		s.On("Get", id).Return(nil, true, nil)
		s.On("Delete", id).Return(nil)
		err := c.Delete(id)
		assert.NoError(t, err)
		s.AssertExpectations(t)
	})
}

func TestController_Get(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		for name, tc := range map[string]struct {
			haveErr error
			wantErr error
		}{
			"storage": {
				haveErr: errors.New("storage error"),
				wantErr: ErrInternal,
			},

			"not found": {
				haveErr: nil,
				wantErr: ErrNotFound,
			},
		} {
			t.Run(name, func(t *testing.T) {
				s := mocks.Storage{}
				c := Controller{&s}
				id := id().String()

				s.On("Get", id).Return(nil, false, tc.haveErr)
				_, err := c.Get(id)
				assert.Equal(t, tc.wantErr, errors.Cause(err))
				s.AssertExpectations(t)
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}
		id := id().String()
		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())
		book := &models.Book{
			ID:              strfmt.UUID(id),
			Title:           &title,
			Author:          &author,
			Publisher:       &publisher,
			PublicationDate: &date,
		}

		s.On("Get", id).Return(book, true, nil)
		res, err := c.Get(id)
		require.NoError(t, err)
		assert.Equal(t, book, res)
		s.AssertExpectations(t)
	})
}

func TestController_GetAll(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}

		s.On("GetAll").Return(nil, errors.New("storage error"))
		_, err := c.GetAll()
		assert.Equal(t, ErrInternal, errors.Cause(err))
		s.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}
		id := id().String()
		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())

		books := []*models.Book{{
			ID:              strfmt.UUID(id),
			Title:           &title,
			Author:          &author,
			Publisher:       &publisher,
			PublicationDate: &date,
		}}

		s.On("GetAll").Return(books, nil)
		res, err := c.GetAll()
		require.NoError(t, err)
		assert.Equal(t, books, res)
		s.AssertExpectations(t)
	})
}

func TestController_Rate(t *testing.T) {
	t.Run("internal error", func(t *testing.T) {
		t.Run("get", func(t *testing.T) {
			s := mocks.Storage{}
			c := Controller{&s}
			id := id().String()

			s.On("Get", id).Return(nil, false, errors.New("storage error"))
			err := c.Rate(id, randRate())
			assert.Equal(t, ErrInternal, errors.Cause(err))
			s.AssertExpectations(t)
		})

		t.Run("update", func(t *testing.T) {
			s := mocks.Storage{}
			c := Controller{&s}
			id := id().String()

			s.On("Get", id).Return(&models.Book{}, true, nil)
			s.On("Update", mock.Anything).Return(errors.New("storage error"))
			err := c.Rate(id, randRate())
			assert.Equal(t, ErrInternal, errors.Cause(err))
			s.AssertExpectations(t)
		})
	})

	t.Run("not found", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}
		id := id().String()

		s.On("Get", id).Return(nil, false, nil)
		err := c.Rate(id, randRate())
		assert.Equal(t, ErrNotFound, errors.Cause(err))
		s.AssertExpectations(t)

	})

	t.Run("success", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}
		id := id().String()
		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())
		rand.Seed(time.Now().UnixNano())
		ratePrecise := rand.Float64()*2 + 1
		newRate := randRate()

		book := &models.Book{
			ID:              strfmt.UUID(id),
			Title:           &title,
			Author:          &author,
			Publisher:       &publisher,
			PublicationDate: &date,
			Rating: &models.BookRating{
				RatesCount: int64(rand.Intn(20)),
				RatePrecise: ratePrecise,
				Rate: rateString(ratePrecise),
			},
		}

		s.On("Get", id).Return(book, true, nil)
		s.On("Update", book).Return(nil)
		err := c.Rate(id, newRate)
		assert.NoError(t, err)
		s.AssertExpectations(t)
	})
}

func TestController_UpdateStatus(t *testing.T) {
	t.Run("internal error", func(t *testing.T) {
		t.Run("get", func(t *testing.T) {
			s := mocks.Storage{}
			c := Controller{&s}
			id := id().String()

			s.On("Get", id).Return(nil, false, errors.New("storage error"))
			err := c.UpdateStatus(id, models.BookStatusCheckedOut)
			assert.Equal(t, ErrInternal, errors.Cause(err))
			s.AssertExpectations(t)
		})

		t.Run("update", func(t *testing.T) {
			s := mocks.Storage{}
			c := Controller{&s}
			id := id().String()

			s.On("Get", id).Return(&models.Book{}, true, nil)
			s.On("Update", mock.Anything).Return(errors.New("storage error"))
			err := c.UpdateStatus(id, models.BookStatusCheckedOut)
			assert.Equal(t, ErrInternal, errors.Cause(err))
			s.AssertExpectations(t)
		})
	})

	t.Run("not found", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}
		id := id().String()

		s.On("Get", id).Return(nil, false, nil)
		err := c.UpdateStatus(id, models.BookStatusCheckedOut)
		assert.Equal(t, ErrNotFound, errors.Cause(err))
		s.AssertExpectations(t)
	})

	t.Run("already checked in error", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}
		id := id().String()
		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())

		book := &models.Book{
			ID:              strfmt.UUID(id),
			Title:           &title,
			Author:          &author,
			Publisher:       &publisher,
			PublicationDate: &date,
			Status: models.BookStatusCheckedIn,
		}

		s.On("Get", id).Return(book, true, nil)
		err := c.UpdateStatus(id, models.BookStatusCheckedIn)
		assert.Equal(t, ErrConflict, errors.Cause(err))
		s.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}
		id := id().String()
		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())

		book := &models.Book{
			ID:              strfmt.UUID(id),
			Title:           &title,
			Author:          &author,
			Publisher:       &publisher,
			PublicationDate: &date,
			Status: models.BookStatusCheckedIn,
		}

		s.On("Get", id).Return(book, true, nil)
		s.On("Update", book).Return(nil)
		err := c.UpdateStatus(id, models.BookStatusCheckedOut)
		assert.NoError(t, err)
		s.AssertExpectations(t)
	})
}

func id() strfmt.UUID {
	return strfmt.UUID(uuid.New().String())
}

func randRate() int64 {
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(2) + 1)
}
