package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/models"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/services"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/utils"
)

type GameController struct {
	service *services.GameService
}

func NewGameController(service *services.GameService) *GameController {
	return &GameController{service: service}
}

var validate = validator.New()

// GetGameById ... Get one game by id
func (ctrl *GameController) GetGameById(c *gin.Context) {
	id := c.Param("id")
	result, err := ctrl.service.GetGame(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.GameResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusOK, models.GameResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
}

// AddGame ... Batch create multiple games
func (ctrl *GameController) AddGame(c *gin.Context) {
	var games []models.GameInput

	if !utils.BindJSONAndValidate(c, &games) {
		return
	}

	for _, game := range games {
		if validationErr := validate.Struct(&game); validationErr != nil {
			c.JSON(http.StatusBadRequest, models.GameResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
	}

	result, err := ctrl.service.BatchCreateGame(games)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.GameResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}

	c.JSON(http.StatusOK, models.GameResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
}
