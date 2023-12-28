package types

import (
	"github.com/rickb777/date"
)

type (
	BookDto struct {
		Id             string          `json:"id"`
		Title          string          `json:"title"`
		Genre          string          `json:"genre"`
		Authors        []Author        `json:"authors,omitempty"`
		BorrowedStatus *BorrowedStatus `json:"borrowedStatus,omitempty"`
		BookedStatus   *BookedStatus   `json:"bookedStatus,omitempty"`
	}
	Author struct {
		Id       string `json:"id"`
		Name     string `json:"name,omitempty"`
		LastName string `json:"lastName,omitempty"`
	}
	BookedStatus struct {
		BookedBy *string    `json:"bookedBy,omitempty"`
		To       *date.Date `json:"to,omitempty"`
	}
	BorrowedStatus struct {
		BorrowedBy *string    `json:"borrowedBy,omitempty"`
		From       *date.Date `json:"from,omitempty"`
		To         *date.Date `json:"to,omitempty"`
	}
	User struct {
		Name     string `json:"name"`
		Lastname string `json:"lastname"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	Credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)

func (b BookDto) IsBooked() bool {
	if b.BookedStatus.BookedBy != nil {
		return true
	}
	return false
}

func (b BookDto) IsBorrowed() bool {
	if b.BorrowedStatus.BorrowedBy != nil {
		return true
	}
	return false
}
