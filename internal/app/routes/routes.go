package routes

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/RajabovIlyas/golang-crud/internal/app/controllers/auth-controller"
	filesController "github.com/RajabovIlyas/golang-crud/internal/app/controllers/files-controller"
	"github.com/RajabovIlyas/golang-crud/internal/app/controllers/users-controller"
	"github.com/RajabovIlyas/golang-crud/internal/app/middleware"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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

func New(gin *gin.Engine, redisClient *redis.Client, params *models.DBConfigParam) *Route {
	return &Route{gin,
		usersController.NewUsersController(params),
		authController.NewAuthController(params),
		gin.Group(constants.ENDPOINT_V1),
		middleware.NewDeserializeMiddleware(params),
		filesController.NewFilesController(params),
	}
}

func (r *Route) PaveRoutes() {

	r.v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.authRouter()

	r.userRouter()

	r.filesRouter()
}
