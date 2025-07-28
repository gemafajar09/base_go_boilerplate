package http

import (
	"go-project/internal/config"
	"go-project/internal/delivery/http/handler"
	"go-project/internal/delivery/middleware"
	"go-project/internal/repository"
	"go-project/internal/usecase"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	Logger      *zap.Logger
	UserUsecase *usecase.UserUsecase
	// tambah usecase lain disini nanti
}

func initDeps(db *gorm.DB) (*Dependencies, error) {
	cfgzap := zap.NewProductionConfig()
	cfgzap.OutputPaths = []string{"app.log", "stdout"}
	logger, err := cfgzap.Build()
	if err != nil {
		return nil, err
	}

	// Setup repository dan usecase
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	return &Dependencies{
		Logger:      logger,
		UserUsecase: userUsecase,
	}, nil
}

func NewRouter(cfg config.Config, db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Setup CORS
	r.Use(middleware.CORSMiddleware())

	deps, err := initDeps(db)
	if err != nil {
		panic(err)
	}
	defer deps.Logger.Sync()

	r.Use(middleware.LoggerMiddleware(deps.Logger))

	authHandler := handler.NewAuthHandler(deps.UserUsecase)
	userHandler := handler.NewUserHandler(deps.UserUsecase)

	// Public routes
	r.GET("/health", handler.GetHealth)

	// Auth routes group
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// Protected routes group
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.JWTAuthMiddleware(cfg.JWT)) // pakai JWT middleware
	{
		apiGroup.GET("/profile", userHandler.GetProfile)
	}

	return r
}
