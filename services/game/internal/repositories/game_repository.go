// Package repositories holds the database operations for the application
package repositories

import (
	// configs "github.com/timojw/S6-Playlist-Playoffs/services/game/config"
	"github.com/timojw/S6-Playlist-Playoffs/services/game/internal/models"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type GameRepository struct{}

// var gameCollection *mongo.Collection = configs.GetCollection(configs.DB, "games")

// BatchInsert inserts multiple games into the database
func (c GameRepository) BatchInsert(games []interface{}) (interface{}, error) {
	// Mock response for testing without MongoDB
	return nil, nil
}

// FindOne finds a game by id
func (c GameRepository) FindOne(id string) (interface{}, error) {
	// Mock response for testing without MongoDB
	return models.Game{ID: id, Name: "Test Game"}, nil
}

// FindPaginated finds games paginated
func (c GameRepository) FindPaginated(pagination models.Pagination) ([]models.Game, int, error) {
	// Mock response for testing without MongoDB
	return []models.Game{{ID: "1", Name: "Test Game"}}, 1, nil
}

// CheckAlreadyExistingGames checks if the games already exist in the database
func (c GameRepository) CheckAlreadyExistingGames(gameNames []string) ([]string, error) {
	// Mock response for testing without MongoDB
	return []string{"Test Game"}, nil
}

// var gameCollection *mongo.Collection = configs.GetCollection(configs.DB, "games")

// // BatchInsert inserts multiple games into the database
// func (c GameRepository) BatchInsert(games []interface{}) (*mongo.BulkWriteResult, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	var operations []mongo.WriteModel
// 	for _, game := range games {
// 		filter := bson.M{"name": game.(models.Game).Name}
// 		update := bson.M{"$set": game}
// 		operation := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
// 		operations = append(operations, operation)
// 	}

// 	result, err := gameCollection.BulkWrite(ctx, operations)
// 	return result, err
// }

// // FindOne finds a game by id
// func (c GameRepository) FindOne(id string) (interface{}, error) {
// 	var game models.Game
// 	objectId, _ := primitive.ObjectIDFromHex(id)
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	err := gameCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&game)
// 	return game, err
// }

// // FindPaginated finds games paginated
// func (c GameRepository) FindPaginated(pagination models.Pagination) ([]models.Game, int, error) {
// 	var games []models.Game
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	totalCount, err := gameCollection.CountDocuments(ctx, pagination.Filter)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	if totalCount == 0 {
// 		return games, 0, nil
// 	}

// 	if pagination.Skip > int(totalCount) {
// 		pagination.Skip = int(totalCount)
// 	}

// 	options := options.Find()
// 	options.SetSkip(int64((pagination.Skip - 1) * pagination.Limit))
// 	options.SetLimit(int64(pagination.Limit))
// 	options.SetSort(pagination.Sort)

// 	cursor, err := gameCollection.Find(ctx, pagination.Filter, options)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	if err = cursor.All(context.TODO(), &games); err != nil {
// 		panic(err)
// 	}

// 	return games, int(totalCount), nil
// }

// // CheckAlreadyExistingGames checks if the games already exist in the database
// func (c GameRepository) CheckAlreadyExistingGames(gameNames []string) ([]string, error) {

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	var games []models.Game

// 	cursor, err := gameCollection.Find(ctx, bson.M{"name": bson.M{"$in": gameNames}})
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = cursor.All(context.Background(), &games)
// 	if err != nil {
// 		return nil, err
// 	}

// 	existingGames := make([]string, len(games))
// 	for i, game := range games {
// 		existingGames[i] = game.Name
// 	}

// 	return existingGames, nil
// }
