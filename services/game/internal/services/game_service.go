// Package services is the business layer between the controllers and the repositories, and it contains the business logic for the application
package services

import (
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/models"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/repositories"
)

type GameService struct{}

var gameRepository repositories.GameRepository

func (c GameService) BatchCreateGame(games []models.GameInput) (interface{}, error) {
	interfaceSlice := make([]interface{}, len(games))
	for i, v := range games {
		interfaceSlice[i] = models.Game{
			Name:     v.Name,
			Deadline: v.Deadline,
		}
	}

	result, err := gameRepository.BatchInsert(interfaceSlice)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c GameService) GetGame(id string) (interface{}, error) {

	result, err := gameRepository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c GameService) FindPaginatedGame(pagination models.Pagination) (interface{}, int, error) {
	result, total, err := gameRepository.FindPaginated(pagination)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (c GameService) CheckAlreadyExistingGames(gameNames []string) ([]string, error) {
	return gameRepository.CheckAlreadyExistingGames(gameNames)
}
