package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/controllers"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/repositories"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/services"
)

func GameRoutes(router *gin.Engine) {
	gameRepo := repositories.GameRepository{}
	gameService := services.NewGameService(gameRepo)
	gameController := controllers.NewGameController(gameService)

	router.GET("/game/:id", gameController.GetGameById)
	router.POST("/game", gameController.AddGame)
}
