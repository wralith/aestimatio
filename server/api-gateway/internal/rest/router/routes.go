package router

func (r router) initRoutes() {
	r.Echo.POST("/auth/login", r.authHandler.Login)
	r.Echo.POST("/auth/register", r.authHandler.Register)

	g := r.Echo.Group("/tasks", r.authHandler.Authenticate)

	g.POST("", r.taskHandler.Create)
	g.GET("/:id", r.taskHandler.Get)
	g.DELETE("/:id", r.taskHandler.Delete)
	g.PUT("/:id", r.taskHandler.Switch)
	g.GET("/list", r.taskHandler.List)

	d := r.Echo.Group("/docs")
	d.File("/spec", "docs/swagger.yaml")
	d.File("", "docs/docs.html")
}
