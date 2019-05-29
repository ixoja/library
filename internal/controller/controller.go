package controller

import "github.com/ixoja/library/internal/models"

type Storage interface {
	Create(book *models.Book) (*models.Book, error)
	Delete(id string) error
	Get(id string) (*models.Book, error)
	GetAll() ([]*models.Book, error)
	Rate(id, rate string) error
	CheckIn(id string) error
	CheckOut(id string) error
}

type Controller struct {
	Storage Storage
}

func New(s Storage) *Controller {
	return &Controller{Storage: s}
}

func (c Controller) Create(book *models.Book) (*models.Book, error) {
	return nil, nil
}

func (c Controller) Delete(id string) error {
	return nil
}

func (c Controller) Get(id string) (*models.Book, error) {
	return nil, nil
}

func (c Controller) GetAll() ([]*models.Book, error) {
	return nil, nil
}

func (c Controller) Rate(id, rate string) error {
	return nil
}

func (c Controller) Update(id, status string) error {
	return nil
}
