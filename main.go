package main

import (
	"errors"
	"flag"
	"fmt"
	"libraryApp/bookList"
	"os"
	"strconv"
	"strings"
)

// type Book struct {
// 	ID          string
// 	Name        string
// 	PageNumber  uint
// 	StockNumber int
// 	Price       float32
// 	StockID     string
// 	ISBN        string
// 	isDeleted   bool
// 	Author      Author
// }

// type Author struct {
// 	ID   string
// 	Name string
// }

// type BookData struct {
// 	IDs          []string
// 	Names        []string
// 	PageNumbers  []uint
// 	StockNumbers []int
// 	Prices       []float32
// 	StockIDs     []string
// 	ISBNs        []string
// 	Authors      []Author
// }
// type AuthorData struct {
// 	IDs   []string
// 	Names []string
// }

// var authors []Author
// var books []Book

var usage = `./homework-2-week-3-cagrikilicoglu

Options: 
	list:    Lists all the books in the library, after list command no additional information submitted.
	search:  Searches the library to check if the search input corresponds any books in the library. The search is elastic, if any book name contains search input words, prints the name of the book.
	get:     Searches the library by the input book ID given by the user after the command.
	delete:  Searches the library by the input book ID given by the user after the command. If the ID matches any of the books, set it's library status to deleted.
	buy:     Searches the library by the input book ID given by the user after the command. If the ID matches any of the books, order the books by the number after input ID.
`
var books []bookList.Book

func main() {
	// dataLength := 5

	// authorData := AuthorData{
	// 	IDs:   []string{"101", "202", "303", "404", "505"},
	// 	Names: []string{"Charles Dickens", "J. R. R. Tolkien", "J. K. Rowling", "Antoine de Saint-Exupéry", "Cao Xueqin"},
	// }

	// for i := 0; i < dataLength; i++ {
	// 	a := *generateAuthor(authorData, i)
	// 	authors = append(authors, a)
	// }

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

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", usage))
	}
	flag.Parse()

	books = bookList.CreateList()
	getCommand()

}

func getCommand() {
	if len(os.Args) > 1 {
		//if an arguement is given, execute respective function
		switch os.Args[1] {
		case "list":
			list()
		case "search":
			searchByName()
		case "get":
			getByID()
		case "delete":
			deleteByID()
		case "buy":
			buyByID()
		default:
			flag.Usage()
		}
	} else {

		flag.Usage()
	}
}

// list : prints the names of all the books in the books slice (the book list)
func list() {
	if strings.Join(os.Args[2:], " ") == "" {
		for _, v := range books {
			fmt.Println(v.Name)
		}
	} else {
		flag.Usage()
	}
}

// searchByName: searches books by the name input given by the user and prints the name of the book if it is included in the book list. The search is elastic so that if the name input word is contained in the name of any book prints its name also. Moreover, the function is case insensitive.

func searchByName() {
	bookName := strings.Join(os.Args[2:], " ")
	isListed := false
	for _, v := range books {
		if strings.ToLower(bookName) == strings.ToLower(v.Name) {
			fmt.Println(v.Name)
			isListed = true
			break
		}
	}

	if isListed == false {
		for _, v := range books {
			if strings.Contains(strings.ToLower(v.Name), strings.ToLower(bookName)) {
				fmt.Println(v.Name)
				isListed = true
			}
		}
	}

	if isListed == false {
		// usage = "This book is not the library. Please try another book."
		// flag.Usage()
		err := errors.New("This book is not the library. Please try another book.")
		fmt.Println(err)
	}

}

// get: searches books by the ID input given by the user and prints the name of the book if the ID is included in the book list. Note that the function is case sensitive.

func getByID() {
	isListed := false
	bookID := strings.Join(os.Args[2:], " ")
	for _, v := range books {
		if bookID == v.ID {

			if !v.IsDeleted {
				fmt.Println(v.Name)
				isListed = true
				break
			} else {
				err := errors.New("The book was on the library. But it is deleted already.")
				fmt.Println(err)
				isListed = true
				break
				// usage = "The book was on the library. But it is deleted already."
				// flag.Usage()
			}

		}

	}

	if isListed == false {
		// usage = "There is no book with this ID. Please try another book."
		// flag.Usage()
		err := errors.New("There is no book with this ID. Please try another book.")
		fmt.Println(err)
	}

}

// delete: delete books by the ID input given by the user and prints a successful deletion message if the ID is included in the book list. Note that the function is case sensitive.
func deleteByID() {
	isListed := false
	bookID := strings.Join(os.Args[2:], " ")
	for i, v := range books {
		if bookID == v.ID {
			if !v.IsDeleted {
				books[i].IsDeleted = true
				fmt.Printf("%s (with the ID: %s) is successfully deleted from the book list.\n", v.Name, v.ID)
				isListed = true
				break
			} else {
				err := errors.New("The book was on the library. But it is deleted already.")
				fmt.Println(err)
				isListed = true
				break
			}
		}

	}
	if isListed == false {
		err := errors.New("There is no book with this ID. Please try another book.")
		fmt.Println(err)
		// usage = "There is no book with this ID. Please try another book."
		// flag.Usage()
	}

}

// buy: buy books by the ID input and of requested quantity inputs given by the user. It prints a successful deletion message if the ID is included in the book list and the book has enough stock. Note that the function is case sensitive.
func buyByID() {
	isListed := false
	bookID := os.Args[2]
	booksOrdered := strings.Join(os.Args[3:], " ")
	for _, v := range books {
		if bookID == v.ID {

			numberOfBooks, err := strconv.Atoi(booksOrdered)
			if err != nil {
				// handle error
				// usage = "Please type only the number of books that you want to order."
				flag.Usage()
				// err := errors.New("Please type only the number of books that you want to order.")
				// fmt.Println(err)
			}
			if numberOfBooks != 0 {
				if v.StockNumber >= numberOfBooks {
					if numberOfBooks == 1 {
						fmt.Printf("%d %s (of ID: %s) book is succesfully ordered.\n", numberOfBooks, v.Name, v.ID)
					} else {
						fmt.Printf("%d %s (of ID: %s) books are succesfully ordered.\n", numberOfBooks, v.Name, v.ID)
					}

				} else {
					if v.StockNumber == 0 {
						fmt.Printf("The book is out of stock. Please try ordering later.\n")
					} else {
						fmt.Printf("No sufficient %s (of ID: %s) books in the stock. Please try ordering less than %d books.\n", v.Name, v.ID, (v.StockNumber + 1))
					}

				}
				isListed = true
				break
			} else {
				err := errors.New("Please type the quantity that you want to order.")
				fmt.Println(err)
				isListed = true
				break
				// fmt.Printf("Please type the quantity that you want to order.\n")
			}

		}
	}

	if isListed == false {
		err := errors.New("There is no book with this ID. Please try another book.")
		fmt.Println(err)

		// flag.Usage()
	}

}

// // generateBook : takes a book data struct as an input and create books with specified attributes given in this struct.
// func generateBook(data BookData, index int) *Book {

// 	book := Book{
// 		ID:          data.IDs[index],
// 		Name:        data.Names[index],
// 		PageNumber:  data.PageNumbers[index],
// 		StockNumber: data.StockNumbers[index],
// 		Price:       data.Prices[index],
// 		StockID:     data.StockIDs[index],
// 		ISBN:        data.ISBNs[index],
// 		Author:      data.Authors[index],
// 	}
// 	return &book
// }

// // generateBook : takes an author data struct as an input and create authors with specified attributes given in this struct.
// func generateAuthor(data AuthorData, index int) *Author {

// 	author := Author{
// 		ID:   data.IDs[index],
// 		Name: data.Names[index],
// 	}
// 	return &author
// }
