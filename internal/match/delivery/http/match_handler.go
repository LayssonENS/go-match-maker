package http

import (
	"net/http"
	"strconv"

	"github.com/LayssonENS/go-match-maker/internal/domain"
	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	mUseCase domain.MatchUseCase
}

func NewMatchHandler(routerGroup *gin.Engine, us domain.MatchUseCase) {
	handler := &MatchHandler{
		mUseCase: us,
	}

	routerGroup.GET("/v1/match/:matchId", handler.GetByID)
	routerGroup.GET("/v1/match/all", handler.GetAllMatch)
	routerGroup.POST("/v1/match", handler.CreateMatch)
}

// GetByID godoc
// @Summary Get match by ID
// @Description get match by ID
// @Tags match
// @Accept  json
// @Produce  json
// @Param matchId path int true "match ID"
// @Success 200 {object} domain.match
// @Failure 400 {object} domain.ErrorResponse
// @Router /v1/match/{id} [get]
func (h *MatchHandler) GetByID(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("matchId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	matchId := int64(idParam)

	response, err := h.mUseCase.GetByID(matchId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllMatch godoc
// @Summary Get all matchs
// @Description get all matchs
// @Tags match
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.match
// @Failure 400 {object} domain.ErrorResponse
// @Router /v1/match/all [get]
func (h *MatchHandler) GetAllMatch(c *gin.Context) {
	response, err := h.mUseCase.GetAllMatch()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// CreateMatch godoc
// @Summary Create a new match
// @Description create new match
// @Tags match
// @Accept  json
// @Produce  json
// @Param match body domain.matchRequest true "Create match"
// @Success 201 {object} string
// @Failure 400 {object} domain.ErrorResponse
// @Failure 422 {object} domain.ErrorResponse
// @Router /v1/match [post]
func (h *MatchHandler) CreateMatch(c *gin.Context) {
	var match domain.MatchRequest
	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusUnprocessableEntity, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	err := h.mUseCase.CreateMatch(match)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{ErrorMessage: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Created"})
}
