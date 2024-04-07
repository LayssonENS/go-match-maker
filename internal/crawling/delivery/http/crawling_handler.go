package http

import (
	"net/http"
	"strconv"

	"github.com/LayssonENS/go-match-maker/internal/domain"
	"github.com/gin-gonic/gin"
)

type crawlingHandler struct {
	cUseCase domain.CrawlingUseCase
}

func NewCrawlingHandler(routerGroup *gin.Engine, us domain.CrawlingUseCase) {
	handler := &crawlingHandler{
		cUseCase: us,
	}

	routerGroup.GET("/v1/crawling/:crawlingId", handler.GetByID)
	routerGroup.GET("/v1/crawling/all", handler.GetAllCrawling)
	routerGroup.POST("/v1/crawling", handler.CreateCrawling)
}

// GetByID godoc
// @Summary Get crawling by ID
// @Description get crawling by ID
// @Tags crawling
// @Accept  json
// @Produce  json
// @Param crawlingId path int true "crawling ID"
// @Success 200 {object} domain.crawling
// @Failure 400 {object} domain.ErrorResponse
// @Router /v1/crawling/{id} [get]
func (h *crawlingHandler) GetByID(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("crawlingId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	crawlingId := int64(idParam)

	response, err := h.cUseCase.GetByID(crawlingId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllCrawling godoc
// @Summary Get all crawlings
// @Description get all crawlings
// @Tags crawling
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.crawling
// @Failure 400 {object} domain.ErrorResponse
// @Router /v1/crawling/all [get]
func (h *crawlingHandler) GetAllCrawling(c *gin.Context) {
	response, err := h.cUseCase.GetAllCrawling()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// CreateCrawling godoc
// @Summary Create a new crawling
// @Description create new crawling
// @Tags crawling
// @Accept  json
// @Produce  json
// @Param crawling body domain.crawlingRequest true "Create crawling"
// @Success 201 {object} string
// @Failure 400 {object} domain.ErrorResponse
// @Failure 422 {object} domain.ErrorResponse
// @Router /v1/crawling [post]
func (h *crawlingHandler) CreateCrawling(c *gin.Context) {
	var crawling domain.CrawlingRequest
	if err := c.ShouldBindJSON(&crawling); err != nil {
		c.JSON(http.StatusUnprocessableEntity, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	err := h.cUseCase.CreateCrawling(crawling)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Created"})
}
