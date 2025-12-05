package summary

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
	// allow injection for tests to vary timeout behavior later if needed
	requestTimeout time.Duration
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service:        service,
		requestTimeout: 2 * time.Second,
	}
}

func (h *Handler) Summary(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.requestTimeout)
	defer cancel()

	snapshot, err := h.service.CollectSnapshot(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, snapshot)
}
