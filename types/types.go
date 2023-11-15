package types

import (
	"time"
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
		To       *time.Time `json:"to,omitempty"`
	}
	BorrowedStatus struct {
		BorrowedBy *string    `json:"borrowedBy,omitempty"`
		From       *time.Time `json:"from,omitempty"`
		To         *time.Time `json:"to,omitempty"`
	}
	User struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Surname  string `json:"author"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}
)

func (b BookDto) isBooked() bool {
	if b.BookedStatus != nil {
		return true
	}
	return false
}

func (b BookDto) isBorrowed() bool {
	if b.BorrowedStatus != nil {
		return true
	}
	return false
}
