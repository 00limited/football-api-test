package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/00limited/football-api/internal/dto/request"
	"github.com/00limited/football-api/internal/models"
	"github.com/00limited/football-api/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeamHandler struct {
	service *services.TeamService
}

func NewTeamHandler(service *services.TeamService) *TeamHandler {
	return &TeamHandler{service: service}
}

func (h *TeamHandler) List(c *gin.Context) {
	teams, err := h.service.List()
	if err != nil {
		fail(c, http.StatusInternalServerError, "failed to list teams", err.Error())
		return
	}
	success(c, http.StatusOK, "teams fetched successfully", teams)
}

func (h *TeamHandler) Create(c *gin.Context) {
	var req request.CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	if errs := validateStruct(req); len(errs) > 0 {
		fail(c, http.StatusBadRequest, "validation failed", errs...)
		return
	}
	team := &models.Team{Name: req.Name, FoundedYear: req.FoundedYear, Address: req.Address, City: req.City}
	if err := h.service.Create(team); err != nil {
		fail(c, http.StatusBadRequest, "failed to create team", err.Error())
		return
	}
	success(c, http.StatusCreated, "team created successfully", team)
}

func (h *TeamHandler) Get(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	team, err := h.service.GetByID(id)
	if err != nil {
		status := http.StatusInternalServerError
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
		}
		fail(c, status, "failed to get team", err.Error())
		return
	}
	success(c, http.StatusOK, "team fetched successfully", team)
}

func (h *TeamHandler) Update(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	team, err := h.service.GetByID(id)
	if err != nil {
		fail(c, http.StatusNotFound, "failed to update team", err.Error())
		return
	}
	var req request.UpdateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	if errs := validateStruct(req); len(errs) > 0 {
		fail(c, http.StatusBadRequest, "validation failed", errs...)
		return
	}
	team.Name = req.Name
	team.FoundedYear = req.FoundedYear
	team.Address = req.Address
	team.City = req.City
	if err := h.service.Update(team); err != nil {
		fail(c, http.StatusBadRequest, "failed to update team", err.Error())
		return
	}
	success(c, http.StatusOK, "team updated successfully", team)
}

func (h *TeamHandler) Delete(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	team, err := h.service.GetByID(id)
	if err != nil {
		fail(c, http.StatusNotFound, "failed to delete team", err.Error())
		return
	}
	if err := h.service.Delete(team); err != nil {
		fail(c, http.StatusBadRequest, "failed to delete team", err.Error())
		return
	}
	success(c, http.StatusOK, "team deleted successfully", nil)
}

func (h *TeamHandler) UploadLogo(c *gin.Context) {
	id, ok := parseUintParam(c, "id")
	if !ok {
		return
	}
	team, err := h.service.GetByID(id)
	if err != nil {
		fail(c, http.StatusNotFound, "failed to upload logo", err.Error())
		return
	}
	file, err := c.FormFile("logo")
	if err != nil {
		fail(c, http.StatusBadRequest, "failed to upload logo", "logo file is required")
		return
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext == "" {
		ext = ".bin"
	}
	fileName := "team-" + strconv.FormatUint(uint64(team.ID), 10) + "-" + strconv.FormatInt(time.Now().UnixNano(), 10) + ext
	path := filepath.Join("uploads", fileName)
	if err := os.MkdirAll("uploads", 0o755); err != nil {
		fail(c, http.StatusInternalServerError, "failed to upload logo", err.Error())
		return
	}
	if err := c.SaveUploadedFile(file, path); err != nil {
		fail(c, http.StatusInternalServerError, "failed to upload logo", err.Error())
		return
	}
	team.LogoURL = "/" + filepath.ToSlash(path)
	if err := h.service.Update(team); err != nil {
		fail(c, http.StatusInternalServerError, "failed to upload logo", err.Error())
		return
	}
	success(c, http.StatusOK, "team logo uploaded successfully", team)
}
