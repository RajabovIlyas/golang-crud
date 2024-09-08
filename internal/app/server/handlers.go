package server

import (
	authHttp "github.com/RajabovIlyas/golang-crud/internal/app/auth/delivery/http"
	authUseCase "github.com/RajabovIlyas/golang-crud/internal/app/auth/usecase"
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	cronUseCase "github.com/RajabovIlyas/golang-crud/internal/app/cron-job/usecase"
	fileHttp "github.com/RajabovIlyas/golang-crud/internal/app/file/delivery/http"
	fileRepository "github.com/RajabovIlyas/golang-crud/internal/app/file/repository"
	fileUseCase "github.com/RajabovIlyas/golang-crud/internal/app/file/usecase"
	"github.com/RajabovIlyas/golang-crud/internal/app/middleware"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	tokenRepository "github.com/RajabovIlyas/golang-crud/internal/app/token/repository"
	tokenUseCase "github.com/RajabovIlyas/golang-crud/internal/app/token/usecase"
	userHttp "github.com/RajabovIlyas/golang-crud/internal/app/user/delivery/http"
	userRepository "github.com/RajabovIlyas/golang-crud/internal/app/user/repository"
	userUseCase "github.com/RajabovIlyas/golang-crud/internal/app/user/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (s *Server) MapHandlers(g *gin.Engine) error {

	//Init Repository
	uRepo := userRepository.NewUserRepository(s.db)
	tRepo := tokenRepository.NewTokenRepository(s.db)
	fRepo := fileRepository.NewFileRepository(s.db)

	// Init useCases
	userUC := userUseCase.NewUserUseCase(s.cfg, uRepo, s.logger)
	tokenUC := tokenUseCase.NewTokenUseCase(s.cfg, tRepo, s.logger)
	authUC := authUseCase.NewAuthUseCase(s.cfg, userUC, tokenUC, s.logger)
	fileUC := fileUseCase.NewFileUseCase(s.cfg, fRepo, s.logger)
	cronUC := cronUseCase.NewCronUC(tokenUC, s.logger)

	// Init handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUC)
	userHandlers := userHttp.NewUserHandlers(s.cfg, userUC)
	fileHandlers := fileHttp.NewFileHandlers(s.cfg, fileUC)

	mw := middleware.NewMiddlewareManager(s.cfg, userUC, tokenUC, s.logger)

	g.Use(mw.Logger)

	v1 := g.Group(constants.ENDPOINT_V1)

	v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	health := v1.Group("/health")
	authGroup := v1.Group("/auth")
	userGroup := v1.Group("/users")
	fileGroup := v1.Group("/files")

	authHttp.MapAuthRoutes(authGroup, authHandlers, mw)
	userHttp.MapUsersRoutes(userGroup, userHandlers)
	fileHttp.MapFileRoutes(fileGroup, fileHandlers)

	health.GET("", func(g *gin.Context) {
		g.JSON(http.StatusOK, models.Message{Message: "OK"})
	})

	cronUC.DeleteAllToken()

	return nil
}
