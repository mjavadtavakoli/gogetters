package summary

import (
	"context"
	"sync"

	"gogetters/internal/models"

	"gorm.io/gorm"
)

// Service concurrently gathers aggregate stats across domain models.
type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

type Snapshot struct {
	BookCount       int64 `json:"books"`
	CoffeeCount     int64 `json:"coffees"`
	MotorcycleCount int64 `json:"motorcycles"`
	//LaptopCount     int64 `json:"laptops"`
}

// CollectSnapshot counts each resource in parallel so that the request latency
// follows the slowest single query instead of the sum of all queries.
func (s *Service) CollectSnapshot(ctx context.Context) (Snapshot, error) {
	type task struct {
		model any
		set   func(*Snapshot, int64)
	}

	tasks := []task{
		{model: &models.Book{}, set: func(snap *Snapshot, count int64) { snap.BookCount = count }},
		{model: &models.Coffee{}, set: func(snap *Snapshot, count int64) { snap.CoffeeCount = count }},
		{model: &models.Motorcycle{}, set: func(snap *Snapshot, count int64) { snap.MotorcycleCount = count }},
		//{model: &models.Laptop{}, set: func(snap *Snapshot, count int64) { snap.LaptopCount = count }},
	}

	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		snap    Snapshot
		errOnce sync.Once
		retErr  error
	)

	ctxDB := s.db.WithContext(ctx)

	for _, t := range tasks {
		wg.Add(1)
		go func(t task) {
			defer wg.Done()

			var count int64
			if err := ctxDB.Model(t.model).Count(&count).Error; err != nil {
				errOnce.Do(func() {
					retErr = err
				})
				return
			}

			mu.Lock()
			t.set(&snap, count)
			mu.Unlock()
		}(t)
	}

	wg.Wait()

	return snap, retErr
}
