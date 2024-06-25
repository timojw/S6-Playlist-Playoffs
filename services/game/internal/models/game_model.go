package models

type GameType string

type Game struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string `json:"name,omitempty" validate:"required"`
	Deadline string `json:"deadline,omitempty" validate:"required"`
}

type GameInput struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Deadline string `json:"deadline,omitempty" validate:"required"`
}

type GameResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
