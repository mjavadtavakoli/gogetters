package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"gogetters/internal/book"
	"gogetters/internal/coffee"
	"gogetters/internal/database"
	"gogetters/internal/laptop"
	"gogetters/internal/models"
	"gogetters/internal/motorcycle"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"time"
	//"testing"
)

func main() {
	// CPU profiling flag
	var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()

	// Start CPU profiling if flag is set
	var profileFile *os.File
	if *cpuProfile != "" {
		var err error
		profileFile, err = os.Create(*cpuProfile)
		if err != nil {
			log.Fatalf("Failed to create CPU profile: %v", err)
		}
		if err := pprof.StartCPUProfile(profileFile); err != nil {
			log.Fatalf("Failed to start CPU profile: %v", err)
		}
		log.Printf("CPU profiling enabled, writing to %s", *cpuProfile)
	}

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
	laptopRepo := laptop.NewRepository(db)
	laptopService := laptop.NewService(laptopRepo)
	laptopHandler := laptop.NewHandler(laptopService)

	// ---- ROUTER ----
	r := gin.Default()

	// Book Routes
	r.POST("/books", bookHandler.Create)
	r.GET("/books", bookHandler.List)
	r.PUT("/books/:id", bookHandler.Update)

	// Coffee Routes
	r.POST("/coffees", coffeeHandler.Create)
	r.GET("/coffees", coffeeHandler.List)
	r.PUT("/coffees/:id", coffeeHandler.Update)

	// Motorcycle Routes
	r.POST("/motorcycles", motorcycleHandler.Create)
	r.GET("/motorcycles", motorcycleHandler.List)
	r.PUT("/motorcycles/:id", motorcycleHandler.Update)
	r.DELETE("/motorcycles/:id", motorcycleHandler.Delete)

	// Laptop Routes
	r.POST("/laptops", laptopHandler.Create)
	r.GET("/laptops", laptopHandler.List)
	r.PUT("/laptops/:id", laptopHandler.Update)

	// Create HTTP server with graceful shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Graceful shutdown handling
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Stop CPU profiling before shutdown
	if profileFile != nil {
		pprof.StopCPUProfile()
		profileFile.Close()
		log.Printf("CPU profile saved to %s", *cpuProfile)
	}

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
