package routes

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/RajabovIlyas/golang-crud/internal/app/controllers/auth-controller"
	filesController "github.com/RajabovIlyas/golang-crud/internal/app/controllers/files-controller"
	"github.com/RajabovIlyas/golang-crud/internal/app/controllers/users-controller"
	"github.com/RajabovIlyas/golang-crud/internal/app/middleware"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Route struct {
	g  *gin.Engine
	uc *usersController.UsersController
	ac *authController.AuthController
	v1 *gin.RouterGroup
	dm *middleware.DeserializeMiddleware
	fc *filesController.FilesController
}

func New(gin *gin.Engine, queries *database.Queries) *Route {
	return &Route{gin,
		usersController.NewUsersController(queries),
		authController.NewAuthController(queries),
		gin.Group(constants.ENDPOINT_V1),
		middleware.NewDeserializeMiddleware(queries),
		filesController.NewFilesController(queries),
	}
}

func (r *Route) PaveRoutes() {

	r.v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.authRouter()

	r.userRouter()

	r.filesRouter()
}
