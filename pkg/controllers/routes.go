package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunRoutes() error {
	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	router.GET("/ping", PingPong)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	userG := router.Group("/users")
	{
		userG.GET("", GetAllUsers)
		userG.GET("/:id", GetUserByID)
		userG.POST("", CreateUser)
		userG.PUT("/:id", UpdateUser)
		userG.DELETE("/:id")
		userG.PATCH("/:id")
	}

	/*
		/operations -> checkUserAuthentication -> GetAllOperations
	*/
	operationsG := router.Group("/operations", checkUserAuthentication)
	{
		operationsG.GET("", GetAllOperations)
		operationsG.POST("")
		operationsG.GET("/:id", GetOperationByID)
		operationsG.PUT("/:id")
		operationsG.DELETE("/:id")
	}

	operationTypesG := router.Group("/operation-types")
	{
		operationTypesG.GET("")
		operationTypesG.POST("")
		operationTypesG.GET("/:id")
		operationTypesG.PUT("/:id")
		operationTypesG.DELETE("/:id")
	}

	operationCategoriesG := router.Group("/operation-categories")
	{
		operationCategoriesG.GET("")
		operationCategoriesG.POST("")
		operationCategoriesG.GET("/:id")
		operationCategoriesG.PUT("/:id")
		operationCategoriesG.DELETE("/:id")
	}

	err := router.Run(":8181")
	if err != nil {
		return err
	}

	return nil
}

func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
