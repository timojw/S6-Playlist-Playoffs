package models

type ConsumeMessage struct {
	GameID   int  `json:"GameID"`
	GetVotes bool `json:"GetVotes"`
}
