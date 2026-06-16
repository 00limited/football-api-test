package handlers

import (
	"net/http"

	"github.com/00limited/football-api/internal/dto/request"
	"github.com/00limited/football-api/internal/models"
	"github.com/00limited/football-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type PlayerHandler struct {
	service *services.PlayerService
}

func NewPlayerHandler(service *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{service: service}
}

func (h *PlayerHandler) ListByTeam(c *gin.Context) {
	teamID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	players, err := h.service.ListByTeamID(teamID)
	if err != nil {
		fail(c, http.StatusBadRequest, "failed to list players", err.Error())
		return
	}
	success(c, http.StatusOK, "players fetched successfully", players)
}

func (h *PlayerHandler) Create(c *gin.Context) {
	teamID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	var req request.CreatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	if errs := validateStruct(req); len(errs) > 0 {
		fail(c, http.StatusBadRequest, "validation failed", errs...)
		return
	}
	player := &models.Player{TeamID: teamID, Name: req.Name, HeightCM: decimal.NewFromFloat(req.HeightCM), WeightKG: decimal.NewFromFloat(req.WeightKG), Position: req.Position, JerseyNumber: req.JerseyNumber}
	if err := h.service.Create(player); err != nil {
		fail(c, http.StatusBadRequest, "failed to create player", err.Error())
		return
	}
	success(c, http.StatusCreated, "player created successfully", player)
}

func (h *PlayerHandler) Get(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	player, err := h.service.GetByID(id)
	if err != nil {
		fail(c, http.StatusNotFound, "failed to get player", err.Error())
		return
	}
	success(c, http.StatusOK, "player fetched successfully", player)
}

func (h *PlayerHandler) Update(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	player, err := h.service.GetByID(id)
	if err != nil {
		fail(c, http.StatusNotFound, "failed to update player", err.Error())
		return
	}
	var req request.UpdatePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	if errs := validateStruct(req); len(errs) > 0 {
		fail(c, http.StatusBadRequest, "validation failed", errs...)
		return
	}
	player.Name = req.Name
	player.HeightCM = decimal.NewFromFloat(req.HeightCM)
	player.WeightKG = decimal.NewFromFloat(req.WeightKG)
	player.Position = req.Position
	player.JerseyNumber = req.JerseyNumber
	if err := h.service.Update(player); err != nil {
		fail(c, http.StatusBadRequest, "failed to update player", err.Error())
		return
	}
	success(c, http.StatusOK, "player updated successfully", player)
}

func (h *PlayerHandler) Delete(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	player, err := h.service.GetByID(id)
	if err != nil {
		fail(c, http.StatusNotFound, "failed to delete player", err.Error())
		return
	}
	if err := h.service.Delete(player); err != nil {
		fail(c, http.StatusBadRequest, "failed to delete player", err.Error())
		return
	}
	success(c, http.StatusOK, "player deleted successfully", nil)
}
