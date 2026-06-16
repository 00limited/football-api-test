package handlers

import (
	"net/http"

	"github.com/00limited/football-api/internal/dto/request"
	"github.com/00limited/football-api/internal/models"
	"github.com/00limited/football-api/internal/services"
	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	service *services.MatchService
}

func NewMatchHandler(service *services.MatchService) *MatchHandler {
	return &MatchHandler{service: service}
}

func (h *MatchHandler) List(c *gin.Context) {
	matches, err := h.service.List()
	if err != nil {
		fail(c, http.StatusInternalServerError, "failed to list matches", err.Error())
		return
	}
	success(c, http.StatusOK, "matches fetched successfully", matches)
}

func (h *MatchHandler) Create(c *gin.Context) {
	var req request.CreateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	if errs := validateStruct(req); len(errs) > 0 {
		fail(c, http.StatusBadRequest, "validation failed", errs...)
		return
	}
	matchDate, err := h.service.ParseDate(req.MatchDate)
	if err != nil {
		fail(c, http.StatusBadRequest, "validation failed", err.Error())
		return
	}
	match := &models.Match{MatchDate: matchDate, MatchTime: req.MatchTime, HomeTeamID: req.HomeTeamID, AwayTeamID: req.AwayTeamID}
	if err := h.service.Create(match); err != nil {
		fail(c, http.StatusBadRequest, "failed to create match", err.Error())
		return
	}
	success(c, http.StatusCreated, "match created successfully", match)
}

func (h *MatchHandler) Get(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	match, err := h.service.GetByID(id)
	if err != nil {
		fail(c, http.StatusNotFound, "failed to get match", err.Error())
		return
	}
	success(c, http.StatusOK, "match fetched successfully", match)
}

func (h *MatchHandler) Update(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	match, err := h.service.GetByID(id)
	if err != nil {
		fail(c, http.StatusNotFound, "failed to update match", err.Error())
		return
	}
	var req request.UpdateMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	if errs := validateStruct(req); len(errs) > 0 {
		fail(c, http.StatusBadRequest, "validation failed", errs...)
		return
	}
	matchDate, err := h.service.ParseDate(req.MatchDate)
	if err != nil {
		fail(c, http.StatusBadRequest, "validation failed", err.Error())
		return
	}
	match.MatchDate = matchDate
	match.MatchTime = req.MatchTime
	match.HomeTeamID = req.HomeTeamID
	match.AwayTeamID = req.AwayTeamID
	if err := h.service.Update(match); err != nil {
		fail(c, http.StatusBadRequest, "failed to update match", err.Error())
		return
	}
	success(c, http.StatusOK, "match updated successfully", match)
}

func (h *MatchHandler) Delete(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	match, err := h.service.GetByID(id)
	if err != nil {
		fail(c, http.StatusNotFound, "failed to delete match", err.Error())
		return
	}
	if err := h.service.Delete(match); err != nil {
		fail(c, http.StatusBadRequest, "failed to delete match", err.Error())
		return
	}
	success(c, http.StatusOK, "match deleted successfully", nil)
}
