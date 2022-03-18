# Assignment | Week 3

> Library simulating app

The app contains a list of top-selling books of all time. It simulates a library that a user can list all of the books in the library, search a book by its name or ID, delete a book by its ID, and order books by their ID.

## Features

Users can:

- Print all the books on the list.
- Search if a book is on the list or not by its name.
- Search if a book is on the list or not by its ID.
- Delete a book from the list by its ID.
- Buy books from the list by their ID and of preferred quantity.

## Configuration

The app can execute five different commands given by the user on the terminal.

### list Arguement

With the "list" arguement, the user can print all the books on the top-selling books list. Example usage can be seen below.

```
go run main.go list
```

### search Arguement

With the "search" arguement, the user can search if a certain book of his/her choice is on the list, by the name of the book. After search arguement, the name of the book to be searched should be written. If this book is on the list, the app prints its name. If it is not, the app gives error.

"search" arguement, searches the books elastically meaning that if any of the books on the list contains the <bookName> input, the app still finds and prints its name.

Note that "search" arguement is case insensitive.

Example usage can be seen below.

```
go run main.go search <bookName>
go run main.go search The Hobbit
```

### get Arguement

With the "get" arguement, the user can search if a certain book of his/her choice is on the list by the ID of the book. After "get" arguement, the ID of the book to be searched should be written. If any book corresponding the input ID is on the list, the app prints its name. If it is not, the app gives error.

Note that the "get" arguement is case sensitive to prevent misleading to another book.

Example usage can be seen below.

```
go run main.go search <bookID>
go run main.go search 3
```

### delete Arguement

With the "delete" arguement, the user can delete a certain book of his/her choice from list by the ID of the book. After "delete" arguement, the ID of the book to be deleted should be written. If any book corresponding the input ID is on the list, the app set the book status to deleted and prints a successful deletion message. If it is not, the app gives error.

Note that the "delete" arguement is case sensitive in terms of ID to prevent misleading to another book.

Also note that, the "delete" command is not literally deleting a book from the list but set its status to deleted. Therefore the book will stay on the list.

Example usage can be seen below.

```
go run main.go delete <bookID>
go run main.go delete 3
```

### buy Arguement

With the "buy" arguement, the user can buy a certain book of his/her choice from list by the ID of the book. After "buy" arguement, the ID of the book to be ordered should be written. User can also specify the number of the books he/she wants to order after the ID of the book. If any book corresponding the input ID is on the list, the app prints a successful ordering message and decreases its stock number. If it is not, the app gives error.

Note that if the stock number of the book is below the number of the books to be ordered, the app prints an unsuccesful ordering message.

Note that the "buy" arguement is case sensitive in terms of ID to prevent misleading to another book.

Example usage can be seen below.

```
go run main.go buy <bookID> <quantity>
go run main.go buy 3 2
```

## Links

- Project repository: https://github.com/Picus-Security-Golang-Bootcamp/homework-2-week-3-cagrikilicoglu

## License

The project has no license.
