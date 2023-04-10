package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"bookstore/book"
	"sort"
	"os"
)


func GetBooks() []book.Book{
	var boooks []book.Book
	//dsn := "host=localhost user=postgres password=dadon2004 dbname=assignmentGo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("db are connected")
	db.AutoMigrate(&book.Book{})
	db.Find(&boooks)

	return boooks
}

func UpdateById(buk book.Book){
	var books book.Book
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("db are connected")
	db.First(&books, buk.ID)
	books.Title = buk.Title
	books.Description = buk.Description
	db.Save(&books)
}

func DeleteById(Id string){
	var books book.Book
	//dsn := "host=localhost user=postgres password=dadon2004 dbname=assignmentGo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("db are connected")
	db.Delete(&books, Id)
}

func GetById(Id string) book.Book{
	var books book.Book
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("db are connected")
	db.First(&books, Id)
	return books
}

func Insert(samp book.Book){
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("db are connected")
	insertProduct := &book.Book{ID: samp.ID, Title: samp.Title, Description: samp.Description, Price: samp.Price}
  	db.Create(insertProduct)
}

func SearchByName(Name string) book.Book{
	var books book.Book
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("db are connected")
	db.Where(fmt.Sprintf("title LIKE '%s'", Name)).Find(&books)
	return books
}

func FilterByPriceAsc() []book.Book{
	items := GetBooks()
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Price < items[j].Price
	})
	return items
}

func FilterByPriceDesc() []book.Book{
	items := GetBooks()
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Price > items[j].Price
	})
	return items
}