package types

import (
	"time"
)

type (
	Book struct {
		Id             string          `json:"id"`
		Name           string          `json:"name"`
		Genre          string          `json:"genre"`
		Author         Author          `json:"author"`
		BorrowedStatus *BorrowedStatus `json:"borrowedStatus"`
		BookedStatus   *BookedStatus   `json:"bookedStatus"`
	}
	Author struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		LastName string `json:"lastName"`
	}
	BookedStatus struct {
		To time.Time `json:"to"`
	}
	BorrowedStatus struct {
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	}
)

func (b Book) isBooked() bool {
	if b.BookedStatus != nil {
		return true
	}
	return false
}

func (b Book) isBorrowed() bool {
	if b.BorrowedStatus != nil {
		return true
	}
	return false
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"author"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
