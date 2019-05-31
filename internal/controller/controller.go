package controller

import (
	"github.com/ixoja/library/internal/models"
	"github.com/pkg/errors"
)

//go:generate mockery -case=underscore -name Storage
type Storage interface {
	Create(book *models.Book) (*models.Book, error)
	Delete(id string) error
	Get(id string) (*models.Book, bool, error)
	GetAll() ([]*models.Book, error)
	Update(book *models.Book) error
}

type Controller struct {
	Storage Storage
}

func New(s Storage) *Controller {
	return &Controller{Storage: s}
}

func (c Controller) Create(book *models.Book) (*models.Book, error) {
	newBook, err := c.Storage.Create(book)
	if err != nil {
		return nil, errors.Wrap(ErrInternal, err.Error())
	}
	return newBook, nil
}

func (c Controller) Delete(id string) error {
	_, ok, err := c.Storage.Get(id)
	if err != nil {
		return errors.Wrap(err, "failed to get book")
	}
	if !ok {
		return ErrNotFound
	}

	if err := c.Storage.Delete(id); err != nil {
		return errors.Wrap(ErrInternal, err.Error())
	}
	return nil
}

func (c Controller) Get(id string) (*models.Book, error) {
	book, ok, err := c.Storage.Get(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get book")
	}
	if !ok {
		return nil, ErrNotFound
	}
	return book, nil
}

func (c Controller) GetAll() ([]*models.Book, error) {
	books, err := c.Storage.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all books")
	}

	return books, nil
}

func (c Controller) Rate(id string, rate int64) error {
	book, ok, err := c.Storage.Get(id)
	if err != nil {
		return errors.Wrap(err, "failed to get book")
	}
	if !ok {
		return ErrNotFound
	}

	if book.Rating == nil {
		book.Rating = &models.BookRating{}
	}

	book.Rating.RatesCount++
	book.Rating.RatePrecise = avgRate(book.Rating.RatePrecise, float64(book.Rating.RatesCount), float64(rate))
	switch int(book.Rating.RatePrecise) {
	case 1:
		book.Rating.Rate = models.BookRatingRateRate1
	case 2:
		book.Rating.Rate = models.BookRatingRateRate2
	case 3:
		book.Rating.Rate = models.BookRatingRateRate3
	}

	if err := c.Storage.Update(book); err != nil {
		return errors.Wrap(err, "failed to rate book")
	}

	return nil
}

func avgRate(currentRate, count, newRate float64) float64 {
	return float64((count-1)/count*currentRate + newRate/count)
}

func (c Controller) UpdateStatus(id, status string) error {
	book, ok, err := c.Storage.Get(id)
	if err != nil {
		return errors.Wrap(err, "failed to get book")
	}
	if !ok {
		return ErrNotFound
	}

	if book.Status != status {
		book.Status = status
		if err := c.Storage.Update(book); err != nil {
			return errors.Wrap(err, "failed to update book")
		}
	} else if book.Status == models.BookStatusCheckedIn {
		return errors.Wrap(ErrConflict, "failed to check in book: it's already checked in")
	}

	return nil
}
