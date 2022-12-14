// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/KwokGH/kratos/internal/conf"
	"github.com/KwokGH/kratos/internal/data"
	"github.com/KwokGH/kratos/internal/server"
	"github.com/KwokGH/kratos/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	iUserRepo := data.NewUserRepo(dataData, logger)
	userService := service.NewUserService(iUserRepo, logger)
	httpServer := server.NewHTTPServer(confServer, userService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
