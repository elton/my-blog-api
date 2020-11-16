package controllers

import "github.com/elton/my-blog-api/api/middlewares"

// Routers routers for the application.
func (s *Server) initializeRouter() {
	s.Router.GET("/status", HealthCheck)

	v1 := s.Router.Group("/api/v1")
	{
		v1.GET("/hello/:name", s.Hello)

		// category routers
		v1.POST("/categories", middlewares.SetMiddlewareJSON(), s.CreateCategory)
		v1.GET("/categories/:id", middlewares.SetMiddlewareJSON(), s.FindCategoryByID)
		v1.GET("/categories/", middlewares.SetMiddlewareJSON(), s.FindCategoryByName)
		v1.GET("/categories", middlewares.SetMiddlewareJSON(), s.FindCategories)
		v1.PUT("/categories", middlewares.SetMiddlewareJSON(), s.UpdateCategory)
		v1.DELETE("/categories/:id", middlewares.SetMiddlewareJSON(), s.DeleteCategory)

		// user routers
		v1.POST("/users", middlewares.SetMiddlewareJSON(), s.CreateUser)
		v1.GET("/users/:id", middlewares.SetMiddlewareJSON(), s.FindUserByID)
		v1.GET("/users", middlewares.SetMiddlewareJSON(), s.FindUsers)
		v1.GET("/users/", middlewares.SetMiddlewareJSON(), s.FindUsersBy)
		v1.PUT("/users/:id", middlewares.SetMiddlewareJSON(), s.UpdateUser)
		v1.DELETE("/users/:id", middlewares.SetMiddlewareJSON(), s.DeleteUser)
	}
}
