package domain

import "time"

type Match struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsActive  bool      `json:"isActive"`
}

type MatchRequest struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}

type MatchUseCase interface {
	GetByID(id int64) (Match, error)
	CreateMatch(match MatchRequest) error
	GetAllMatch() ([]Match, error)
}

type MatchRepository interface {
	GetByID(id int64) (Match, error)
	CreateMatch(match MatchRequest) error
	GetAllMatch() ([]Match, error)
}
