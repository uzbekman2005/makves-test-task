package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makves-test-task/config"
	"github.com/makves-test-task/helper"
	"github.com/makves-test-task/pkg/logger"
)

type handlerV1 struct {
	log         *logger.Logger
	cfg         config.Config
	fileManager *helper.FileManager
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger      *logger.Logger
	Cfg         config.Config
	FileManager *helper.FileManager
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:         c.Logger,
		cfg:         c.Cfg,
		fileManager: c.FileManager,
	}
}

func (h *handlerV1) ParseLimitQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("limit", "10"))
}

func (h *handlerV1) ParsePageQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("page", "1"))
}
