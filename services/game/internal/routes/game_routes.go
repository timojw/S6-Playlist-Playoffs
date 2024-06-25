package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/controllers"
)

func GameRoutes(gameGroup *gin.Engine) {
	game := new(controllers.GameController)
	gameGroup.GET("/game/:id", game.GetGameById())
	gameGroup.POST("/game", game.AddGame())
}
