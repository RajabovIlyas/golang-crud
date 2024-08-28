package routes

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/controllers"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Route struct {
	g  *gin.Engine
	cu *controllers.UserController
	v1 *gin.RouterGroup
}

func New(gin *gin.Engine, queries *database.Queries) *Route {
	return &Route{gin, controllers.NewUserController(queries), gin.Group("/api/v1")}
}

func (r *Route) PaveRoutes() {

	r.v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.userRouter()
}
