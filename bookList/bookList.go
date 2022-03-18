package bookList

var books []Book

func init() {
	books = CreateList()
}

// CreateList : creates list of books in the library
func CreateList() []Book {
	authors := createAuthorData()
	books := createBookData(authors)
	return books
}
