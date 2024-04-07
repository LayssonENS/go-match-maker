package matchRepository

import (
	"database/sql"
	"errors"

	"github.com/LayssonENS/go-match-maker/internal/domain"
)

const dateLayout = "2006-01-02"

type postgresMatchRepo struct {
	DB *sql.DB
}

func NewPostgresMatchRepository(db *sql.DB) domain.MatchRepository {
	return &postgresMatchRepo{
		DB: db,
	}
}

func (p *postgresMatchRepo) GetByID(id int64) (domain.Match, error) {
	var match domain.Match
	err := p.DB.QueryRow(
		"SELECT id, created_at FROM match WHERE id = $1", id).Scan(
		&match.ID, &match.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return match, domain.ErrRegistrationNotFound
		}
		return match, err
	}
	return match, nil
}

func (p *postgresMatchRepo) CreateMatch(match domain.MatchRequest) error {
	query := `INSERT INTO match (birth_date) VALUES ($1)`
	_, err := p.DB.Exec(query, match.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgresMatchRepo) GetAllMatch() ([]domain.Match, error) {
	var matchs []domain.Match

	rows, err := p.DB.Query("SELECT id, name, email, birth_date, created_at FROM match")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var match domain.Match
		err := rows.Scan(
			&match.ID,
			&match.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		matchs = append(matchs, match)
	}

	if len(matchs) == 0 {
		return nil, domain.ErrRegistrationNotFound
	}

	return matchs, nil
}
