package types

import (
	"time"
)

type (
	BookDto struct {
		Id             string          `json:"id"`
		Title          string          `json:"title"`
		Genre          string          `json:"genre"`
		Authors        []Author        `json:"authors"`
		BorrowedStatus *BorrowedStatus `json:"borrowedStatus"`
		BookedStatus   *BookedStatus   `json:"bookedStatus"`
	}
	Author struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		LastName string `json:"lastName"`
	}
	BookedStatus struct {
		BookedBy string     `json:"bookedBy"`
		To       *time.Time `json:"to"`
	}
	BorrowedStatus struct {
		BorrowedBy string     `json:"borrowedBy"`
		From       *time.Time `json:"from"`
		To         *time.Time `json:"to"`
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
