package storage

import (
	"github.com/go-openapi/strfmt"
	"github.com/icrowley/fake"
	"github.com/ixoja/library/internal/models"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCache_Save(t *testing.T) {
	c := Cache{Storage: cache.New(cache.NoExpiration, cache.NoExpiration)}

	title := fake.Title()
	author := fake.FullName()
	publisher := fake.Company()
	date := strfmt.Date(time.Now())

	newBook, err := c.Save(&models.Book{
		Title:           &title,
		Author:          &author,
		Publisher:       &publisher,
		PublicationDate: &date,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, newBook.ID)
}

func TestCache_Get(t *testing.T) {
	c := Cache{Storage: cache.New(cache.NoExpiration, cache.NoExpiration)}

	title := fake.Title()
	author := fake.FullName()
	publisher := fake.Company()
	date := strfmt.Date(time.Now())

	newBook, err := c.Save(&models.Book{
		Title:           &title,
		Author:          &author,
		Publisher:       &publisher,
		PublicationDate: &date,
	})
	require.NoError(t, err)
	b, ok, err := c.Get(newBook.ID.String())
	require.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, newBook, b)
}

func TestCache_GetAll(t *testing.T) {
	t.Run("empty success", func(t *testing.T) {
		c := Cache{Storage: cache.New(cache.NoExpiration, cache.NoExpiration)}
		b, err := c.GetAll()
		require.NoError(t, err)
		assert.Equal(t, make([]*models.Book, 0), b)
	})

	t.Run("not empty success", func(t *testing.T) {
		c := Cache{Storage: cache.New(cache.NoExpiration, cache.NoExpiration)}

		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())
		books := []*models.Book{{
			Title:           &title,
			Author:          &author,
			Publisher:       &publisher,
			PublicationDate: &date,
		},
			{
				Title:           &title,
				Author:          &author,
				Publisher:       &publisher,
				PublicationDate: &date,
			},
		}

		_, err := c.Save(books[0])
		require.NoError(t, err)
		_, err = c.Save(books[1])
		require.NoError(t, err)
		b, err := c.GetAll()
		require.NoError(t, err)
		assert.Equal(t, books, b)
	})
}

func TestCache_Delete(t *testing.T) {
	c := Cache{Storage: cache.New(cache.NoExpiration, cache.NoExpiration)}

	title := fake.Title()
	author := fake.FullName()
	publisher := fake.Company()
	date := strfmt.Date(time.Now())

	newBook, err := c.Save(&models.Book{
		Title:           &title,
		Author:          &author,
		Publisher:       &publisher,
		PublicationDate: &date,
	})
	require.NoError(t, err)
	_, ok, err := c.Get(newBook.ID.String())
	require.NoError(t, err)
	assert.True(t, ok)
	err = c.Delete(newBook.ID.String())
	require.NoError(t, err)
	_, ok, err = c.Get(newBook.ID.String())
	require.NoError(t, err)
	assert.False(t, ok)
}

func TestCache_Update(t *testing.T) {
	t.Run("error empty id", func(t *testing.T) {
		c := Cache{Storage: cache.New(cache.NoExpiration, cache.NoExpiration)}

		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())

		err := c.Update(&models.Book{
			Title:           &title,
			Author:          &author,
			Publisher:       &publisher,
			PublicationDate: &date,
		})
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		c := Cache{Storage: cache.New(cache.NoExpiration, cache.NoExpiration)}

		title := fake.Title()
		author := fake.FullName()
		publisher := fake.Company()
		date := strfmt.Date(time.Now())

		newBook, err := c.Save(&models.Book{
			Title:           &title,
			Author:          &author,
			Publisher:       &publisher,
			PublicationDate: &date,
		})
		require.NoError(t, err)

		updBook := &models.Book{
			ID: newBook.ID,
			Title:           &title,
			Author:          &author,
			Publisher:       &publisher,
			PublicationDate: &date,
			Status: models.BookStatusCheckedIn,
		}
		err = c.Update(updBook)
		require.NoError(t, err)
		book, ok, err := c.Get(newBook.ID.String())
		require.NoError(t, err)
		require.True(t, ok)
		assert.Equal(t, updBook.Status, book.Status)
	})
}
