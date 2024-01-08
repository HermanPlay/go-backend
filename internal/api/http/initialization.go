package http

import (
	"github.com/HermanPlay/backend/internal/api/http/routes"
	"github.com/HermanPlay/backend/internal/config"
	"github.com/HermanPlay/backend/internal/database"
	"github.com/HermanPlay/backend/pkg/repository"
	"github.com/HermanPlay/backend/pkg/service"
)

type Initialization struct {
	Cfg            *config.Config
	UserRepository repository.UserRepository
	UserService    service.UserService
	UserRoute      routes.UserRoute
	AuthRepository repository.AuthRepository
	AuthService    service.AuthService
	AuthRoute      routes.AuthRoute
}

func NewInitialization(
	config *config.Config,
	userRepo repository.UserRepository,
	userService service.UserService,
	UserRoute routes.UserRoute,
	authRepo repository.AuthRepository,
	authService service.AuthService,
	authRoute routes.AuthRoute,
) *Initialization {
	return &Initialization{
		Cfg:            config,
		UserRepository: userRepo,
		UserService:    userService,
		UserRoute:      UserRoute,
		AuthRepository: authRepo,
		AuthService:    authService,
		AuthRoute:      authRoute,
	}
}

func Init(cfg *config.Config) *Initialization {
	pgDb := database.NewPostgresDatabase(cfg).Connect()
	userRepositoryImpl := repository.NewUserRepository(pgDb)
	userServiceImpl := service.NewUserService(userRepositoryImpl)
	userRouteImpl := routes.NewUserRoute(userServiceImpl)
	authRepositoryImpl := repository.NewAuthRepository(pgDb, cfg)
	authServiceImpl := service.NewAuthService(authRepositoryImpl, userRepositoryImpl)
	authRouteImpl := routes.NewAuthRoute(authServiceImpl)
	initialization := NewInitialization(cfg, userRepositoryImpl, userServiceImpl, userRouteImpl, authRepositoryImpl, authServiceImpl, authRouteImpl)

	return initialization
}
