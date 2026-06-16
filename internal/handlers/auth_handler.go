package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/00limited/football-api/internal/config"
	"github.com/00limited/football-api/internal/dto/request"
	"github.com/00limited/football-api/internal/models"
	"github.com/00limited/football-api/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	cfg       *config.Config
	adminRepo *repositories.AdminRepository
}

func NewAuthHandler(cfg *config.Config, adminRepo *repositories.AdminRepository) *AuthHandler {
	return &AuthHandler{cfg: cfg, adminRepo: adminRepo}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	if errs := validateStruct(req); len(errs) > 0 {
		fail(c, http.StatusBadRequest, "validation failed", errs...)
		return
	}
	if _, err := h.adminRepo.GetByUsername(req.Username); err == nil {
		fail(c, http.StatusConflict, "register failed", "username already exists")
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		fail(c, http.StatusInternalServerError, "register failed", err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		fail(c, http.StatusInternalServerError, "register failed", err.Error())
		return
	}
	admin := &models.Admin{Username: req.Username, Password: string(hashedPassword)}
	if err := h.adminRepo.Create(admin); err != nil {
		fail(c, http.StatusInternalServerError, "register failed", err.Error())
		return
	}
	success(c, http.StatusCreated, "admin registered successfully", gin.H{"id": admin.ID, "username": admin.Username})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fail(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	if errs := validateStruct(req); len(errs) > 0 {
		fail(c, http.StatusBadRequest, "validation failed", errs...)
		return
	}
	admin, err := h.adminRepo.GetByUsername(req.Username)
	if err != nil {
		fail(c, http.StatusUnauthorized, "login failed", "invalid credentials")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		fail(c, http.StatusUnauthorized, "login failed", "invalid credentials")
		return
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      admin.ID,
		"username": admin.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}).SignedString([]byte(h.cfg.JWTSecret))
	if err != nil {
		fail(c, http.StatusInternalServerError, "login failed", err.Error())
		return
	}

	success(c, http.StatusOK, "login successful", gin.H{"token": token})
}
