package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/parodos-dev/parodos-service/docs"
	"github.com/parodos-dev/parodos-service/pkg/workflows"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Parodos API Documentation
// @version 2.0
// @description This is a project to run enterprise workflows on demand

// @contact.name API Support
// @contact.url http://www.parodos.dev
// @contact.email parodos@redhat.com

// @host parodos-dev:8080
// @BasePath /

func main() {
	router := gin.New()

	workflows.InitWorkflows(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run()
	if err != nil {
		log.Fatalf("Cannot initialize the http server: %v", err)
	}
}
