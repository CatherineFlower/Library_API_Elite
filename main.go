package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Определение структур данных
type AdressLibrary struct {
	ID      int    `json:"id"`
	Country string `json:"страна"`
	City    string `json:"город"`
	Street  string `json:"улица"`
	House   string `json:"дом"`
}

type Library struct {
	ID       int    `json:"id"`
	Name     string `json:"название"`
	AdressID *int   `json:"id_адрес,omitempty"`
}

type Publisher struct {
	ID   int    `json:"id"`
	Name string `json:"название"`
	City string `json:"город"`
}

type Author struct {
	ID         int     `json:"id"`
	LastName   string  `json:"фамилия"`
	FirstName  string  `json:"имя"`
	Patronymic *string `json:"отчество,omitempty"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"название"`
}

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"название"`
	Year  int    `json:"год_издания"`
}

type Customer struct {
	ID         int     `json:"id"`
	LastName   string  `json:"фамилия"`
	FirstName  string  `json:"имя"`
	Patronymic *string `json:"отчество,omitempty"`
	Phone      string  `json:"телефон"`
}

type AddressDelivery struct {
	ID         int    `json:"id"`
	CustomerID int    `json:"id_заказчик"`
	Address    string `json:"адрес"`
	Active     bool   `json:"активность"`
}

type Order struct {
	ID             int    `json:"id"`
	CustomerID     int    `json:"id_заказчик"`
	DeliveryAddrID *int   `json:"id_адрес_доставки,omitempty"`
	IssueDate      string `json:"дата_выдачи"`
	Status         string `json:"статус"`
}

type BookLibrary struct {
	BookID    int     `json:"id_книга"`
	LibraryID int     `json:"id_библиотека"`
	Quantity  int     `json:"количество_экземпляров"`
	Price     float64 `json:"цена"`
	Condition string  `json:"состояние"`
}

type PublisherBook struct {
	BookID      int `json:"id_книга"`
	PublisherID int `json:"id_издательство"`
}

type AuthorBook struct {
	BookID   int `json:"id_книга"`
	AuthorID int `json:"id_автор"`
}

type GenreBook struct {
	BookID  int `json:"id_книга"`
	GenreID int `json:"id_жанр"`
}

type PublisherLibrary struct {
	PublisherID int `json:"id_издательство"`
	LibraryID   int `json:"id_библиотека"`
}

type AuthorLibrary struct {
	AuthorID  int `json:"id_автор"`
	LibraryID int `json:"id_библиотека"`
}

type OrderLibrary struct {
	OrderID   int `json:"id_заказ"`
	LibraryID int `json:"id_библиотека"`
}

type OrderBook struct {
	OrderID   int     `json:"id_заказ"`
	BookID    int     `json:"id_книга"`
	Type      string  `json:"тип"`
	ReturnDue *string `json:"срок_возврата,omitempty"`
}

// Глобальные массивы данных
var (
	adressLibraries = []AdressLibrary{
		{ID: 1, Country: "Россия", City: "Москва", Street: "Ленина", House: "10"},
		{ID: 2, Country: "США", City: "Нью-Йорк", Street: "Бродвей", House: "15"},
	}

	libraries = []Library{
		{ID: 1, Name: "Центральная библиотека", AdressID: &adressLibraries[0].ID},
		{ID: 2, Name: "Городская библиотека", AdressID: &adressLibraries[1].ID},
	}

	publishers = []Publisher{
		{ID: 1, Name: "Издательство АСТ", City: "Москва"},
		{ID: 2, Name: "Penguin Books", City: "Лондон"},
	}

	authors = []Author{
		{ID: 1, LastName: "Толстой", FirstName: "Лев", Patronymic: nil},
		{ID: 2, LastName: "Оруэлл", FirstName: "Джордж", Patronymic: nil},
	}

	genres = []Genre{
		{ID: 1, Name: "Фантастика"},
		{ID: 2, Name: "Классика"},
	}

	books = []Book{
		{ID: 1, Title: "1984", Year: 1949},
		{ID: 2, Title: "Война и мир", Year: 1869},
	}

	customers = []Customer{
		{ID: 1, LastName: "Иванов", FirstName: "Алексей", Patronymic: nil, Phone: "+79111234567"},
		{ID: 2, LastName: "Смит", FirstName: "Джон", Patronymic: nil, Phone: "+12025550123"},
	}

	addressDeliveries = []AddressDelivery{
		{ID: 1, CustomerID: customers[0].ID, Address: "ул. Ленина, д. 10", Active: true},
		{ID: 2, CustomerID: customers[1].ID, Address: "5-я авеню, д. 12", Active: true},
	}

	orders = []Order{
		{ID: 1, CustomerID: customers[0].ID, DeliveryAddrID: &addressDeliveries[0].ID, IssueDate: "2024-02-20", Status: "в обработке"},
		{ID: 2, CustomerID: customers[1].ID, DeliveryAddrID: &addressDeliveries[1].ID, IssueDate: "2024-02-18", Status: "доставлен"},
	}

	booksLibraries = []BookLibrary{
		{BookID: books[0].ID, LibraryID: libraries[0].ID, Quantity: 5, Price: 500.00, Condition: "новая"},
		{BookID: books[1].ID, LibraryID: libraries[1].ID, Quantity: 3, Price: 1200.00, Condition: "б/у"},
	}

	publishersBooks = []PublisherBook{
		{BookID: books[0].ID, PublisherID: publishers[1].ID},
		{BookID: books[1].ID, PublisherID: publishers[0].ID},
	}

	authorsBooks = []AuthorBook{
		{BookID: books[0].ID, AuthorID: authors[1].ID},
		{BookID: books[1].ID, AuthorID: authors[0].ID},
	}

	genresBooks = []GenreBook{
		{BookID: books[0].ID, GenreID: genres[0].ID},
		{BookID: books[1].ID, GenreID: genres[1].ID},
	}

	publishersLibraries = []PublisherLibrary{
		{PublisherID: publishers[0].ID, LibraryID: libraries[0].ID},
		{PublisherID: publishers[1].ID, LibraryID: libraries[1].ID},
	}

	authorsLibraries = []AuthorLibrary{
		{AuthorID: authors[0].ID, LibraryID: libraries[0].ID},
		{AuthorID: authors[1].ID, LibraryID: libraries[1].ID},
	}

	ordersLibraries = []OrderLibrary{
		{OrderID: orders[0].ID, LibraryID: libraries[0].ID},
		{OrderID: orders[1].ID, LibraryID: libraries[1].ID},
	}

	ordersBooks = []OrderBook{
		{OrderID: orders[0].ID, BookID: books[0].ID, Type: "аренда", ReturnDue: nil},
		{OrderID: orders[1].ID, BookID: books[1].ID, Type: "покупка", ReturnDue: nil},
	}
)

func main() {
	r := gin.Default()

	// Маршруты для всех 17 таблиц
	addCRUDRoutes(r, "/address_libraries", &adressLibraries)
	addCRUDRoutes(r, "/libraries", &libraries)
	addCRUDRoutes(r, "/publishers", &publishers)
	addCRUDRoutes(r, "/authors", &authors)
	addCRUDRoutes(r, "/genres", &genres)
	addCRUDRoutes(r, "/books", &books)
	addCRUDRoutes(r, "/customers", &customers)
	addCRUDRoutes(r, "/address_deliveries", &addressDeliveries)
	addCRUDRoutes(r, "/orders", &orders)
	addCRUDRoutes(r, "/book_libraries", &booksLibraries)
	addCRUDRoutes(r, "/publisher_books", &publishersBooks)
	addCRUDRoutes(r, "/author_books", &authorsBooks)
	addCRUDRoutes(r, "/genre_books", &genresBooks)
	addCRUDRoutes(r, "/publisher_libraries", &publishersLibraries)
	addCRUDRoutes(r, "/author_libraries", &authorsLibraries)
	addCRUDRoutes(r, "/order_libraries", &ordersLibraries)
	addCRUDRoutes(r, "/order_books", &ordersBooks)

	r.Run(":8080")
}

// Универсальная функция для создания CRUD-маршрутов
func addCRUDRoutes[T any](r *gin.Engine, path string, data *[]T) {
	r.GET(path, func(c *gin.Context) {
		c.JSON(http.StatusOK, *data)
	})

	r.POST(path, func(c *gin.Context) {
		var newItem T
		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		*data = append(*data, newItem)
		c.JSON(http.StatusCreated, newItem)
	})

	r.GET(path+"/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
			return
		}
		for _, item := range *data {
			if id == getItemID(item) {
				c.JSON(http.StatusOK, item)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	})

	r.PUT(path+"/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
			return
		}
		var updatedItem T
		if err := c.ShouldBindJSON(&updatedItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		for i, item := range *data {
			if id == getItemID(item) {
				(*data)[i] = updatedItem
				c.JSON(http.StatusOK, updatedItem)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	})

	r.DELETE(path+"/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
			return
		}
		for i, item := range *data {
			if id == getItemID(item) {
				*data = append((*data)[:i], (*data)[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	})
}

// Получение ID структуры (универсально)
func getItemID[T any](item T) int {
	switch v := any(item).(type) {
	case AdressLibrary:
		return v.ID
	case Library:
		return v.ID
	case Publisher:
		return v.ID
	case Author:
		return v.ID
	case Genre:
		return v.ID
	case Book:
		return v.ID
	case Customer:
		return v.ID
	case AddressDelivery:
		return v.ID
	case Order:
		return v.ID
	}
	return 0
}
