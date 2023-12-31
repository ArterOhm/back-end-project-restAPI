package servers

import (
	"github.com/ArterOhm/back-end-project-restAPI/modules/middlewares/middlewaresHandlers"
	"github.com/ArterOhm/back-end-project-restAPI/modules/middlewares/middlewaresRepositories"
	middlewaresUsecases "github.com/ArterOhm/back-end-project-restAPI/modules/middlewares/middlewaresUsecase"
	monitorHandlers "github.com/ArterOhm/back-end-project-restAPI/modules/monitor/monitor_handlers"
	"github.com/ArterOhm/back-end-project-restAPI/modules/users/usersHandlers"
	"github.com/ArterOhm/back-end-project-restAPI/modules/users/usersRepositories"
	"github.com/ArterOhm/back-end-project-restAPI/modules/users/usersUsecases"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
	UsersModule()
}

type moduleFactory struct {
	router fiber.Router
	server *server
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		router: r,
		server: s,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	return middlewaresHandlers.MiddlewaresHandler(s.cfg, usecase)
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.server.cfg)

	m.router.Get("/", handler.HealthCheck)
}
func (m *moduleFactory) UsersModule() {
	repository := usersRepositories.UsersRepository(m.server.db)
	usecase := usersUsecases.UsersUsecase(m.server.cfg, repository)
	handler := usersHandlers.UsersHandler(m.server.cfg, usecase)

	router := m.router.Group("/users")

	router.Post("/signup", handler.SignUpCustomer)

}
