package server

import (
	"fmt"

	"users-api/internal/server/handlers"

	"users-api/internal/core/domain"
	"users-api/internal/core/services/processors"

	"github.com/labstack/echo/v4"
)

type Server struct {
	server             *echo.Echo
	getUsersHandler    *handlers.GetUsersHandler
	findUserHandler    *handlers.FindUserHandler
	saveUserHandler    *handlers.SaveUserHandler
	getAddresssHandler *handlers.GetAddresssHandler
	findAddressHandler *handlers.FindAddressHandler
	saveAddressHandler *handlers.SaveAddressHandler
}

func Start(options *domain.Options, processors *processors.Processors) {
	api := &Server{
		server:             newHttpServer(),
		getUsersHandler:    handlers.NewGetUsersHandler(processors.GetUsersProcessor),
		findUserHandler:    handlers.NewFindUserHandler(processors.FindUserProcessor),
		saveUserHandler:    handlers.NewSaveUserHandler(options, processors.SaveUserProcessor),
		getAddresssHandler: handlers.NewGetAddresssHandler(processors.GetAddressProcessor),
		findAddressHandler: handlers.NewFindAddressHandler(processors.FindAddressProcessor),
		saveAddressHandler: handlers.NewSaveAddressHandler(options, processors.SaveAddressProcessor),
	}

	api.StartRoutes()

	if err := api.server.Start(":" + options.Config.Http.Port); err != nil {
		fmt.Println("Shutting down the server")
	}
}

func newHttpServer() (server *echo.Echo) {
	server = echo.New()
	return
}

func (s *Server) StartRoutes() {
	s.server.GET("/users", s.findUserHandler.FindUser)
	s.server.GET("/users/:id", s.getUsersHandler.GetUsers)
	s.server.GET("/users/:userId/address", s.findAddressHandler.FindAddress)
	s.server.POST("/users/:userId/address", s.saveAddressHandler.SaveAddress)
	s.server.GET("/users/:userId/address/:addressId", s.getAddresssHandler.GetAddress)
	s.server.POST("/users", s.saveUserHandler.SaveUser)
}
