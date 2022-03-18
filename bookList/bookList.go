package bookList

type Book struct {
	ID          string
	Name        string
	PageNumber  uint
	StockNumber int
	Price       float32
	StockID     string
	ISBN        string
	IsDeleted   bool
	Author      Author
}

type Author struct {
	ID   string
	Name string
}

type BookData struct {
	IDs          []string
	Names        []string
	PageNumbers  []uint
	StockNumbers []int
	Prices       []float32
	StockIDs     []string
	ISBNs        []string
	Authors      []Author
}
type AuthorData struct {
	IDs   []string
	Names []string
}

var authors []Author
var books []Book
var dataLength = 5

func CreateList() []Book {
	// dataLength := 5

	// authorData := AuthorData{
	// 	IDs:   []string{"101", "202", "303", "404", "505"},
	// 	Names: []string{"Charles Dickens", "J. R. R. Tolkien", "J. K. Rowling", "Antoine de Saint-Exupéry", "Cao Xueqin"},
	// }

	// for i := 0; i < dataLength; i++ {
	// 	a := *generateAuthor(authorData, i)
	// 	authors = append(authors, a)
	// }
	authors := createAuthorData()
	books := createBookData(authors)
	return books
	// bookData := BookData{
	// 	IDs:          []string{"1", "2", "3", "4", "5"},
	// 	Names:        []string{"A Tale of Two Cities", "The Hobbit", "Harry Potter and the Philosophers Stone", "The Little Prince", "Dream of the Red Chamber"},
	// 	PageNumbers:  []uint{320, 376, 560, 102, 350},
	// 	StockNumbers: []int{10, 10, 10, 10, 0},
	// 	Prices:       []float32{15.30, 24.00, 32.20, 7.80, 17.00},
	// 	StockIDs:     []string{"21AC", "44UY", "22OL", "09UJ", "77II"},
	// 	ISBNs: []string{"9780451530578", "9780547928227", "9781408855898", "9781853261589",
	// 		"9780385093798"},
	// 	Authors: authors,
	// }

	// for i := 0; i < dataLength; i++ {
	// 	b := *generateBook(bookData, i)
	// 	books = append(books, b)
	// }
	// return books
}
func createAuthorData() []Author {
	authorData := AuthorData{
		IDs:   []string{"101", "202", "303", "404", "505"},
		Names: []string{"Charles Dickens", "J. R. R. Tolkien", "J. K. Rowling", "Antoine de Saint-Exupéry", "Cao Xueqin"},
	}

	for i := 0; i < dataLength; i++ {
		a := *generateAuthor(authorData, i)
		authors = append(authors, a)
	}
	return authors
}
func createBookData(authors []Author) []Book {
	bookData := BookData{
		IDs:          []string{"1", "2", "3", "4", "5"},
		Names:        []string{"A Tale of Two Cities", "The Hobbit", "Harry Potter and the Philosophers Stone", "The Little Prince", "Dream of the Red Chamber"},
		PageNumbers:  []uint{320, 376, 560, 102, 350},
		StockNumbers: []int{10, 10, 10, 10, 0},
		Prices:       []float32{15.30, 24.00, 32.20, 7.80, 17.00},
		StockIDs:     []string{"21AC", "44UY", "22OL", "09UJ", "77II"},
		ISBNs: []string{"9780451530578", "9780547928227", "9781408855898", "9781853261589",
			"9780385093798"},
		Authors: authors,
	}

	for i := 0; i < dataLength; i++ {
		b := *generateBook(bookData, i)
		books = append(books, b)
	}
	return books
}

// generateBook : takes a book data struct as an input and create books with specified attributes given in this struct.
func generateBook(data BookData, index int) *Book {

	book := Book{
		ID:          data.IDs[index],
		Name:        data.Names[index],
		PageNumber:  data.PageNumbers[index],
		StockNumber: data.StockNumbers[index],
		Price:       data.Prices[index],
		StockID:     data.StockIDs[index],
		ISBN:        data.ISBNs[index],
		Author:      data.Authors[index],
	}
	return &book
}

// generateBook : takes an author data struct as an input and create authors with specified attributes given in this struct.
func generateAuthor(data AuthorData, index int) *Author {

	author := Author{
		ID:   data.IDs[index],
		Name: data.Names[index],
	}
	return &author
}
