package restserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/alxhtp/monogo/config"
	"github.com/alxhtp/monogo/internal/server/rest/router"
	databasehelper "github.com/alxhtp/monogo/pkg/helper/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

type RestServer struct {
	app *fiber.App
	cfg *config.Config
	db  *gorm.DB
	ctx context.Context
}

func NewRestServer(ctx context.Context, cfg *config.Config) *RestServer {
	if cfg == nil {
		log.Fatalf("config is nil")
	}

	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
		BodyLimit:    cfg.BodyLimit,
		AppName:      cfg.AppName,
	})

	db, err := databasehelper.NewGormDB(ctx, &cfg.DatabaseConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return &RestServer{
		app: app,
		cfg: cfg,
		db:  db,
		ctx: ctx,
	}
}

func (s *RestServer) Start() error {
	// Add global middleware
	s.app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	s.app.Use(logger.New(logger.Config{
		Format:     "${time} ${status} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(s.cfg.AllowedOrigins, ","),
		AllowMethods: strings.Join(s.cfg.AllowedMethods, ","),
		AllowHeaders: strings.Join(s.cfg.AllowedHeaders, ","),
	}))

	// Add routes
	s.app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Swagger routes
	s.app.Get("/swagger/*", basicauth.New(basicauth.Config{
		Users: map[string]string{
			s.cfg.SwaggerUsername: s.cfg.SwaggerPassword,
		},
	}), swagger.New())

	// Register routes
	s.RegisterRoutes()

	// Start server
	errCh := make(chan error, 1)
	go func() {
		addr := fmt.Sprintf(":%d", s.cfg.AppPort)
		log.Printf("HTTP server listening on %s", addr)
		if err := s.app.Listen(addr); err != nil {
			errCh <- err
		}
		close(errCh)
	}()

	// Graceful shutdown on signal or context cancellation
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-s.ctx.Done():
		return s.Shutdown()
	case sig := <-sigCh:
		log.Printf("received signal: %s, shutting down", sig)
		return s.Shutdown()
	case err := <-errCh:
		return errors.New("failed to start server: " + err.Error())
	}
}

func (s *RestServer) Shutdown() error {
	if s.app == nil {
		return nil
	}
	shutdownCtx, cancel := context.WithTimeout(s.ctx, 10*time.Second)
	defer cancel()
	return s.app.ShutdownWithContext(shutdownCtx)
}

func (s *RestServer) RegisterRoutes() {
	dependencies := router.NewDependencies(s.app, s.db, s.cfg)
	router.UserRouter(dependencies)
}
