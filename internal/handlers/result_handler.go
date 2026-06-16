package handlers

import (
	"net/http"

	"github.com/00limited/football-api/internal/dto/request"
	"github.com/00limited/football-api/internal/models"
	"github.com/00limited/football-api/internal/services"
	"github.com/gin-gonic/gin"
)

type ResultHandler struct {
	service *services.ResultService
}

func NewResultHandler(service *services.ResultService) *ResultHandler {
	return &ResultHandler{service: service}
}

func (h *ResultHandler) Create(c *gin.Context) {
	matchID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	var req request.CreateResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	if errs := validateStruct(req); len(errs) > 0 {
		fail(c, http.StatusBadRequest, "validation failed", errs...)
		return
	}
	goals := make([]models.Goal, 0, len(req.Goals))
	for _, goal := range req.Goals {
		goals = append(goals, models.Goal{PlayerID: goal.PlayerID, GoalMinute: goal.GoalMinute})
	}
	result, err := h.service.Create(matchID, goals)
	if err != nil {
		fail(c, http.StatusBadRequest, "failed to create result", err.Error())
		return
	}
	success(c, http.StatusCreated, "match result created successfully", result)
}

func (h *ResultHandler) Get(c *gin.Context) {
	matchID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	result, err := h.service.GetByMatchID(matchID)
	if err != nil {
		fail(c, http.StatusNotFound, "failed to get result", err.Error())
		return
	}
	success(c, http.StatusOK, "match result fetched successfully", result)
}
