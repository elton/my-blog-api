package controllers

// Routers routers for the application.
func (s *Server) initializeRouter() {
	s.Router.GET("/status", HealthCheck)

	v1 := s.Router.Group("/api/v1")
	{
		v1.GET("/hello/:name", Hello)
	}
}
