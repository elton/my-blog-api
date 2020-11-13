package controllers

// Routers routers for the application.
func (s *Server) initializeRouter() {
	s.Routers.GET("/status", HealthCheck)

	v1 := s.Routers.Group("/api/v1")
	{
		v1.GET("/hello/:name", Hello)
	}
}
