package di

import (
	"go.mongodb.org/mongo-driver/mongo"
	"shop-smart-api/internal/infrastructure/repository/mongo/mapper"
	"shop-smart-api/internal/infrastructure/repository/mongo/repository"
	"shop-smart-api/internal/pkg/jwt"
	"shop-smart-api/internal/usecase"
	"shop-smart-api/internal/usecase/user"
	"shop-smart-api/pkg"
)

type Container interface {
	ProvideUserUseCase() usecase.UserUseCase
}

type container struct {
	baseRepository repository.BaseRepository
	baseMapper     mapper.BaseMapper
	database       *mongo.Database
	serverConfig   pkg.Server
}

func CreateContainer(db *mongo.Database, sc pkg.Server) Container {
	br := repository.CreateBaseRepository()
	bm := mapper.CreateBaseMapper()

	return &container{br, bm, db, sc}
}

func (c *container) ProvideUserUseCase() usecase.UserUseCase {
	return c.resolveUserUseCaseDependencies(c.baseRepository, c.baseMapper)
}

func (c *container) resolveUserUseCaseDependencies(
	br repository.BaseRepository,
	bm mapper.BaseMapper,
) usecase.UserUseCase {
	jwtManager := jwt.CreateManager(c.serverConfig.Secret)

	userMapper := mapper.CreateUserMapper(bm)
	userRepository := repository.CreateUserRepository(br, c.database.Collection("users"), userMapper)
	userCreator := user.CreateCreator(userRepository)
	userAuthService := user.CreateAuthService(userRepository, jwtManager, userCreator)
	userFinder := user.CreateFinder(userRepository)
	userCollector := user.CreateCollector(userRepository)

	return user.CreateUserUseCase(userAuthService, userFinder, userCollector)
}
