package motorcycle

import (
	"errors"
	"gogetters/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock implementation of the RepositoryInterface
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateMotorcycle(motorcycle *models.Motorcycle) error {
	args := m.Called(motorcycle)
	return args.Error(0)
}

func (m *MockRepository) GetAllMotorcycle() ([]models.Motorcycle, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Motorcycle), args.Error(1)
}

func (m *MockRepository) UpdateMotorcycle(id uint, motorcycle *models.Motorcycle) error {
	args := m.Called(id, motorcycle)
	return args.Error(0)
}

func (m *MockRepository) DeleteMotorcycle(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestService_CreateMotorcycle(t *testing.T) {
	tests := []struct {
		name       string
		motorcycle *models.Motorcycle
		mockSetup  func(*MockRepository)
		wantErr    bool
	}{
		{
			name: "successful creation",
			motorcycle: &models.Motorcycle{
				Brand:      "Yamaha",
				Totalspeed: 80,
				Fueltype:   "Gasoline",
				Price:      15000.00,
			},
			mockSetup: func(mr *MockRepository) {
				mr.On("CreateMotorcycle", mock.AnythingOfType("*models.Motorcycle")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "repository error",
			motorcycle: &models.Motorcycle{
				Brand:      "Yamaha",
				Totalspeed: 80,
				Fueltype:   "Gasoline",
				Price:      15000.00,
			},
			mockSetup: func(mr *MockRepository) {
				mr.On("CreateMotorcycle", mock.AnythingOfType("*models.Motorcycle")).Return(errors.New("database error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockSetup(mockRepo)

			service := NewService(mockRepo)
			err := service.CreateMotorcycle(tt.motorcycle)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_GetAllMotorcycle(t *testing.T) {
	tests := []struct {
		name          string
		mockSetup     func(*MockRepository)
		expectedCount int
		wantErr       bool
	}{
		{
			name: "successful retrieval",
			mockSetup: func(mr *MockRepository) {
				motorcycles := []models.Motorcycle{
					{ID: 1, Brand: "Yamaha", Totalspeed: 80, Fueltype: "Gasoline", Price: 15000.00},
					{ID: 2, Brand: "Honda", Totalspeed: 90, Fueltype: "Gasoline", Price: 18000.00},
				}
				mr.On("GetAllMotorcycle").Return(motorcycles, nil)
			},
			expectedCount: 2,
			wantErr:       false,
		},
		{
			name: "empty list",
			mockSetup: func(mr *MockRepository) {
				mr.On("GetAllMotorcycle").Return([]models.Motorcycle{}, nil)
			},
			expectedCount: 0,
			wantErr:       false,
		},
		{
			name: "repository error",
			mockSetup: func(mr *MockRepository) {
				mr.On("GetAllMotorcycle").Return(nil, errors.New("database error"))
			},
			expectedCount: 0,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockSetup(mockRepo)

			service := NewService(mockRepo)
			motorcycles, err := service.GetAllMotorcycle()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, motorcycles, tt.expectedCount)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_UpdateMotorcycle(t *testing.T) {
	tests := []struct {
		name       string
		id         uint
		motorcycle *models.Motorcycle
		mockSetup  func(*MockRepository)
		wantErr    bool
	}{
		{
			name: "successful update",
			id:   1,
			motorcycle: &models.Motorcycle{
				Brand:      "Yamaha R1",
				Totalspeed: 90,
				Fueltype:   "Gasoline",
				Price:      20000.00,
			},
			mockSetup: func(mr *MockRepository) {
				mr.On("UpdateMotorcycle", uint(1), mock.AnythingOfType("*models.Motorcycle")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "repository error",
			id:   1,
			motorcycle: &models.Motorcycle{
				Brand:      "Yamaha R1",
				Totalspeed: 90,
				Fueltype:   "Gasoline",
				Price:      20000.00,
			},
			mockSetup: func(mr *MockRepository) {
				mr.On("UpdateMotorcycle", uint(1), mock.AnythingOfType("*models.Motorcycle")).Return(errors.New("database error"))
			},
			wantErr: true,
		},
		{
			name: "not found error",
			id:   999,
			motorcycle: &models.Motorcycle{
				Brand:      "Yamaha R1",
				Totalspeed: 90,
				Fueltype:   "Gasoline",
				Price:      20000.00,
			},
			mockSetup: func(mr *MockRepository) {
				mr.On("UpdateMotorcycle", uint(999), mock.AnythingOfType("*models.Motorcycle")).Return(errors.New("record not found"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockSetup(mockRepo)

			service := NewService(mockRepo)
			err := service.UpdateMotorcycle(tt.id, tt.motorcycle)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_DeleteMotorcycle(t *testing.T) {
	tests := []struct {
		name      string
		id        uint
		mockSetup func(*MockRepository)
		wantErr   bool
	}{
		{
			name: "successful delete",
			id:   1,
			mockSetup: func(mr *MockRepository) {
				mr.On("DeleteMotorcycle", uint(1)).Return(nil)
			},
			wantErr: false,
		},
		{
			name:      "invalid id",
			id:        0,
			mockSetup: func(mr *MockRepository) {},
			wantErr:   true,
		},
		{
			name: "repository error",
			id:   2,
			mockSetup: func(mr *MockRepository) {
				mr.On("DeleteMotorcycle", uint(2)).Return(errors.New("database error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockSetup(mockRepo)

			service := NewService(mockRepo)
			err := service.DeleteMotorcycle(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
