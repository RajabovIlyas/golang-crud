package server

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

type Server struct {
	gin         *gin.Engine
	cfg         *config.Config
	db          *database.Queries
	redisClient *redis.Client
}

func NewServer(gin *gin.Engine, cfg *config.Config, db *database.Queries, redisClient *redis.Client) *Server {
	return &Server{gin: gin, cfg: cfg, db: db, redisClient: redisClient}
}

func (s *Server) Run() error {

	srv := &http.Server{
		Addr:    s.cfg.Server.Port,
		Handler: s.gin,
	}

	go func() {
		log.Printf("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.gin.Run(srv.Addr); err != nil {
			log.Fatal().Msg("Error starting Server: " + err.Error())
		}
	}()

	if err := s.MapHandlers(s.gin); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msg("Server Shutdown:" + err.Error())
	}

	log.Info().Msg("Server Exited Properly")
	return nil
}
