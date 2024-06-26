package services

import (
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/models"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/repositories"
)

type GameService struct {
	repo repositories.GameRepository
}

func NewGameService(repo repositories.GameRepository) *GameService {
	return &GameService{repo: repo}
}

func (s *GameService) BatchCreateGame(games []models.GameInput) (interface{}, error) {
	interfaceSlice := make([]interface{}, len(games))
	for i, v := range games {
		interfaceSlice[i] = models.Game{
			Name:     v.Name,
			Deadline: v.Deadline,
		}
	}

	result, err := s.repo.BatchInsert(interfaceSlice)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *GameService) GetGame(id string) (interface{}, error) {
	result, err := s.repo.FindOne(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *GameService) FindPaginatedGame(pagination models.Pagination) (interface{}, int, error) {
	result, total, err := s.repo.FindPaginated(pagination)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (s *GameService) CheckAlreadyExistingGames(gameNames []string) ([]string, error) {
	return s.repo.CheckAlreadyExistingGames(gameNames)
}
