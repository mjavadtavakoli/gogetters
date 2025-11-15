package book

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
    var book models.Book

    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.CreateBook(&book); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, book)
}

func (h *Handler) List(c *gin.Context) {
    books, _ := h.service.ListBooks()
    c.JSON(http.StatusOK, books)
}

