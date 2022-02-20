// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"lopingbest/GolangRESTFullAPI/app"
	"lopingbest/GolangRESTFullAPI/controller"
	"lopingbest/GolangRESTFullAPI/middleware"
	"lopingbest/GolangRESTFullAPI/repository"
	"lopingbest/GolangRESTFullAPI/service"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	categoryRepositoryImplementation := repository.NewCategoryRepository()
	db := app.NewDB()
	validate := validator.New()
	categoryServiceImplementation := service.NewCategoryService(categoryRepositoryImplementation, db, validate)
	categoryControllerImplementation := controller.NewCategoryController(categoryServiceImplementation)
	router := app.NewRouter(categoryControllerImplementation)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var categorySet = wire.NewSet(repository.NewCategoryRepository, wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImplementation)), service.NewCategoryService, wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImplementation)), controller.NewCategoryController, wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImplementation)))
