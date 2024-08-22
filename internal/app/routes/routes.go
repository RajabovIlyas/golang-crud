package routes

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/controllers"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/gin-gonic/gin"
)

type Route struct {
	g  *gin.Engine
	cu *controllers.UserController
}

func New(gin *gin.Engine, queries *database.Queries) *Route {
	return &Route{gin, controllers.NewUserController(queries)}
}

func (r *Route) PaveRoutes() {

	v1 := r.g.Group("/api/v1")
	{

	}
	users := v1.Group("/users")
	{
		users.GET("/", r.cu.GetUsers)
		users.POST("/", r.cu.CreateUser)
		users.GET("/:userID", r.cu.GetUser)
		users.PUT("/:userID", r.cu.ChangeUser)
		users.DELETE("/:userID", r.cu.DeleteUser)
	}
}
