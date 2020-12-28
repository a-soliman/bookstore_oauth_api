package app

import (
	"github.com/a-soliman/bookstore_oauth_api/src/domain/access_token"
	"github.com/a-soliman/bookstore_oauth_api/src/http"
	"github.com/a-soliman/bookstore_oauth_api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication starts the application
func StartApplication() {
	dbRepository := db.New()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)

	router.Run(":8080")
}
