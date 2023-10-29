package main

import (
	"api/controllers"
	"api/database"
	"api/events"
	"api/middlewares"
	"api/presenters"
	"api/repositories"
	"api/services"
	"api/usecases"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database, err := database.NewDatabase("postgres://postgres:postgres@db:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Create listeners
	onUserCreate := events.NewOnUserCreate()
	onForgotPassword := events.NewOnForgotPassword()

	// Create the repositories
	sourceRepository := repositories.NewSourceRepository(database)
	userRepository := repositories.NewUserRepository(database)
	roleRepository := repositories.NewRoleRepository(database)
	siteRepository := repositories.NewSiteRepository(database)
	alertRepository := repositories.NewAlertRepository(database)

	// services
	authService := services.NewAuthService("secret", userRepository)

	// Create the usecases
	userUsecase := usecases.NewUserUsecase(authService, userRepository, roleRepository, onUserCreate)
	siteUsecase := usecases.NewSiteUsecase(siteRepository, sourceRepository, alertRepository)
	alertUsecase := usecases.NewAlertUsecase(alertRepository)
	optionUsecase := usecases.NewOptionUsecase(roleRepository, alertRepository)
	authUsecase := usecases.NewAuthUsecase(authService, userRepository, onForgotPassword)

	// Create the presenters
	userPresenter := presenters.NewUserPresenter()
	optionPresenter := presenters.NewOptionPresenter()
	alertPresenter := presenters.NewAlertPresenter()
	sourcePresenter := presenters.NewSourcePresenter()
	sitePresenter := presenters.NewSitePresenter(alertPresenter, sourcePresenter)

	// Create the controllers
	userController := controllers.NewUserController(userUsecase, userPresenter)
	optionController := controllers.NewOptionController(optionUsecase, optionPresenter)
	siteController := controllers.NewSiteController(siteUsecase, sitePresenter)
	alertController := controllers.NewAlertController(alertUsecase, alertPresenter)
	authController := controllers.NewAuthController(authUsecase, authService)

	// Create cors configuration
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins: []string{"https://dashboard.seopilot.dev", "http://dashboard.seopilot.dev"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	})

	router := gin.Default()
	router.Use(corsMiddleware)

	optionsGroup := router.Group("/options", middlewares.Authentication(authService))
	authorizedGroup := router.Group("/", middlewares.Authentication(authService))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/auth", authController.Login)
	router.POST("/auth/forgot-password", authController.ForgotPassword)
	router.POST("/auth/reset-password", authController.ResetPassword)

	authorizedGroup.GET("/users/me", userController.ViewMe)
	authorizedGroup.GET("/users/:id", userController.View)
	authorizedGroup.PATCH("/users/:id", userController.Update)
	authorizedGroup.DELETE("/users/:id", userController.Delete)
	authorizedGroup.GET("/users", userController.ViewAny)
	authorizedGroup.POST("/users", userController.Create)

	authorizedGroup.GET("/alerts/:id", alertController.View)
	authorizedGroup.PATCH("/alerts/:id", alertController.Update)
	authorizedGroup.DELETE("/alerts/:id", alertController.Delete)
	authorizedGroup.GET("/alerts", alertController.ViewAny)
	authorizedGroup.POST("/alerts", alertController.Create)

	authorizedGroup.GET("/sites/:id", siteController.View)
	authorizedGroup.PATCH("/sites/:id", siteController.Update)
	authorizedGroup.DELETE("/sites/:id", siteController.Delete)
	authorizedGroup.GET("/sites", siteController.ViewAny)
	authorizedGroup.POST("/sites", siteController.Create)
	router.POST("/sites/:id/verify", siteController.Verify)

	optionsGroup.GET("/roles", optionController.ViewRoleOptions)
	optionsGroup.GET("/alerts", optionController.ViewAlertOptions)

	router.Run() // listen and serve on 0.0.0.0:8080
}
