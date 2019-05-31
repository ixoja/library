package storage

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/ixoja/library/internal/models"
	"github.com/patrickmn/go-cache"
)

type Cache struct {
	Storage *cache.Cache
}

func New(cache *cache.Cache) *Cache {
	return &Cache{Storage: cache}
}

func (s *Cache) Create(book *models.Book) (*models.Book, error) {
	book.ID = strfmt.UUID(uuid.New().String())
	s.Storage.SetDefault(string(book.ID), book)
	return book, nil
}

func (s *Cache) Delete(id string) error {
	s.Storage.Delete(id)
	return nil
}

func (s *Cache) Get(id string) (*models.Book, bool, error) {
	book, ok := s.Storage.Get(id)
	if ok {
		return book.(*models.Book), true, nil
	}

	return nil, false, nil
}

func (s *Cache) GetAll() ([]*models.Book, error) {
	m := s.Storage.Items()
	books := make([]*models.Book, 0, len(m))
	for _, value := range m {
		books = append(books, value.Object.(*models.Book))
	}

	return books, nil
}

func (s *Cache) Update(book *models.Book) error {
	s.Storage.SetDefault(string(book.ID), book)
	return nil
}
