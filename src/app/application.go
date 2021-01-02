package app

import (
	"github.com/a-soliman/bookstore_oauth_api/src/http"
	"github.com/a-soliman/bookstore_oauth_api/src/repository/db"
	"github.com/a-soliman/bookstore_oauth_api/src/repository/rest"
	"github.com/a-soliman/bookstore_oauth_api/src/services/access_token_service"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication starts the application
func StartApplication() {

	dbRepository := db.New()
	restUsersRepository := rest.New()
	atService := access_token_service.NewService(restUsersRepository, dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token/", atHandler.Create)

	router.Run(":8080")
}
