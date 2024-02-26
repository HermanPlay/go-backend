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
	DevRoute       routes.DevRoute
	UserRepository repository.UserRepository
	UserService    service.UserService
	UserRoute      routes.UserRoute
	AuthRepository repository.AuthRepository
	AuthService    service.AuthService
	AuthRoute      routes.AuthRoute
}

func NewInitialization(
	config *config.Config,
	devRoute routes.DevRoute,
	userRepo repository.UserRepository,
	userService service.UserService,
	UserRoute routes.UserRoute,
	authRepo repository.AuthRepository,
	authService service.AuthService,
	authRoute routes.AuthRoute,
) *Initialization {
	return &Initialization{
		Cfg:            config,
		DevRoute:       devRoute,
		UserRepository: userRepo,
		UserService:    userService,
		UserRoute:      UserRoute,
		AuthRepository: authRepo,
		AuthService:    authService,
		AuthRoute:      authRoute,
	}
}

func Init(cfg *config.Config) *Initialization {
	db, err := database.NewPostgresDatabase(cfg)
	if err != nil {
		panic(err)
	}
	pgDb := db.Connect()
	devRouteImpl := routes.NewDevRoute()
	userRepositoryImpl := repository.NewUserRepository(pgDb)
	userServiceImpl := service.NewUserService(userRepositoryImpl)
	userRouteImpl := routes.NewUserRoute(userServiceImpl)
	authRepositoryImpl := repository.NewAuthRepository(pgDb, cfg)
	authServiceImpl := service.NewAuthService(authRepositoryImpl, userRepositoryImpl)
	authRouteImpl := routes.NewAuthRoute(authServiceImpl)
	initialization := NewInitialization(cfg, devRouteImpl, userRepositoryImpl, userServiceImpl, userRouteImpl, authRepositoryImpl, authServiceImpl, authRouteImpl)

	return initialization
}
