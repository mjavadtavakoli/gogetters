package summary_test

import (
	"context"
	"testing"

	"gogetters/internal/models"
	"gogetters/internal/summary"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite db: %v", err)
	}

	if err := db.AutoMigrate(&models.Book{}, &models.Coffee{}, &models.Motorcycle{}, &models.Laptop{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	return db
}

func TestCollectSnapshotCountsRecordsInParallel(t *testing.T) {
	db := setupTestDB(t)

	db.Create(&models.Book{Title: "book", Author: "auth", Year: 2020, Lan: "en"})
	db.Create(&models.Coffee{Late: "latte", Amount: 1})
	db.Create(&models.Motorcycle{Brand: "Yamaha", Fueltype: "gas", Price: 12000, Totalspeed: 80})
	db.Create(&models.Laptop{Cpu: "i7"})

	service := summary.NewService(db)

	snap, err := service.CollectSnapshot(context.Background())
	if err != nil {
		t.Fatalf("CollectSnapshot returned error: %v", err)
	}

	if snap.BookCount != 1 || snap.CoffeeCount != 1 || snap.MotorcycleCount != 1 || snap.LaptopCount != 1 {
		t.Fatalf("unexpected snapshot: %+v", snap)
	}
}

func TestCollectSnapshotRespectsContextCancellation(t *testing.T) {
	db := setupTestDB(t)
	service := summary.NewService(db)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := service.CollectSnapshot(ctx)
	if err == nil {
		t.Fatal("expected error when context cancelled")
	}
}
