// Package controllers is the entry point for the application
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/models"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/services"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/utils"
)

// set the struct for the controller
type GameController struct{}

var gameService services.GameService

var validate = validator.New()

// GetGame ... Get one game
// @Summary Get one game by id
// @Description Get one game by id
// @Tags Games
// @Success 200 {object} models.GameResponse
// @Failure 404 {object} object
// @Router /game/{id} [get]
// @Param id path string true "Game ID"
func (c GameController) GetGameById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result, err := gameService.GetGame(id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.GameResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, models.GameResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})

	}
}

// BatchCreateGame ... Batch create multiple games
// @Summary Batch create multiple games
// @Description Batch create multiple games
// @Tags Games
// @Success 200 {object} models.GameResponse
// @Failure 404 {object} object
// @Router /component [post]
// @Param games body []models.GameInput true "Games"
func (c GameController) AddGame() gin.HandlerFunc {
	return func(c *gin.Context) {
		var games []models.GameInput

		if !utils.BindJSONAndValidate(c, &games) {
			return
		}

		for _, component := range games {
			if validationErr := validate.Struct(&component); validationErr != nil {
				c.JSON(http.StatusBadRequest, models.GameResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
				return
			}
		}

		result, err := gameService.BatchCreateGame(games)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.GameResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, models.GameResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
