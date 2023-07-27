package di

import (
	"database/sql"
	smsru "github.com/dmitriy-borisov/go-smsru"
	"shop-smart-api/internal/infrastructure/repository"
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
	database     *sql.DB
	serverConfig pkg.Server
}

func CreateContainer(db *sql.DB, sc pkg.Server) Container {
	return &container{db, sc}
}

func (c *container) ProvideUserUseCase() usecase.UserUseCase {
	return c.resolveUserUseCaseDependencies()
}

func (c *container) ProvideOTPUseCase() usecase.OTPUseCase {
	return c.resolveOTPUseCaseDependencies()
}

func (c *container) resolveUserUseCaseDependencies() usecase.UserUseCase {
	jwtManager := jwt.CreateManager(c.serverConfig.Secret)

	userRepository := repository.CreateUserRepository(c.database)
	userCreator := user.CreateCreator(userRepository)
	userAuthService := user.CreateAuthService(userRepository, jwtManager, userCreator)
	userFinder := user.CreateFinder(userRepository)
	userCollector := user.CreateCollector(userRepository)
	userModifier := user.CreateModifier(userRepository)

	return user.CreateUserUseCase(userAuthService, userFinder, userCollector, userModifier, userCreator)
}

func (c *container) resolveOTPUseCaseDependencies() usecase.OTPUseCase {
	debug, _ := strconv.ParseBool(c.serverConfig.Debug)
	smsClient := sms.CreateClient(smsru.NewClient(c.serverConfig.SmsApiKey), debug)

	otpGenerator := otp.CreateGenerator()
	otpRepository := repository.CreateOTPRepository(c.database)
	otpCreator := otp.CreateCreator(otpRepository, otpGenerator)
	otpSender := otp.CreateSender(otpCreator, smsClient)
	otpValidator := otp.CreateValidator(otpRepository, debug)

	return otp.CreateOTPUseCase(otpSender, otpValidator)
}
