package routes

func (r *Route) userRouter() {

	users := r.v1.Group("/users")
	{
		users.GET("/", r.cu.GetUsers)
		users.POST("/", r.cu.CreateUser)
		users.GET("/:userID", r.cu.GetUser)
		users.PUT("/:userID", r.cu.ChangeUser)
		users.DELETE("/:userID", r.cu.DeleteUser)
	}
}
