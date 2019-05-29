// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Book book
// swagger:model book
type Book struct {

	// author
	// Required: true
	Author *string `json:"author"`

	// id
	ID string `json:"id,omitempty"`

	// publication date
	// Required: true
	PublicationDate *string `json:"publication_date"`

	// publisher
	// Required: true
	Publisher *string `json:"publisher"`

	// rating
	// Enum: [rate1 rate2 rate3]
	Rating string `json:"rating,omitempty"`

	// status
	// Enum: [checked_in checked_out]
	Status string `json:"status,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this book
func (m *Book) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublisher(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRating(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Book) validateAuthor(formats strfmt.Registry) error {

	if err := validate.Required("author", "body", m.Author); err != nil {
		return err
	}

	return nil
}

func (m *Book) validatePublicationDate(formats strfmt.Registry) error {

	if err := validate.Required("publication_date", "body", m.PublicationDate); err != nil {
		return err
	}

	return nil
}

func (m *Book) validatePublisher(formats strfmt.Registry) error {

	if err := validate.Required("publisher", "body", m.Publisher); err != nil {
		return err
	}

	return nil
}

var bookTypeRatingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rate1","rate2","rate3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bookTypeRatingPropEnum = append(bookTypeRatingPropEnum, v)
	}
}

const (

	// BookRatingRate1 captures enum value "rate1"
	BookRatingRate1 string = "rate1"

	// BookRatingRate2 captures enum value "rate2"
	BookRatingRate2 string = "rate2"

	// BookRatingRate3 captures enum value "rate3"
	BookRatingRate3 string = "rate3"
)

// prop value enum
func (m *Book) validateRatingEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bookTypeRatingPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Book) validateRating(formats strfmt.Registry) error {

	if swag.IsZero(m.Rating) { // not required
		return nil
	}

	// value enum
	if err := m.validateRatingEnum("rating", "body", m.Rating); err != nil {
		return err
	}

	return nil
}

var bookTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["checked_in","checked_out"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bookTypeStatusPropEnum = append(bookTypeStatusPropEnum, v)
	}
}

const (

	// BookStatusCheckedIn captures enum value "checked_in"
	BookStatusCheckedIn string = "checked_in"

	// BookStatusCheckedOut captures enum value "checked_out"
	BookStatusCheckedOut string = "checked_out"
)

// prop value enum
func (m *Book) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bookTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Book) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Book) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Book) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Book) UnmarshalBinary(b []byte) error {
	var res Book
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
