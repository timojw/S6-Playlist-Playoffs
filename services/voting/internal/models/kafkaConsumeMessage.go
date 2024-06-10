package models

type ConsumeMessage struct {
	UserID             string `json:"user_id"`
	Players            int    `json:"players"`
	SubmissionDeadline string `json:"submission_deadline"`
	GameDeadline       string `json:"game_deadline"`
	Stage              string `json:"stage"`
	AdminPlayer        bool   `json:"admin_player"`
	Private            bool   `json:"private"`
}
