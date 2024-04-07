package usecase

import (
	"github.com/LayssonENS/go-match-maker/internal/domain"
)

type matchUseCase struct {
	matchRepository domain.MatchRepository
}

func NewMatchUseCase(matchRepository domain.MatchRepository) domain.MatchUseCase {
	return &matchUseCase{
		matchRepository: matchRepository,
	}
}

func (a *matchUseCase) GetByID(id int64) (domain.Match, error) {
	match, err := a.matchRepository.GetByID(id)
	if err != nil {
		return match, err
	}

	return match, nil
}

func (a *matchUseCase) CreateMatch(match domain.MatchRequest) error {
	err := a.matchRepository.CreateMatch(match)
	if err != nil {
		return err
	}

	return nil
}

func (a *matchUseCase) GetAllMatch() ([]domain.Match, error) {
	match, err := a.matchRepository.GetAllMatch()
	if err != nil {
		return match, err
	}

	return match, nil
}
