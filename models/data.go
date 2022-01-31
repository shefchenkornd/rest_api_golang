package models

var db []Book

type Book struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Author        Autor  `json:"author"`
	YearPublished int    `json:"year_published"`
}

type Autor struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	BornYear int    `json:"born_year"`
}

func init() {
	book1 := Book{
		Id:    1,
		Title: "Практика программирования",
		Author: Autor{
			Name:     "Rob",
			LastName: "Pike",
			BornYear: 1956,
		},
		YearPublished: 2001,
	}

	db = append(db, book1)
}

func GetAllBooks() []Book {
	return db
}

func CreateBook(book Book) {
	db = append(db, book)
}

func UpdateBook(newBook Book, oldBook Book) (Book, error) {
	newBook.Id = oldBook.Id
	index := newBook.Id - 1
	db[index] = newBook

	return newBook, nil
}

func DeleteBook(book Book) {
	index := book.Id - 1
	db = append(db[:index], db[index+1:]...)
}

func FindBookById(id int) (Book, bool) {
	var book Book
	var found bool

	for _, b := range db {
		if b.Id == id {
			book = b
			found = true
			break
		}
	}

	return book, found
}
