package motorcycle

import (
	"bytes"
	"encoding/json"
	"gogetters/internal/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*MockRepository)
		expectedStatus int
	}{
		{
			name: "successful creation",
			requestBody: models.Motorcycle{
				Brand:      "Yamaha",
				Totalspeed: 80,
				Fueltype:   "Gasoline",
				Price:      15000.00,
			},
			mockSetup: func(mr *MockRepository) {
				mr.On("CreateMotorcycle", mock.AnythingOfType("*models.Motorcycle")).Return(nil).Run(func(args mock.Arguments) {
					m := args.Get(0).(*models.Motorcycle)
					m.ID = 1
				})
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "invalid JSON",
			requestBody:    "invalid json",
			mockSetup:      func(mr *MockRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "service error",
			requestBody: models.Motorcycle{
				Brand:      "Yamaha",
				Totalspeed: 80,
				Fueltype:   "Gasoline",
				Price:      15000.00,
			},
			mockSetup: func(mr *MockRepository) {
				mr.On("CreateMotorcycle", mock.AnythingOfType("*models.Motorcycle")).Return(assert.AnError)
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockSetup(mockRepo)

			service := NewService(mockRepo)
			handler := NewHandler(service)

			router := gin.New()
			router.POST("/motorcycles", handler.Create)

			body, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/motorcycles", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusCreated {
				var response models.Motorcycle
				json.Unmarshal(w.Body.Bytes(), &response)
				assert.NotZero(t, response.ID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_Create_Integration(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockRepository)
	mockRepo.On("CreateMotorcycle", mock.AnythingOfType("*models.Motorcycle")).Return(nil).Run(func(args mock.Arguments) {
		m := args.Get(0).(*models.Motorcycle)
		m.ID = 1
	})

	service := NewService(mockRepo)
	handler := NewHandler(service)

	router := gin.New()
	router.POST("/motorcycles", handler.Create)

	motorcycle := models.Motorcycle{
		Brand:      "Yamaha",
		Totalspeed: 80,
		Fueltype:   "Gasoline",
		Price:      15000.00,
	}

	body, _ := json.Marshal(motorcycle)
	req, _ := http.NewRequest("POST", "/motorcycles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Motorcycle
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, "Yamaha", response.Brand)
	mockRepo.AssertExpectations(t)
}

func TestHandler_List(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockRepository)
	expectedMotorcycles := []models.Motorcycle{
		{ID: 1, Brand: "Yamaha", Totalspeed: 80, Fueltype: "Gasoline", Price: 15000.00},
		{ID: 2, Brand: "Honda", Totalspeed: 90, Fueltype: "Gasoline", Price: 18000.00},
	}
	mockRepo.On("GetAllMotorcycle").Return(expectedMotorcycles, nil)

	service := NewService(mockRepo)
	handler := NewHandler(service)

	router := gin.New()
	router.GET("/motorcycles", handler.List)

	req, _ := http.NewRequest("GET", "/motorcycles", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Motorcycle
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Len(t, response, 2)
	assert.Equal(t, "Yamaha", response[0].Brand)
	mockRepo.AssertExpectations(t)
}

func TestHandler_Update(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		id             string
		requestBody    interface{}
		mockSetup      func(*MockRepository)
		expectedStatus int
	}{
		{
			name: "successful update",
			id:   "1",
			requestBody: models.Motorcycle{
				Brand:      "Yamaha R1",
				Totalspeed: 90,
				Fueltype:   "Gasoline",
				Price:      20000.00,
			},
			mockSetup: func(mr *MockRepository) {
				mr.On("UpdateMotorcycle", uint(1), mock.AnythingOfType("*models.Motorcycle")).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid id",
			id:             "invalid",
			requestBody:    models.Motorcycle{},
			mockSetup:      func(mr *MockRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid JSON body",
			id:             "1",
			requestBody:    "not a valid json object",
			mockSetup:      func(mr *MockRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "service error",
			id:   "1",
			requestBody: models.Motorcycle{
				Brand:      "Yamaha R1",
				Totalspeed: 90,
				Fueltype:   "Gasoline",
				Price:      20000.00,
			},
			mockSetup: func(mr *MockRepository) {
				mr.On("UpdateMotorcycle", uint(1), mock.AnythingOfType("*models.Motorcycle")).Return(assert.AnError)
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockSetup(mockRepo)

			service := NewService(mockRepo)
			handler := NewHandler(service)

			router := gin.New()
			router.PUT("/motorcycles/:id", handler.Update)

			body, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("PUT", "/motorcycles/"+tt.id, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_Delete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		id             string
		mockSetup      func(*MockRepository)
		expectedStatus int
	}{
		{
			name: "successful delete",
			id:   "1",
			mockSetup: func(mr *MockRepository) {
				mr.On("DeleteMotorcycle", uint(1)).Return(nil)
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "invalid id",
			id:             "abc",
			mockSetup:      func(mr *MockRepository) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "service error",
			id:   "1",
			mockSetup: func(mr *MockRepository) {
				mr.On("DeleteMotorcycle", uint(1)).Return(assert.AnError)
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockSetup(mockRepo)

			service := NewService(mockRepo)
			handler := NewHandler(service)

			router := gin.New()
			router.DELETE("/motorcycles/:id", handler.Delete)

			req, _ := http.NewRequest("DELETE", "/motorcycles/"+tt.id, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			mockRepo.AssertExpectations(t)
		})
	}
}

type benchRepository struct{}

func (s *benchRepository) CreateMotorcycle(*models.Motorcycle) error { return nil }
func (s *benchRepository) GetAllMotorcycle() ([]models.Motorcycle, error) {
	return nil, nil
}
func (s *benchRepository) UpdateMotorcycle(uint, *models.Motorcycle) error { return nil }
func (s *benchRepository) DeleteMotorcycle(uint) error                     { return nil }

func BenchmarkHandlerUpdate(b *testing.B) {
	gin.SetMode(gin.TestMode)

	h := &Handler{
		service: NewService(&benchRepository{}),
	}

	for n := 0; n < b.N; n++ {
		body := bytes.NewBuffer([]byte(`{"brand":"Honda","model":"CBR"}`))
		req := httptest.NewRequest("PUT", "/motorcycles/1", body)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Params = gin.Params{
			{Key: "id", Value: strconv.Itoa(1)},
		}

		h.Update(c)
	}
}
