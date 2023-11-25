package main

import (
	"backend/src/api"
	"backend/src/config"
	"backend/src/repository"
	"backend/src/rpc"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @title           FarAway NFT
// @version         1.0
// @description     FarAway NFT Test API
func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	appConfig := config.NewConfig()

	storage := repository.NewInMemoryStorage(sugar)
	rpcConnection := rpc.NewRpcConnection(sugar, storage)
	handler := api.NewAPIHandler(sugar, storage)
	
	rpcConnection.Start(appConfig)

	if appConfig.IsRelease {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	r.Use(cors.New(corsConfig))

	r.GET("/collections/", handler.GetCollections)
	r.GET("/collections/:collection", handler.GetTokensByCollection)
	r.GET("/tokens/", handler.GetAllTokens)

	r.Run(fmt.Sprintf(":%s", appConfig.GinPort))
}
