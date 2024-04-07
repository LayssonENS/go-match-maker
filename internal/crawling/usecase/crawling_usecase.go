package usecase

import (
	"github.com/LayssonENS/go-match-maker/internal/domain"
)

type crawlingUseCase struct {
	CrawlingRepository domain.CrawlingRepository
}

func NewCrawlingUseCase(CrawlingRepository domain.CrawlingRepository) domain.CrawlingUseCase {
	return &crawlingUseCase{
		CrawlingRepository: CrawlingRepository,
	}
}

func (a *crawlingUseCase) GetByID(id int64) (domain.Crawling, error) {
	crawling, err := a.CrawlingRepository.GetByID(id)
	if err != nil {
		return crawling, err
	}

	return crawling, nil
}

func (a *crawlingUseCase) CreateCrawling(crawling domain.CrawlingRequest) error {
	err := a.CrawlingRepository.CreateCrawling(crawling)
	if err != nil {
		return err
	}

	return nil
}

func (a *crawlingUseCase) GetAllCrawling() ([]domain.Crawling, error) {
	crawling, err := a.CrawlingRepository.GetAllCrawling()
	if err != nil {
		return crawling, err
	}

	return crawling, nil
}
