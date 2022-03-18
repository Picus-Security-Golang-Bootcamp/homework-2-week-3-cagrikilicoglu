package main

import (
	"flag"
	"fmt"
	"libraryApp/bookList"
	"os"
)

var usage = `./homework-2-week-3-cagrikilicoglu

Options: 
	list:    Lists all the books in the library, after list command no additional information submitted.
	search:  Searches the library to check if the search input corresponds any books in the library. The search is elastic, if any book name contains search input words, prints the name of the book.
	get:     Searches the library by the input book ID given by the user after the command.
	delete:  Searches the library by the input book ID given by the user after the command. If the ID matches any of the books, set it's library status to deleted.
	buy:     Searches the library by the input book ID given by the user after the command. If the ID matches any of the books, order the books by the number after input ID.
`

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n", usage))
	}
	flag.Parse()

	getCommand()
}

// getCommand: get user input from the terminal
func getCommand() {
	if len(os.Args) > 1 {
		//if an arguement is given, execute respective function
		switch os.Args[1] {
		case "list":
			bookList.List()
		case "search":
			bookList.SearchByName()
		case "get":
			bookList.GetByID()
		case "delete":
			bookList.DeleteByID()
		case "buy":
			bookList.BuyByID()
		default:
			flag.Usage()
		}
	} else {
		flag.Usage()
	}
}
