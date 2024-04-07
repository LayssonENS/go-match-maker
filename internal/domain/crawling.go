package domain

import "time"

type Crawling struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsActive  bool      `json:"isActive"`
}

type CrawlingRequest struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}

type CrawlingUseCase interface {
	GetByID(id int64) (Crawling, error)
	CreateCrawling(crawling CrawlingRequest) error
	GetAllCrawling() ([]Crawling, error)
}

type CrawlingRepository interface {
	GetByID(id int64) (Crawling, error)
	CreateCrawling(crawling CrawlingRequest) error
	GetAllCrawling() ([]Crawling, error)
}
