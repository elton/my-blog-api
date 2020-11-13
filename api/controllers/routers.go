package controllers

import "github.com/elton/my-blog-api/api/middlewares"

// Routers routers for the application.
func (s *Server) initializeRouter() {
	s.Router.GET("/status", HealthCheck)

	v1 := s.Router.Group("/api/v1")
	{
		v1.GET("/hello/:name", s.Hello)
		v1.POST("/categories", middlewares.SetMiddlewareJSON(), s.CreateCategory)
	}
}
