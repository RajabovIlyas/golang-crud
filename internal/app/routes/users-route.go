package routes

func (r *Route) userRouter() {

	users := r.v1.Group("/users")
	{
		users.GET("/", r.uc.GetUsers)
		users.POST("/", r.uc.CreateUser)
		users.POST("/password/:userID", r.uc.UpdateUserPassword)
		users.GET("/:userID", r.uc.GetUser)
		users.PUT("/:userID", r.uc.UpdateUser)
		users.DELETE("/:userID", r.uc.DeleteUser)
	}
}
