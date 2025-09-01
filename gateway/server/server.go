package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	kitlog "github.com/go-kit/log"

	"github.com/ppeymann/Planora.git/pkg/auth"
	"github.com/ppeymann/Planora.git/pkg/env"
	"github.com/ppeymann/Planora/gateway/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Router        *gin.Engine
	Logger        kitlog.Logger
	paseto        auth.TokenMaker
	instrumenting serviceInstrumenting
}

// EnvMode specified the running env 'release' represents production mode and ” represents development.
// it depended on gin GIN_MODE env for unifying and simplicity of setting.
var EnvMode = ""

func NewServer(logger kitlog.Logger) *Server {
	svr := &Server{
		Logger:        logger,
		instrumenting: newServiceInstrumenting(),
	}

	router := gin.New()
	router.Use(gin.Recovery())

	// setting swagger info if not in production mode
	if env.GetEnv("SWAGGER_ENABLED", "false") == "true" {
		docs.SwaggerInfo.Title = fmt.Sprintf("TODO App Backend [ AuthMode: %s ]", "Paseto")
		docs.SwaggerInfo.Description = "The Swagger Documentation For TODO Backend API server."
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = env.GetEnv("HOST_URL", "localhost:8080")
		docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	}

	if env.GetEnv("CORS_ENABLE", "false") == "true" {
		router.Use(svr.cors())
	}

	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalln(err)
	}

	svr.Router = router

	svr.Router.GET("/metrics", svr.prometheus())

	svr.paseto, err = auth.NewPasetoMaker(env.GetEnv("JWT", ""))
	if err != nil {
		log.Fatal(err)
	}

	return svr
}

func (s *Server) Listen() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()

	if env.GetEnv("SWAGGER_ENABLE", "false") == "true" {
		s.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	srv := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		Addr:              fmt.Sprintf("%s%s", env.GetEnv("HOST", "localhost"), env.GetEnv("PORT", ":8080")),
		Handler:           s.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http listener stopped: %s", err)
		}
	}()

	<-ctx.Done()
	stop()

	log.Println("shutting down gracefully Planora app server, press Ctrl+C again to force")

	// The context is used to inform the server it has 30 seconds to finish
	// the request it is currently handling

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown: ", err)
	}

	log.Println("Planora service exiting")
}
