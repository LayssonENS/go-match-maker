package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LayssonENS/go-match-maker/config"
	"github.com/LayssonENS/go-match-maker/config/database"
	crawlingHttpDelivery "github.com/LayssonENS/go-match-maker/internal/crawling/delivery/http"
	"github.com/LayssonENS/go-match-maker/internal/crawling/repository"
	crawlingUCase "github.com/LayssonENS/go-match-maker/internal/crawling/usecase"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title go-match-maker API
// @version 1.0
// @description This is go-match-maker API in Go.

func main() {
	ctx := context.Background()
	log := logrus.New()

	dbInstance, err := database.NewPostgresConnection()
	if err != nil {
		log.WithError(err).Fatal("failed connection database")
		return
	}

	//err = database.DBMigrate(dbInstance, config.GetEnv().DbConfig)
	//if err != nil {
	//	log.WithError(err).Fatal("failed to migrate")
	//	return
	//}

	router := gin.Default()

	postgresCrawlingRepository := crawlingRepository.NewPostgresCrawlingRepository(dbInstance)
	crawlingService := crawlingUCase.NewCrawlingUseCase(postgresCrawlingRepository)

	crawlingHttpDelivery.NewCrawlingHandler(router, crawlingService)
	router.GET("/go-match-maker/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	gin.SetMode(gin.ReleaseMode)
	if config.GetEnv().Debug {
		gin.SetMode(gin.DebugMode)
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.GetEnv().Port),
		Handler: router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down API...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("API Server forced to shutdown:", err)
	}

	log.Println("API Server exiting")
}
