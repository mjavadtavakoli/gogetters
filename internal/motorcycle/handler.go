package motorcycle

import (
	"gogetters/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Create(c *gin.Context) {
	var motorcycle models.Motorcycle

	if err := c.ShouldBindJSON(&motorcycle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateMotorcycle(&motorcycle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, motorcycle)

}
func (h *Handler) List(c *gin.Context) {
	motorcycles, _ := h.service.GetAllMotorcycle()

	c.JSON(http.StatusOK, motorcycles)	
}

func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")
	var motorcycle models.Motorcycle
	
	if err := c.ShouldBindJSON(&motorcycle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.UpdateMotorcycle(uint(idUint), &motorcycle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, motorcycle)
}