package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/s-take/knowledge-like/handler"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// vue.js static files
	e.Static("/", "public/")

	e.POST("/signup", handler.Signup)
	e.POST("/signin", handler.Signin)

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(handler.Config))

	api.GET("/knowledge", handler.GetKnowledges)
	api.POST("/knowledge", handler.AddKnowledge)
	api.DELETE("/knowledge/:id", handler.DeleteKnowledge)

	api.POST("/like", handler.AddLike)
	api.GET("/like", handler.GetLikes)
	api.GET("/like/my", handler.GetLikesByUserID)
	api.GET("/like/:id", handler.GetLikesByKnowledgeID)
	api.DELETE("/like/my", handler.DeleteLikeByUserID)

	return e
}
