package factory

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	mainbiz "my-source/sheet-payment/be/biz"
	authenhandler "my-source/sheet-payment/be/biz/auth"
	middlewarelogging "my-source/sheet-payment/be/biz/logging"
	"my-source/sheet-payment/be/repository"
)

var (
	app         *fiber.App
	bizInst     *mainbiz.MainBusiness
	authInst    *authenhandler.AuthHandler
	ssoInst     *authenhandler.SsoHandler
	loggingInst *middlewarelogging.Logger
)

func Factory() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it")
	}

	db := repository.InitDB()
	logRepo := repository.NewLogRepository(db)
	loggingInst = middlewarelogging.NewLogger(logRepo)

	memberRepo := repository.NewMemberRepository(db)
	blockRepo := repository.NewBlockRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	userRepo := repository.NewUserRepository(db)
	authInst = authenhandler.NewAuthHandler(userRepo)
	ssoInst = authenhandler.NewSSOHandler(userRepo)
	bizInst = mainbiz.NewMainBusiness(memberRepo, blockRepo, transactionRepo)
	app = fiber.New()
}

func GetApp() *fiber.App {
	return app
}

func GetBiz() *mainbiz.MainBusiness {
	return bizInst
}

func GetLogging() *middlewarelogging.Logger {
	return loggingInst
}

func GetAuth() *authenhandler.AuthHandler {
	return authInst
}

func GetSSO() *authenhandler.SsoHandler {
	return ssoInst
}
