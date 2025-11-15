package coffee

import (
	"gogetters/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Create(c *gin.Context) {
    var coffee models.Coffee
    if err := c.ShouldBindJSON(&coffee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.CreateCoffee(&coffee); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, coffee)
}

func (h *Handler) List(c *gin.Context) {
    coffees, err := h.service.GetAllCoffee()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, coffees)
}
