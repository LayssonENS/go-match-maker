package crawlingRepository

import (
	"database/sql"

	"github.com/LayssonENS/go-match-maker/internal/domain"
)

const dateLayout = "2006-01-02"

type postgresCrawlingRepo struct {
	DB *sql.DB
}

func NewPostgresCrawlingRepository(db *sql.DB) domain.CrawlingRepository {
	return &postgresCrawlingRepo{}
}

func (p *postgresCrawlingRepo) GetByID(id int64) (domain.Crawling, error) {
	var crawling domain.Crawling
	err := p.DB.QueryRow(
		"SELECT id, created_at FROM crawling WHERE id = $1", id).Scan(
		&crawling.ID, &crawling.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return crawling, domain.ErrRegistrationNotFound
		}
		return crawling, err
	}
	return crawling, nil
}

func (p *postgresCrawlingRepo) CreateCrawling(crawling domain.CrawlingRequest) error {
	query := `INSERT INTO crawling (name, email, birth_date) VALUES ($1, $2, $3)`
	_, err := p.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgresCrawlingRepo) GetAllCrawling() ([]domain.Crawling, error) {
	var crawlings []domain.Crawling

	rows, err := p.DB.Query("SELECT id, name, email, birth_date, created_at FROM crawling")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var crawling domain.Crawling
		err := rows.Scan(
			&crawling.ID,

			&crawling.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		crawlings = append(crawlings, crawling)
	}

	if len(crawlings) == 0 {
		return nil, domain.ErrRegistrationNotFound
	}

	return crawlings, nil
}
