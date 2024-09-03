package routes

func (r *Route) filesRouter() {

	users := r.v1.Group("/files")
	{
		users.POST("/upload", r.fc.UploadFile)
		users.GET("/:file_name", r.fc.GetFile)
		users.DELETE("/:file_name", r.fc.DeleteFile)
	}
}
