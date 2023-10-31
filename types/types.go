package types

import (
	"time"
)

type Book struct {
	Id             int             `json:"id"`
	Name           string          `json:"name"`
	Author         string          `json:"author"`
	BorrowedStatus *BorrowedStatus `json:"borrowedStatus"`
	BookedStatus   *BookedStatus   `json:"bookedStatus"`
}

func (b Book) isBorrowed() bool {
	if b.BorrowedStatus != nil {
		return true
	}
	return false
}

type BorrowedStatus struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type BookedStatus struct {
	To time.Time `json:"to"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"author"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
