package servers

import (
	"github.com/ArterOhm/back-end-project-restAPI/modules/middlewares/middlewaresHandlers"
	"github.com/ArterOhm/back-end-project-restAPI/modules/middlewares/middlewaresRepositories"
	middlewaresUsecases "github.com/ArterOhm/back-end-project-restAPI/modules/middlewares/middlewaresUsecase"
	monitorHandlers "github.com/ArterOhm/back-end-project-restAPI/modules/monitor/monitor_handlers"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
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
