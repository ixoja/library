package controller

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/ixoja/library/internal/controller/mocks"
	"github.com/ixoja/library/internal/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			Title: &title,
			Author: &author,
			Publisher: &publisher,
			PublicationDate: &date,
		}

		s.On("Create", &book).Return(nil, errors.New("message"))
		_, err := c.Create(&book)
		assert.Equal(t, ErrInternal, errors.Cause(err))
	})

	t.Run("success", func(t *testing.T) {
		s := mocks.Storage{}
		c := Controller{&s}
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

		s.On("Create", &book).Return(&retBook, nil)
		res, err := c.Create(&book)
		require.NoError(t, err)
		assert.Equal(t, &retBook, res)
	})
}

func TestController_Delete(t *testing.T) {

}

func TestController_Get(t *testing.T) {

}

func TestController_GetAll(t *testing.T) {

}

func TestController_Rate(t *testing.T) {

}

func TestController_UpdateStatus(t *testing.T) {

}

func id() strfmt.UUID {
	return strfmt.UUID(uuid.New().String())
}
