package handlers

import (
	"errors"
	"net/http"

	"github.com/00limited/football-api/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) List(c *gin.Context) {
	reports, err := h.service.List()
	if err != nil {
		fail(c, http.StatusInternalServerError, "failed to list reports", err.Error())
		return
	}
	success(c, http.StatusOK, "reports fetched successfully", reports)
}

func (h *ReportHandler) Get(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	report, err := h.service.GetByMatchID(id)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		fail(c, status, "failed to get report", err.Error())
		return
	}
	success(c, http.StatusOK, "report fetched successfully", report)
}
