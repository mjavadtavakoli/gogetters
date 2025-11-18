package motorcycle

import (
	"gogetters/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Auto migrate the Motorcycle model
	err = db.AutoMigrate(&models.Motorcycle{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

func TestRepository_CreateMotorcycle(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)

	tests := []struct {
		name       string
		motorcycle *models.Motorcycle
		wantErr    bool
	}{
		{
			name: "successful creation",
			motorcycle: &models.Motorcycle{
				Brand:      "Yamaha",
				Totalspeed: 180,
				Fueltype:   "Gasoline",
				Price:      15000.00,
			},
			wantErr: false,
		},
		{
			name: "create another motorcycle",
			motorcycle: &models.Motorcycle{
				Brand:      "Honda",
				Totalspeed: 200,
				Fueltype:   "Gasoline",
				Price:      18000.00,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.CreateMotorcycle(tt.motorcycle)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotZero(t, tt.motorcycle.ID, "ID should be set after creation")
			}
		})
	}
}

func TestRepository_GetAllMotorcycle(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)

	// Create test data
	motorcycles := []models.Motorcycle{
		{Brand: "Yamaha", Totalspeed: 180, Fueltype: "Gasoline", Price: 15000.00},
		{Brand: "Honda", Totalspeed: 200, Fueltype: "Gasoline", Price: 18000.00},
		{Brand: "Kawasaki", Totalspeed: 190, Fueltype: "Gasoline", Price: 17000.00},
	}

	for i := range motorcycles {
		err := repo.CreateMotorcycle(&motorcycles[i])
		assert.NoError(t, err)
	}

	t.Run("get all motorcycles", func(t *testing.T) {
		result, err := repo.GetAllMotorcycle()
		assert.NoError(t, err)
		assert.Len(t, result, 3)
		assert.Equal(t, "Yamaha", result[0].Brand)
		assert.Equal(t, "Honda", result[1].Brand)
		assert.Equal(t, "Kawasaki", result[2].Brand)
	})

	t.Run("get empty list", func(t *testing.T) {
		emptyDB := setupTestDB(t)
		emptyRepo := NewRepository(emptyDB)
		result, err := emptyRepo.GetAllMotorcycle()
		assert.NoError(t, err)
		assert.Len(t, result, 0)
	})
}

func TestRepository_UpdateMotorcycle(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)

	// Create a test motorcycle
	motorcycle := &models.Motorcycle{
		Brand:      "Yamaha",
		Totalspeed: 180,
		Fueltype:   "Gasoline",
		Price:      15000.00,
	}
	err := repo.CreateMotorcycle(motorcycle)
	assert.NoError(t, err)
	originalID := motorcycle.ID

	tests := []struct {
		name       string
		id         uint
		updateData *models.Motorcycle
		wantErr    bool
	}{
		{
			name: "successful update",
			id:   originalID,
			updateData: &models.Motorcycle{
				Brand:      "Yamaha R1",
				Totalspeed: 200,
				Fueltype:   "Gasoline",
				Price:      20000.00,
			},
			wantErr: false,
		},
		{
			name: "update non-existent id",
			id:   999,
			updateData: &models.Motorcycle{
				Brand:      "Yamaha R1",
				Totalspeed: 200,
				Fueltype:   "Gasoline",
				Price:      20000.00,
			},
			wantErr: false, // GORM Updates doesn't return error for non-existent records
		},
		{
			name: "partial update",
			id:   originalID,
			updateData: &models.Motorcycle{
				Brand: "Yamaha R1 Updated",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.UpdateMotorcycle(tt.id, tt.updateData)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Verify the update if it was successful
			if !tt.wantErr && tt.id == originalID {
				var updated models.Motorcycle
				db.First(&updated, originalID)
				if tt.updateData.Brand != "" {
					assert.Equal(t, tt.updateData.Brand, updated.Brand)
				}
			}
		})
	}
}

func TestRepository_Integration(t *testing.T) {
	db := setupTestDB(t)
	repo := NewRepository(db)

	// Create
	motorcycle := &models.Motorcycle{
		Brand:      "Yamaha",
		Totalspeed: 180,
		Fueltype:   "Gasoline",
		Price:      15000.00,
	}
	err := repo.CreateMotorcycle(motorcycle)
	assert.NoError(t, err)
	assert.NotZero(t, motorcycle.ID)

	// Read
	all, err := repo.GetAllMotorcycle()
	assert.NoError(t, err)
	assert.Len(t, all, 1)
	assert.Equal(t, "Yamaha", all[0].Brand)

	// Update
	updateData := &models.Motorcycle{
		Brand:      "Yamaha R1",
		Totalspeed: 200,
		Fueltype:   "Gasoline",
		Price:      20000.00,
	}
	err = repo.UpdateMotorcycle(motorcycle.ID, updateData)
	assert.NoError(t, err)

	// Verify update
	all, err = repo.GetAllMotorcycle()
	assert.NoError(t, err)
	assert.Len(t, all, 1)
	assert.Equal(t, "Yamaha R1", all[0].Brand)
	assert.Equal(t, 200, all[0].Totalspeed)
	assert.Equal(t, 20000.00, all[0].Price)
}

