package routes

// Index/info page

const INDEX = "/"

// Books endpoint

const GetBooks = "/books"
const GetBookWithId = "/books/:id"
const PostBook = "/books"
const PutBook = "/books/:id"
const DeleteBook = "/books/:id"

// Borrowed Books endpoint

const PostBorrowedBooks = "/borrowed"
const DeleteBorrowedBook = "/borrowed/:id"

// Booked Books endpoint

const PostBookedBooks = "/booked"
const DeleteBookedBook = "/booked/:id"

// Users endpoint
