package handlers

import (
	"math"
	"net/http"
	"strconv"

	resp "github.com/00limited/football-api/internal/dto/response"
	"github.com/gin-gonic/gin"
)

func success(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, resp.APIResponse{Status: "success", Message: message, Data: data})
}

func fail(c *gin.Context, code int, message string, errs ...string) {
	c.JSON(code, resp.APIResponse{Status: "error", Message: message, Errors: errs})
}

func parseUintParam(c *gin.Context, key string) (uint, bool) {
	value, err := strconv.ParseUint(c.Param(key), 10, 64)
	if err != nil {
		fail(c, http.StatusBadRequest, "invalid parameter", key+" must be a positive integer")
		return 0, false
	}
	if value > math.MaxUint {
		fail(c, http.StatusBadRequest, "invalid parameter", key+" exceeds supported size")
		return 0, false
	}
	return uint(value), true
}
