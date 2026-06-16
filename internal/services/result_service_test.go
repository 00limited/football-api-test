package services

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/00limited/football-api/internal/models"
	"github.com/00limited/football-api/internal/repositories"
	"github.com/shopspring/decimal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupResultServiceTest(t *testing.T) (*gorm.DB, *ResultService) {
	t.Helper()
	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", strings.ReplaceAll(t.Name(), "/", "_"))
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.AutoMigrate(&models.Team{}, &models.Player{}, &models.Match{}, &models.MatchResult{}, &models.Goal{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db, NewResultService(repositories.NewMatchRepository(db), repositories.NewResultRepository(db), repositories.NewPlayerRepository(db))
}

func seedMatchData(t *testing.T, db *gorm.DB) (models.Match, models.Player, models.Player, models.Player) {
	t.Helper()
	home := models.Team{Name: "Home FC", FoundedYear: 1900, Address: "A", City: "Home"}
	away := models.Team{Name: "Away FC", FoundedYear: 1901, Address: "B", City: "Away"}
	other := models.Team{Name: "Other FC", FoundedYear: 1902, Address: "C", City: "Other"}
	for _, team := range []*models.Team{&home, &away, &other} {
		if err := db.Create(team).Error; err != nil {
			t.Fatalf("create team: %v", err)
		}
	}
	homePlayer := models.Player{TeamID: home.ID, Name: "Home Striker", HeightCM: decimal.NewFromFloat(180), WeightKG: decimal.NewFromFloat(75), Position: models.PositionForward, JerseyNumber: 9}
	awayPlayer := models.Player{TeamID: away.ID, Name: "Away Striker", HeightCM: decimal.NewFromFloat(181), WeightKG: decimal.NewFromFloat(76), Position: models.PositionForward, JerseyNumber: 10}
	otherPlayer := models.Player{TeamID: other.ID, Name: "Other Striker", HeightCM: decimal.NewFromFloat(182), WeightKG: decimal.NewFromFloat(77), Position: models.PositionForward, JerseyNumber: 11}
	for _, player := range []*models.Player{&homePlayer, &awayPlayer, &otherPlayer} {
		if err := db.Create(player).Error; err != nil {
			t.Fatalf("create player: %v", err)
		}
	}
	match := models.Match{MatchDate: time.Date(2026, 6, 16, 0, 0, 0, 0, time.UTC), MatchTime: "19:30", HomeTeamID: home.ID, AwayTeamID: away.ID, Status: models.MatchStatusScheduled}
	if err := db.Create(&match).Error; err != nil {
		t.Fatalf("create match: %v", err)
	}
	return match, homePlayer, awayPlayer, otherPlayer
}

func TestResultServiceCreateCalculatesScoresAndFinishesMatch(t *testing.T) {
	db, service := setupResultServiceTest(t)
	match, homePlayer, awayPlayer, _ := seedMatchData(t, db)

	result, err := service.Create(match.ID, []models.Goal{{PlayerID: homePlayer.ID, GoalMinute: 10}, {PlayerID: homePlayer.ID, GoalMinute: 40}, {PlayerID: awayPlayer.ID, GoalMinute: 50}})
	if err != nil {
		t.Fatalf("create result: %v", err)
	}
	if result.HomeScore != 2 || result.AwayScore != 1 {
		t.Fatalf("unexpected score: %+v", result)
	}
	updatedMatch, err := repositories.NewMatchRepository(db).GetByID(match.ID)
	if err != nil {
		t.Fatalf("get match: %v", err)
	}
	if updatedMatch.Status != models.MatchStatusFinished {
		t.Fatalf("expected finished status, got %s", updatedMatch.Status)
	}
}

func TestResultServiceRejectsGoalScorerOutsideMatch(t *testing.T) {
	db, service := setupResultServiceTest(t)
	match, _, _, otherPlayer := seedMatchData(t, db)

	_, err := service.Create(match.ID, []models.Goal{{PlayerID: otherPlayer.ID, GoalMinute: 10}})
	if err == nil {
		t.Fatal("expected error for scorer outside the match")
	}
}
