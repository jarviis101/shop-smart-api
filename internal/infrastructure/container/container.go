package container

import (
	"database/sql"
	smsru "github.com/dmitriy-borisov/go-smsru"
	"shop-smart-api/internal/infrastructure/repository"
	"shop-smart-api/internal/pkg/jwt"
	"shop-smart-api/internal/pkg/sms"
	"shop-smart-api/internal/service"
	"shop-smart-api/internal/service/organization"
	"shop-smart-api/internal/service/otp"
	"shop-smart-api/internal/service/transaction"
	"shop-smart-api/internal/service/user"
	"shop-smart-api/pkg"
	"strconv"
)

type Container interface {
	ProvideUserService() service.UserService
	ProvideOTPService() service.OTPService
	ProvideOrganizationService() service.OrganizationService
	ProvideTransactionService() service.TransactionService
}

type container struct {
	database     *sql.DB
	serverConfig pkg.Server
}

func CreateContainer(db *sql.DB, sc pkg.Server) Container {
	return &container{db, sc}
}

func (c *container) ProvideUserService() service.UserService {
	return c.resolveUserServiceDependencies()
}

func (c *container) ProvideOTPService() service.OTPService {
	return c.resolveOTPServiceDependencies()
}

func (c *container) ProvideOrganizationService() service.OrganizationService {
	return c.resolveOrganizationServiceDependencies()
}

func (c *container) ProvideTransactionService() service.TransactionService {
	return c.resolveTransactionServiceDependencies()
}

func (c *container) resolveUserServiceDependencies() service.UserService {
	jwtManager := jwt.CreateManager(c.serverConfig.Secret)

	userRepository := repository.CreateUserRepository(c.database)
	userCreator := user.CreateCreator(userRepository)
	userFinder := user.CreateFinder(userRepository)
	userCollector := user.CreateCollector(userRepository)
	userModifier := user.CreateModifier(userRepository)

	return service.CreateUserService(userFinder, userCollector, userModifier, userCreator, jwtManager)
}

func (c *container) resolveOTPServiceDependencies() service.OTPService {
	debug, _ := strconv.ParseBool(c.serverConfig.Debug)
	smsClient := sms.CreateClient(smsru.NewClient(c.serverConfig.SmsApiKey), debug)

	otpGenerator := otp.CreateGenerator()
	otpRepository := repository.CreateOTPRepository(c.database)
	otpCreator := otp.CreateCreator(otpRepository, otpGenerator)
	otpSender := otp.CreateSender(otpCreator, smsClient)
	otpValidator := otp.CreateValidator(otpRepository, debug)

	return service.CreateOTPService(otpSender, otpValidator)
}

func (c *container) resolveOrganizationServiceDependencies() service.OrganizationService {
	organizationRepository := repository.CreateOrganizationRepository(c.database)
	organizationFinder := organization.CreateFinder(organizationRepository)

	return service.CreateOrganizationService(organizationFinder)
}

func (c *container) resolveTransactionServiceDependencies() service.TransactionService {
	transactionRepository := repository.CreateTransactionRepository(c.database)
	transactionFinder := transaction.CreateFinder(transactionRepository)

	return service.CreateTransactionService(transactionFinder)
}
