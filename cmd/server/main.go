package main

import (
    "gogetters/internal/laptop"
    "gogetters/internal/motorcycle"
    "gogetters/internal/book"
    "gogetters/internal/coffee"
    "gogetters/internal/database"
    "gogetters/internal/models"
    "github.com/gin-gonic/gin"
)

func main() {
    dsn := "host=localhost user=postgres password=7878 dbname=gogetters port=5432 sslmode=disable"
    db := database.Connect(dsn)

    // Migrate Models
    db.AutoMigrate(&models.Book{}, &models.Coffee{}, &models.Motorcycle{}, &models.Laptop{})

    // ---- BOOK SETUP ----
    bookRepo := book.NewRepository(db)
    bookService := book.NewService(bookRepo)
    bookHandler := book.NewHandler(bookService)

    // ---- COFFEE SETUP ----
    coffeeRepo := coffee.NewRepository(db)
    coffeeService := coffee.NewService(coffeeRepo)
    coffeeHandler := coffee.NewHandler(coffeeService)

    // ---- MOTORCYCLE SETUP ----
    motorcycleRepo := motorcycle.NewRepository(db)
    motorcycleService := motorcycle.NewService(motorcycleRepo)
    motorcycleHandler := motorcycle.NewHandler(motorcycleService)

    // ---- LAPTOP SETUP ----
    laptopRepo :=  laptop.NewRepository(db)
    laptopService :=  laptop.NewService( laptopRepo)
    laptopHandler :=  laptop.NewHandler( laptopService)

    // ---- ROUTER ----
    r := gin.Default()

    // Book Routes
    r.POST("/books", bookHandler.Create)
    r.GET("/books", bookHandler.List)

    // Coffee Routes
    r.POST("/coffees", coffeeHandler.Create)
    r.GET("/coffees", coffeeHandler.List)

    // Motorcycle Routes
    r.POST("/motorcycles", motorcycleHandler.Create)
    r.GET("/motorcycles", motorcycleHandler.List)

    // Laptop Routes
     r.POST("/laptops", laptopHandler.Create)
    r.GET("/laptops", laptopHandler.List)


    r.Run(":8080")
} 
