package routes

// Index/info page

const INDEX = "/"

// Books endpoint

const GetBooks = "/books"
const GetBookWithId = "/books/:id"
const PostBook = "/books"
const PutBook = "/books/:id"
const DeleteBook = "/books/:id"

// Authors endpoint

const GetAuthors = "/authors"
const GetAuthorWithId = "/authors/:id"
const PostAuthor = "/authors"

// Borrowed Books endpoint

const PostBorrowedBook = "/borrowed/:id"
const DeleteBorrowedBook = "/borrowed/:user/:id"

// Booked Books endpoint

const PostBookedBooks = "/booked/:id"
const DeleteBookedBook = "/booked/:user/:id"

// Archive endpoint

const GetArchive = "/archive/:user"

// Users endpoint

const PostLoginUsers = "/users"
