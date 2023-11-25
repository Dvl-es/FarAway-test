package api

import (
	"backend/src/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	logger  *zap.SugaredLogger
	storage repository.Repository
}

func NewAPIHandler(
	logger *zap.SugaredLogger,
	storage repository.Repository) *Handler {
	return &Handler{
		logger:  logger,
		storage: storage,
	}
}

func (h *Handler) GetCollections(c *gin.Context) {
	collections := h.storage.GetAllCollections()

	c.JSON(http.StatusOK, collections)
}

func (h *Handler) GetAllTokens(c *gin.Context) {
	tokens := h.storage.GetAllTokens()

	c.JSON(http.StatusOK, tokens)
}

func (h *Handler) GetTokensByCollection(c *gin.Context) {
	collection := c.Param("collection")
	tokens := h.storage.GetTokensByCollection(common.HexToAddress(collection))
	c.JSON(http.StatusOK, tokens)
}
