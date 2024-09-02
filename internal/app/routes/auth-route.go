package routes

func (r *Route) authRouter() {

	auth := r.v1.Group("/auth")
	{
		auth.POST("/login", r.ac.Login)
		auth.POST("/registration", r.ac.Registration)
		auth.POST("/logout", r.dm.DeserializeUser(), r.ac.LogoutMe)
		auth.GET("/auth-me", r.dm.DeserializeUser(), r.ac.AuthMe)
		auth.POST("/refresh-token", r.ac.RefreshToken)
	}
}
