// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/ixoja/library/internal/models"
)

// GetAllBooksOKCode is the HTTP code returned for type GetAllBooksOK
const GetAllBooksOKCode int = 200

/*GetAllBooksOK OK

swagger:response getAllBooksOK
*/
type GetAllBooksOK struct {

	/*
	  In: Body
	*/
	Payload *GetAllBooksOKBody `json:"body,omitempty"`
}

// NewGetAllBooksOK creates GetAllBooksOK with default headers values
func NewGetAllBooksOK() *GetAllBooksOK {

	return &GetAllBooksOK{}
}

// WithPayload adds the payload to the get all books o k response
func (o *GetAllBooksOK) WithPayload(payload *GetAllBooksOKBody) *GetAllBooksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all books o k response
func (o *GetAllBooksOK) SetPayload(payload *GetAllBooksOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllBooksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAllBooksInternalServerErrorCode is the HTTP code returned for type GetAllBooksInternalServerError
const GetAllBooksInternalServerErrorCode int = 500

/*GetAllBooksInternalServerError Internal error.

swagger:response getAllBooksInternalServerError
*/
type GetAllBooksInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllBooksInternalServerError creates GetAllBooksInternalServerError with default headers values
func NewGetAllBooksInternalServerError() *GetAllBooksInternalServerError {

	return &GetAllBooksInternalServerError{}
}

// WithPayload adds the payload to the get all books internal server error response
func (o *GetAllBooksInternalServerError) WithPayload(payload *models.Error) *GetAllBooksInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all books internal server error response
func (o *GetAllBooksInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllBooksInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
