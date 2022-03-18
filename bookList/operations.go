package bookList

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// list : prints the names of all the books in the books slice (the book list)
func List() {
	if strings.Join(os.Args[2:], " ") == "" {
		for _, v := range books {
			fmt.Println(v.Name)
		}
	} else {
		flag.Usage()
	}
}

// searchByName: searches books in the book list by the name input. The search is elastic. The function is case sensitive.
func SearchByName() {
	bookName := strings.Join(os.Args[2:], " ")
	isListed := false

	for _, v := range books {
		if bookName != "" && strings.Contains(strings.ToLower(v.Name), strings.ToLower(bookName)) {
			fmt.Println(v.Name)
			isListed = true
		}
	}

	if !isListed {
		err := errors.New("This book is not the library. Please try another book.")
		fmt.Println(err)
	}
}

// GetByID: searches books by the ID input. The function is case sensitive.
func GetByID() {
	isListed := false
	bookID := strings.Join(os.Args[2:], " ")
	for _, v := range books {
		if bookID == v.ID {
			if !v.IsDeleted {
				fmt.Println(v.Name)
				isListed = true
				break
			} else {
				err := errors.New("The book was on the library. But it is deleted already")
				fmt.Println(err)
				isListed = true
				break
			}
		}
	}
	if !isListed {
		err := errors.New("There is no book with this ID. Please try another book.")
		fmt.Println(err)
	}
}

// delete: delete books by the ID input. The function is case sensitive.
func DeleteByID() {
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
	if !isListed {
		err := errors.New("There is no book with this ID. Please try another book.")
		fmt.Println(err)
	}
}

// buy: buy books by the ID input and of requested quantity inputs given by the user. It checks stock status of the book and decrease if succesfully purchased. Note that the function is case sensitive.
func BuyByID() {
	isListed := false
	bookID := os.Args[2]
	booksOrdered := strings.Join(os.Args[3:], " ")
	for _, v := range books {
		if bookID == v.ID {
			numberOfBooks, err := strconv.Atoi(booksOrdered)
			if err != nil {

				flag.Usage()
			}
			if numberOfBooks != 0 {
				if v.StockNumber >= numberOfBooks {
					if numberOfBooks == 1 {

						fmt.Printf("%d %s (of ID: %s) book is succesfully ordered.\n", numberOfBooks, v.Name, v.ID)
					} else {

						fmt.Printf("%d %s (of ID: %s) books are succesfully ordered.\n", numberOfBooks, v.Name, v.ID)

					}
					UpdateStock(numberOfBooks, v)
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
			}
		}
	}
	if !isListed {
		err := errors.New("There is no book with this ID. Please try another book.")
		fmt.Println(err)
	}
}

// UpdateStock: update stock number of a book by the requsted quantity
func UpdateStock(quantity int, book Book) {
	book.StockNumber -= quantity
}
