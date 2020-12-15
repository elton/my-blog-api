package controllers

import "github.com/elton/my-blog-api/api/middlewares"

// Routers routers for the application.
func (s *Server) initializeRouter() {
	s.Router.GET("/status", HealthCheck)

	v1 := s.Router.Group("/api/v1")
	{
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

		// post routers
		v1.POST("/posts", middlewares.SetMiddlewareJSON(), s.CreatePost)
		v1.GET("/posts/", middlewares.SetMiddlewareJSON(), s.FindPostsBy)
		v1.GET("/posts/:id", middlewares.SetMiddlewareJSON(), s.FindPostByID)
		v1.GET("/posts", middlewares.SetMiddlewareJSON(), s.FindPosts)
		v1.PUT("/posts/:id", middlewares.SetMiddlewareJSON(), s.UpdatePost)
		v1.DELETE("/posts/:id", middlewares.SetMiddlewareJSON(), s.DeletePost)

		// comment routers
		v1.POST("/comments", middlewares.SetMiddlewareJSON(), s.CreateComment)
		v1.GET("/comments/:id", middlewares.SetMiddlewareJSON(), s.FindCommentByID)
		v1.GET("/comments/", middlewares.SetMiddlewareJSON(), s.FindCommentsBy)
		v1.PUT("/comments/:id", middlewares.SetMiddlewareJSON(), s.UpdateComment)
		v1.DELETE("/comments/:id", middlewares.SetMiddlewareJSON(), s.DeleteComment)

		// like routers
		v1.POST("/likes", middlewares.SetMiddlewareJSON(), s.CreateLike)
		v1.GET("/likes/:id", middlewares.SetMiddlewareJSON(), s.FindLikeByID)
		v1.GET("/likes/", middlewares.SetMiddlewareJSON(), s.FindLikesBy)
		v1.PUT("/likes/:id", middlewares.SetMiddlewareJSON(), s.UpdateLike)
		v1.DELETE("/likes/:id", middlewares.SetMiddlewareJSON(), s.DeleteLike)

		// Auth routers
		v1.POST("/login", middlewares.SetMiddlewareJSON(), s.Login)
		v1.POST("/register", middlewares.SetMiddlewareJSON(), s.CreateUser)
	}
}
