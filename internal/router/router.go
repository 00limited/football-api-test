package router

import (
	"github.com/00limited/football-api/internal/config"
	"github.com/00limited/football-api/internal/handlers"
	"github.com/00limited/football-api/internal/middleware"
	"github.com/00limited/football-api/internal/repositories"
	"github.com/00limited/football-api/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(cfg *config.Config, db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Static("/uploads", "./uploads")

	adminRepo := repositories.NewAdminRepository(db)
	teamRepo := repositories.NewTeamRepository(db)
	playerRepo := repositories.NewPlayerRepository(db)
	matchRepo := repositories.NewMatchRepository(db)
	resultRepo := repositories.NewResultRepository(db)
	reportRepo := repositories.NewReportRepository(db)

	authHandler := handlers.NewAuthHandler(cfg, adminRepo)
	teamHandler := handlers.NewTeamHandler(services.NewTeamService(teamRepo))
	playerHandler := handlers.NewPlayerHandler(services.NewPlayerService(playerRepo, teamRepo))
	matchHandler := handlers.NewMatchHandler(services.NewMatchService(matchRepo, teamRepo))
	resultHandler := handlers.NewResultHandler(services.NewResultService(matchRepo, resultRepo, playerRepo))
	reportHandler := handlers.NewReportHandler(services.NewReportService(reportRepo))

	api := r.Group("/api/v1")
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	protected := api.Group("")
	protected.Use(middleware.JWT(cfg))
	{
		protected.GET("/teams", teamHandler.List)
		protected.POST("/teams", teamHandler.Create)
		protected.GET("/teams/:id", teamHandler.Get)
		protected.PUT("/teams/:id", teamHandler.Update)
		protected.DELETE("/teams/:id", teamHandler.Delete)
		protected.POST("/teams/:id/logo", teamHandler.UploadLogo)

		protected.GET("/teams/:id/players", playerHandler.ListByTeam)
		protected.POST("/teams/:id/players", playerHandler.Create)
		protected.GET("/players/:id", playerHandler.Get)
		protected.PUT("/players/:id", playerHandler.Update)
		protected.DELETE("/players/:id", playerHandler.Delete)

		protected.GET("/matches", matchHandler.List)
		protected.POST("/matches", matchHandler.Create)
		protected.GET("/matches/:id", matchHandler.Get)
		protected.PUT("/matches/:id", matchHandler.Update)
		protected.DELETE("/matches/:id", matchHandler.Delete)

		protected.POST("/matches/:id/result", resultHandler.Create)
		protected.GET("/matches/:id/result", resultHandler.Get)

		protected.GET("/reports/matches", reportHandler.List)
		protected.GET("/reports/matches/:id", reportHandler.Get)
	}

	return r
}
