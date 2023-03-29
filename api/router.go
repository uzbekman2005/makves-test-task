package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/makves-test-task/config"
	"github.com/makves-test-task/helper"
	"github.com/makves-test-task/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/makves-test-task/api/docs" //swag
	v1 "github.com/makves-test-task/api/handlers/v1"
)

// Option ...
type Option struct {
	Conf        config.Config
	Logger      *logger.Logger
	FileManager *helper.FileManager
}

// New ...
// @title           Makves Test Project API endpoints
// @version         1.0
// @description     Here QA can test and frontend or mobile developers can get information of API endpoints

// @host     localhost:8000
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:      option.Logger,
		Cfg:         option.Conf,
		FileManager: option.FileManager,
	})

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	router.GET("/", handlerV1.AppIsRunning)
	router.GET("/get-all", handlerV1.GetAllInfo)
	router.GET("/get-by-id/:id", handlerV1.GetById)
	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
