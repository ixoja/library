package model

import "time"

type Book struct {
	ID string
	Title string
	Author string
	Publisher string
	PublicationDate time.Time
	Rating Rating
	Status Status
}

type Rating int

const (
	Rating1 Rating = 1 + iota
	Rating2
	Rating3
)

type Status int

const (
	StatusCheckedIn Status = iota
	StatusCheckedOut
)
