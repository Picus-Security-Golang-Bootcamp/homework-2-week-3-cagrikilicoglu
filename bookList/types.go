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
