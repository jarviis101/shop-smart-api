package di

import (
	smsru "github.com/dmitriy-borisov/go-smsru"
	"go.mongodb.org/mongo-driver/mongo"
	"shop-smart-api/internal/infrastructure/repository/mongo/mapper"
	"shop-smart-api/internal/infrastructure/repository/mongo/repository"
	"shop-smart-api/internal/pkg/jwt"
	"shop-smart-api/internal/pkg/sms"
	"shop-smart-api/internal/usecase"
	"shop-smart-api/internal/usecase/otp"
	"shop-smart-api/internal/usecase/user"
	"shop-smart-api/pkg"
	"strconv"
)

type Container interface {
	ProvideUserUseCase() usecase.UserUseCase
	ProvideOTPUseCase() usecase.OTPUseCase
}

type container struct {
	baseRepository repository.BaseRepository
	baseMapper     mapper.BaseMapper
	database       *mongo.Database
	serverConfig   pkg.Server
}

func CreateContainer(db *mongo.Database, sc pkg.Server) Container {
	baseRepository := repository.CreateBaseRepository()
	baseMapper := mapper.CreateBaseMapper()

	return &container{baseRepository, baseMapper, db, sc}
}

func (c *container) ProvideUserUseCase() usecase.UserUseCase {
	return c.resolveUserUseCaseDependencies(c.baseRepository, c.baseMapper)
}

func (c *container) ProvideOTPUseCase() usecase.OTPUseCase {
	return c.resolveOTPUseCaseDependencies(c.baseRepository, c.baseMapper)
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

func (c *container) resolveOTPUseCaseDependencies(
	br repository.BaseRepository,
	bm mapper.BaseMapper,
) usecase.OTPUseCase {
	debug, _ := strconv.ParseBool(c.serverConfig.Debug)
	smsClient := sms.CreateClient(smsru.NewClient(c.serverConfig.SmsApiKey), debug)

	otpGenerator := otp.CreateGenerator()
	otpMapper := mapper.CreateOTPMapper(bm)
	otpRepository := repository.CreateOTPRepository(br, c.database.Collection("otp"), otpMapper)
	otpCreator := otp.CreateCreator(otpRepository, otpGenerator)
	otpSender := otp.CreateSender(otpCreator, smsClient)
	otpValidator := otp.CreateValidator(otpRepository, debug)

	return otp.CreateOTPUseCase(otpSender, otpValidator)
}
