package routes

import (
	authController "github.com/RajabovIlyas/golang-crud/internal/app/controllers/auth-controller"
	"github.com/RajabovIlyas/golang-crud/internal/app/controllers/users-controller"
	"github.com/RajabovIlyas/golang-crud/internal/app/middleware"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Route struct {
	g  *gin.Engine
	uc *usersController.UsersController
	ac *authController.AuthController
	v1 *gin.RouterGroup
	dm *middleware.DeserializeMiddleware
}

func New(gin *gin.Engine, queries *database.Queries) *Route {
	return &Route{gin,
		usersController.NewUsersController(queries),
		authController.NewAuthController(queries),
		gin.Group("/api/v1"),
		middleware.NewDeserializeMiddleware(queries),
	}
}

func (r *Route) PaveRoutes() {

	r.v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.authRouter()

	r.userRouter()
}
